package parametervalidation

/*
 * Follow the link below for more details.
 * https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/
 */

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	pbCommon "GRPC-Middleware-Example/releases"

	"google.golang.org/grpc"
)

// RPCParameterValidator is an example of unary server interceptor
func RPCParameterValidator(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	var status interface{}
	var err error = errors.New("Parameter validation failed")

	log.Println("ServerInterceptor:Parameter Validation - Preprocessing step(s) before calling the handler", info.FullMethod, req)

	valid, status := validateParameters(info.FullMethod, req) // TODO Check if the status can be used eliminating the boolean ? Faced problem with status == nil

	if valid {
		log.Printf("ServerInerceptor:Parameter Validation - Invoking handler / Next interceptor - Method:%s\n", info.FullMethod)
		status, err = handler(ctx, req)
		log.Printf("ServerInterceptor:Parameter Validation - Response from handler - Method:%s\tDuration:%s\tError:%v\n",
			info.FullMethod,
			time.Since(start),
			err)

		log.Println("ServerInterceptor:Parameter Validation - Postprocessing step(s) after calling the handler")
	} else {
		log.Print("ServerInterceptor:Parameter Validation failed with ", status)
		err = errors.New("Parameter validation failed")
	}

	return status, err
}

func validateParameters(rpc string, req interface{}) (bool, *pbCommon.Status) {
	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	for i := 0; i < t.Elem().NumField(); i++ {
		if strings.HasPrefix(t.Elem().Field(i).Name, "XXX_") {
			continue
		}

		fieldName := t.Elem().Field(i).Name
		fieldValue := strings.TrimSpace(fmt.Sprintf("%v", v.Elem().Field(i).Interface())) // Convert interface to string - https://yourbasic.org/golang/interface-to-string/
		validationConf := getValidationConf(rpc, fieldName)

		if errDetails := validateField(fieldName, fieldValue, validationConf); len(errDetails) != 0 {
			return false, &pbCommon.Status{StatusCode: pbCommon.StatusCode_FAILURE,
				StatusDescription: &pbCommon.StatusDescription{
					DescriptionCode: pbCommon.DescriptionCode_INVALID_ARGUMENT,
					Description:     errDetails}}
		}
	}

	return true, nil
}

// ValidateDetails ... Checks for valid length and then valid regex
func validateField(name string, value string, validationConf *ValidationConfiguration) string {

	length := utf8.RuneCountInString(value)

	if length < validationConf.MinLength || length > validationConf.MaxLength {
		// errDetails += fmt.Sprintf("Parameter: Name(%s), Value(%s): Input validation has failed for length check\n", name, value)
		return "Parameter: Name(" + name + "), " + "Value(" + value + "): Input validation has failed for length check\n"
	}

	matched, _ := regexp.MatchString(validationConf.RegExp, value)
	if !matched {
		// errDetails += fmt.Sprintf("Parameter(%s)(%s): %s\n", name, value, description)
		return "Parameter: Name(" + name + "), " + "Value(" + value + "): " + validationConf.ErrDescription + "\n"
	}

	return ""
}

func getValidationConf(rpc string, fieldName string) *ValidationConfiguration {
	// s := strings.Split(rpc, "/")
	// fmt.Println(s)
	// fmt.Println("Proto file : ", strings.Split(s[1], ".")[0])
	// fmt.Println("Servce : ", strings.Split(s[1], ".")[1])
	// fmt.Println("RPC : ", s[])

	// if mapper == nil {
	// 	mapper = make(map[string]ValidationConfiguration)
	// 	mapper[rpc+"/Version"] = ValidationConfiguration{RegExp: "^([0-9A-Za-z.-]+)$", MinLength: 1, MaxLength: 40, ErrDescription: ""}
	// 	mapper[rpc+"/Param2"] = ValidationConfiguration{RegExp: "^([0-9A-Za-z.-]+)$", MinLength: 2, MaxLength: 40, ErrDescription: ""}
	// 	mapper[rpc+"/Param3"] = ValidationConfiguration{RegExp: "^([0-9A-Za-z.-]+)$", MinLength: 3, MaxLength: 40, ErrDescription: ""}
	// }
	// return mapper[rpc+"/"+fieldName]

	// fmt.Println("Mapped Field Name: " + GetDomainFieldName(rpc, fieldName))
	// fmt.Println("Domain field name", GetValidationConfiguration(GetDomainFieldName(rpc, fieldName)))
	return GetValidationConfiguration(GetDomainFieldName(rpc, fieldName))
}
