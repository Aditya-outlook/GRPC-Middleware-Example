package parametervalidation

// package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var rpc2DomainMappingPath = "../common"
var rpcField2DomainField = make(map[string]string, 0)

func loadRPC2DomainMappingFromFile(filePath string) {
	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("RPC-Common: Error opening mapping file: ", err)
	}

	var jsonContent interface{}
	err = json.Unmarshal(jsonBytes, &jsonContent)
	if err != nil {
		log.Fatal("unmarshaling failed for mapping JSON")
	}

	serviceDetails := jsonContent.(map[string]interface{})

	s := strings.Split(filePath, "/")
	var fieldPrefix string = fmt.Sprintf(strings.Split(s[len(s)-1], ".")[0])

	for serviceName, serviceContent := range serviceDetails {
		fieldPrefix += "."
		fieldPrefix += serviceName
		parseServiceContent(fieldPrefix, serviceContent)
	}
}

func parseServiceContent(fieldPrefix string, serviceContent interface{}) {
	rpcDetails, err := serviceContent.(map[string]interface{})
	if !err {
		log.Fatal("RPC-Common: ServiceObject: Expecting a JSON object got something else")
	}
	for rpcName, rpcContent := range rpcDetails {
		fieldPrefix = fieldPrefix + "/" + rpcName
		parseRPCContent(fieldPrefix, rpcContent)
	}
}

func parseRPCContent(fieldPrefix string, rpcContent interface{}) {
	fields, err := rpcContent.(map[string]interface{})
	if !err {
		log.Fatal("RPC-Common: RPCObject: Expecting a JSON object; got something else")
	}

	for rpcFieldName, domainFieldName := range fields {
		rpcField2DomainField["/"+fieldPrefix+"/"+rpcFieldName] = fmt.Sprintf("%v", domainFieldName)
	}
}

// GetDomainFieldName returns the mapped domain field name of a given rpc field
func GetDomainFieldName(rpc string, rpcFieldName string) string {
	key := rpc + "/" + rpcFieldName
	if len(rpcField2DomainField[key]) == 0 {
		protoFileName := strings.Split(strings.Split(rpc, "/")[1], ".")[0]
		loadRPC2DomainMappingFromFile("C:/Users/akadiyal/Documents/Go/src/GRPC-Middleware-Example/releases/" + protoFileName + ".json")
	}

	return rpcField2DomainField[key]
}

// func main() {
// 	loadRPC2DomainMappingFromFile("C:/Users/akadiyal/Documents/Go/src/GRPC-Middleware-Example/releases/releases.json")
// 	fmt.Println(rpcField2DomainField)
// }
