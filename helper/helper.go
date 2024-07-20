package helper

import (
	"encoding/json"
	"fmt"
	"grpc-rest-api-crud/protofile"
)

// it takes 2 values in a key-value and reture json object format,
// e.g `{"name":"person_name"}` in string type
func GetJsonObj(key string, val any) (jsonObj string) {
	jsonStr := fmt.Sprintf(`{"%s":%v}`, key, val)
	return jsonStr
}

// it takes 2 values in a key-value and reture proto message type as JsonStruct format
func GetProtoJsonStruct(key string, val any) *protofile.JsonStruct {
	jsonStr := fmt.Sprintf(`{"%s":%v}`, key, val)
	return &protofile.JsonStruct{JsonData: []byte(jsonStr)}
}

// it takes json bytes values and reture proto message and error
func GetProtoJsonStructByStruct(val any) (*protofile.JsonStruct, error) {
	valByte, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}
	p := &protofile.JsonStruct{JsonData: valByte}
	return p, nil
}
