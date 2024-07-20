package main

import (
	"context"
	"encoding/json"

	"grpc-rest-api-crud/helper"
	"grpc-rest-api-crud/model"
	"grpc-rest-api-crud/protofile"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client protofile.PersonClient //interface in go but in proto called service

func main() {
	// Connection to internal grpc server
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials())) //to make grpc request to server
	if err != nil {
		log.Println("error invoked to established grpc connetion to server")
		panic(err)
	}
	client = protofile.NewPersonClient(conn)

	// implement rest api
	r := gin.Default()
	r.GET("/api/person/:id", getPersonById)       //reply to endpoint user
	r.POST("/api/person", createPerson)           //reply to endpoint user
	r.GET("/api/persons", getPersonByAnyField)    //reply to endpoint user
	r.DELETE("/api/person/:id", deletePersonById) //reply to endpoint user
	r.Run(":8080")                                // 8080

}

func createPerson(c *gin.Context) {
	//variableName := c.Param("message") //get data from endpoint user
	// var stData model.Person_ST
	respPerson := model.Person_ST{}
	err := c.BindJSON(&respPerson)
	if err != nil {
		log.Println("-- error in unmarshalling into object:", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}
	// //fmt.Println("-- stData:", stData)

	ubyte, err := json.Marshal(respPerson)
	if err != nil {
		log.Println("-- error in marshalling from object:", err)
		return
	}

	req := &protofile.JsonStruct{JsonData: ubyte} //data pass as pointer in variable

	resp, err := client.CreateData(context.TODO(), req) //send data to grpc server
	if err != nil {
		log.Println("-- error from server:", err.Error())
		c.JSON(http.StatusOK, gin.H{"response": nil, "error": err.Error()})
		return
	}
	//fmt.Println("err req to grpc server :", err)
	//fmt.Printf("-- resp req to grpc server :%s\n", resp)

	err = json.Unmarshal(resp.JsonData, &respPerson)
	if err != nil {
		log.Println("-- error in unmarshalling into object:", err)
		c.JSON(http.StatusOK, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": respPerson, "error": nil})
}

func deletePersonById(c *gin.Context) {
	//fmt.Println("deletePerson")
	id := c.Param("id") //get data from endpoint user
	//fmt.Println("-- id:", id)
	idInt, _ := strconv.Atoi(id)
	//req := helper.GetprotofileJsonStruct("id", idInt)
	idObj := helper.GetJsonObj("id", idInt)
	//fmt.Println("-- idObj:", idObj)

	req := &protofile.JsonStruct{JsonData: []byte(idObj)} //data pass as pointer in variable

	resp, err := client.DeleteData(context.TODO(), req) //send data to grpc server

	//fmt.Printf("-- err:%v\n", err)
	if err != nil {
		log.Println("-- error in deleting:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "err": err.Error()})
		return
	}

	log.Println("-- resp.jsonData:", resp.JsonData)
	c.JSON(http.StatusOK, gin.H{"response": "deleted:" + id, "error": nil})
}

func getPersonById(c *gin.Context) {
	//fmt.Println("getPersonById")
	id := c.Param("id") //get data from endpoint user
	//fmt.Println("-- id:", id)
	idInt, _ := strconv.Atoi(id)
	req := helper.GetProtoJsonStruct("id", idInt)

	//fmt.Println("-- rep grpc struct:", req)
	//req := &proto.JsonStruct{JsonData: []byte(`{"id":1}`)} //data pass as pointer in variable

	resp, err := client.GetDataById(context.TODO(), req) //send data to grpc server
	if err != nil {
		log.Println("-- error in founding:", err)
		c.JSON(http.StatusNotFound, gin.H{"response": nil, "error": err.Error()})
		return
	}
	//fmt.Println("err req to grpc server :", err)
	//fmt.Printf("-- resp req to grpc server :%s\n", resp)

	respPerson := model.Person_ST{}
	err = json.Unmarshal(resp.JsonData, &respPerson)
	if err != nil {
		log.Println("-- error in unmarshalling into object:", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"response": nil, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": respPerson, "error": nil})
}

func getPersonByAnyField(c *gin.Context) {
	//fmt.Println("getPersonByAnyField")

	var pSrch model.SearchField
	err := c.BindJSON(&pSrch)
	//fmt.Println("-- pSrch:", pSrch)
	if err != nil {
		log.Println("-- invalid json format:", err)
		c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
		return
	}
	find, err := json.Marshal(pSrch)
	if err != nil {
		log.Println("-- error in marshalling data:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"response": nil, "error": err.Error()})
		return
	}
	req := &protofile.JsonStruct{JsonData: find}               //data pass as pointer in variable
	resp, err := client.GetAllDataByField(context.TODO(), req) //send data to grpc server, multiple json obj in 1 string ,e.g.: "[{"id":1,"name":"name_any"},{"id":1,"name":"name_any"}]"
	if err != nil {
		log.Println("-- error in getting data:", err.Error())
		c.JSON(http.StatusNotFound, gin.H{"response": nil, "error": err.Error()})
		return
	}

	//fmt.Println("err req to grpc server :", err)
	//fmt.Printf("-- resp req to grpc server :%s\n", resp)

	raws := strings.Split(string(resp.JsonData), ";")
	respPerson := []model.Person_ST{}

	//fmt.Printf("--  raws :%s\n:", raws)
	for i, v := range raws {
		p := model.Person_ST{}
		err = json.Unmarshal([]byte(v), &p)
		if err != nil {
			log.Println("-- error in unmarshalling into object:", err.Error(), "at: ", i)
			c.JSON(http.StatusBadRequest, gin.H{"response": nil, "error": err.Error()})
			return
		}
		respPerson = append(respPerson, p)

	}

	//fmt.Println("--  respPerson:", respPerson)
	c.JSON(http.StatusOK, gin.H{"response": respPerson, "error": nil})
}
