package main

import (
	"context"
	"encoding/json"
	"grpc-rest-api-crud/model"
	"grpc-rest-api-crud/protofile"

	"github.com/cpradeepc/findinbytes"

	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// 3rd party package to manage the CRUD operation with pointer person, data's all operation work in [][]byte like buffer
var person = findinbytes.NewByteJsonObj()

type server struct {
	protofile.UnimplementedPersonServer
}

// grpc service to  create person via http client json input
func (s *server) CreateData(c context.Context, req *protofile.JsonStruct) (*protofile.JsonStruct, error) {
	//fmt.Println(">> hello from server")
	uRec := model.Person_ST{}

	err := json.Unmarshal(req.JsonData, &uRec)
	if err != nil {
		log.Println(">> error in unmarshalling into object:", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	////fmt.Println(">> uRec: ", uRec, person.Data == nil)
	dbSize := person.GetDataSize()
	//fmt.Println(">> dbsize: ", dbSize)
	uRec.ID = dbSize + 1

	//fmt.Println(">> uRec :", uRec)

	nByte, err := json.Marshal(&uRec)
	if err != nil {
		log.Println(">> error in marshalling from object:", err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	person.Data = append(person.Data, nByte)
	//fmt.Printf(">> data.Data :%s\n", person.Data)

	resp := &protofile.JsonStruct{JsonData: nByte}
	return resp, nil
}

// grpc service to  delete person data by id
func (s *server) DeleteData(c context.Context, req *protofile.JsonStruct) (*protofile.JsonStruct, error) {
	//fmt.Println(">> DeleteDataById")

	ok, err := person.DeleteByteById(req.JsonData)
	resp := &protofile.JsonStruct{}
	log.Println(">> ok, err :", ok, err)
	if !ok || err != nil {
		//resp.JsonData = []byte("false")
		//log.Printf(">> resp.JsonData:%s\n", resp.JsonData)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	resp.JsonData = []byte("true")
	log.Printf(">> resp.JsonData:%s\n", resp.JsonData)
	return resp, nil

}

// grpc service to  get person data by id
func (s *server) GetDataById(c context.Context, req *protofile.JsonStruct) (*protofile.JsonStruct, error) {
	//fmt.Println(">> hello from server")

	idx, err := person.FindGetExactExistsIndexs(req.JsonData)
	// idx, err := person.FindGetExactExistsIndexs([]byte(find))
	if err != nil {
		log.Println(">> error in finding :", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	//fmt.Printf(">> index : %+v\n", idx)

	found := person.Data[idx]

	//fmt.Printf(">> uRec : %+v\n", found)

	resp := &protofile.JsonStruct{JsonData: found}
	return resp, nil
}

// grpc service to  get person data by person any json field-value
func (s *server) GetAllDataByField(c context.Context, req *protofile.JsonStruct) (*protofile.JsonStruct, error) {
	//fmt.Println(">> hello from server,GetAllDataByField")

	indices, err := person.GetIndexesExistsAllFields(req.JsonData)
	// idx, err := person.FindGetExactExistsIndexs([]byte(find))
	if err != nil {
		log.Println(">> error in finding :", err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	//fmt.Printf(">> indices : %+v\n", indices)

	found, err := person.GetDataByIndicesIn1D(indices)
	if err != nil {
		log.Println(">> error in tracking the records in db :", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	//fmt.Printf(">> uRec : %+v\n", found)

	resp := &protofile.JsonStruct{JsonData: found}
	return resp, nil
}

func main() {

	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()

	protofile.RegisterPersonServer(srv, &server{})
	reflection.Register(srv)
	e := srv.Serve(listener)
	if e != nil {
		log.Panic(e)
	}
}
