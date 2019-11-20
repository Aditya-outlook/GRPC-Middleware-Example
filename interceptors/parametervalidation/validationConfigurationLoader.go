package parametervalidation

// package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//RegExFilePath ... path to read RegEx in production
var RegExFilePath = "/etc/common/ParameterValidator.json"

// ValidationConfiguration ... regEx structure
type ValidationConfiguration struct {
	RegExp         string
	MinLength      int
	MaxLength      int
	ErrDescription string
}

// ValidationConfiguration asfsafgsg sg sgsg
var validationConfigurationMap = make(map[string]*ValidationConfiguration) // Key is a string, combination of msName and fieldName

func loadValidationRegExFromFile(filePath string) {
	if len(filePath) == 0 {
		filePath = RegExFilePath
	}

	jsonBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("error opening regex file: ", err)
	}

	var jsonContent interface{}
	err = json.Unmarshal(jsonBytes, &jsonContent)
	if err != nil {
		log.Fatal("unmarshaling failed for regEx JSON")
	}

	regExContent := jsonContent.(map[string]interface{})

	for msName, msContent := range regExContent {
		parseMsContent(msName, msContent)
	}
}

func parseMsContent(msName string, msContent interface{}) {
	// Use type assertions to ensure that the msContent is a JSON object
	msField, err := msContent.(map[string]interface{})
	if !err {
		log.Fatal("msObj: Expecting a JSON object; got something else")
	}
	for msFieldName, msFieldContent := range msField {
		parseMsFieldContent(msFieldContent, msFieldName, msName)
	}
}

func parseMsFieldContent(msFieldContent interface{}, msFieldName string, msName string) {
	msAttr, err := msFieldContent.(map[string]interface{})
	if !err {
		log.Fatal("msObj: Expecting a JSON object; got something else")
	}
	for msAttrName, msAttrValue := range msAttr {
		parseMsAttrValues(msAttrValue, msFieldName, msName, msAttrName)
	}
}

func parseMsAttrValues(msAttrValue interface{}, msFieldName string, msName string, msAttrName string) {
	var key string = msName + "." + msFieldName

	if validationConfigurationMap[key] == nil {
		validationConfigurationMap[key] = &ValidationConfiguration{}
	}

	switch msAttrName {
	case "regex":
		// Make sure that regex is string
		switch msAttrValue := msAttrValue.(type) {
		case string:
			validationConfigurationMap[key].RegExp = msAttrValue
		default:
			log.Println("Incorrect Type for", msAttrValue, msAttrName, msFieldName)
		}
	case "max_length":
		// Make sure that max_length is int
		switch msAttrValue := msAttrValue.(type) {
		case float64:
			validationConfigurationMap[key].MaxLength = int(msAttrValue)
		default:
			log.Println("Incorrect type for", msAttrValue, msAttrName, msFieldName)
		}
	case "min_length":
		// Make sure that min_length is int
		switch msAttrValue := msAttrValue.(type) {
		case float64:
			validationConfigurationMap[key].MinLength = int(msAttrValue)
		default:
			log.Println("Incorrect type for:", msAttrValue, msAttrName, msFieldName)
		}
	case "description":
		// Make sure that description is string
		switch msAttrValue := msAttrValue.(type) {
		case string:
			validationConfigurationMap[key].ErrDescription = msAttrValue
		default:
			log.Println("Incorrect Type for", msAttrValue, msAttrName, msFieldName)
		}
	}
}

// GetValidationConfiguration returns the validation configuration for a given field
func GetValidationConfiguration(key string) *ValidationConfiguration {

	if validationConfigurationMap[key] == nil {
		loadValidationRegExFromFile("C:/Users/akadiyal/Documents/Go/src/GRPC-Middleware-Example/common/ParameterValidation.json")
	}
	return validationConfigurationMap[key]
}

// func main() {
// 	loadValidationRegExFromFile("C:/Users/akadiyal/Documents/Go/src/GRPC-Middleware-Example/common/ParameterValidation.json")
// 	fmt.Println(validationConfigurationMap)
// }
