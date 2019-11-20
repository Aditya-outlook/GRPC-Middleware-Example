package interceptors

/*
 * Follow the link below for more details.
 * https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/
 */

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"google.golang.org/grpc"
)

// ServerInterceptorA is an example of unary server interceptor
func ServerInterceptorA(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	log.Println("ServerInterceptor A - Request Params: ", req)

	t := reflect.TypeOf(req)
	v := reflect.ValueOf(req)

	var fieldName string
	var fieldValue string

	s := strings.Split(info.FullMethod, "/")
	fmt.Println(s)
	fmt.Println("Proto file : ", strings.Split(s[1], ".")[0])
	fmt.Println("Service : ", strings.Split(s[1], ".")[1])
	fmt.Println("RPC : ", s[2])

	for i := 0; i < t.Elem().NumField(); i++ {
		if strings.HasPrefix(t.Elem().Field(i).Name, "XXX_") {
			continue
		}

		fieldName = t.Elem().Field(i).Name
		fieldValue = fmt.Sprintf("%v", v.Elem().Field(i).Interface()) // Convert interface to string - https://yourbasic.org/golang/interface-to-string/
		fmt.Printf("Field: %s\tValue: %v\n", fieldName, fieldValue)
		fmt.Printf("Mapped validation configuration : WIFI/%s\n", fieldName)

		mapper := make(map[string]regEx)
		mapper["Version"] = regEx{RegExp: "^([0-9A-Za-z_.-]+)$", MinLength: 9, MaxLength: 40}

		regexp.MatchString(mapper[fieldName].RegExp, fieldValue)

	}

	// Preprocessing
	log.Println("ServerInterceptor A - Preprocessing step(s) before calling the handler")

	// Before calling the handler
	log.Printf("ServerInterceptor A - Invoking handler / Next interceptor - Method:%s\n", info.FullMethod)
	// Calls the handler
	h, err := handler(ctx, req)

	// After calling the handler
	log.Printf("ServerInterceptor A - Response from handler - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	// Postprocessing
	log.Println("ServerInterceptor A - Postprocessing step(s) after calling the handler")

	return h, err
}

// ServerInterceptorB is another example of unary server interceptor
func ServerInterceptorB(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	// Preprocessing
	log.Println("ServerInterceptor B - Preprocessing step(s) before calling the handler")

	// Before calling the handler
	log.Printf("ServerInterceptor B - Invoking handler / Next interceptor - Method:%s\n", info.FullMethod)
	// Calls the handler
	h, err := handler(ctx, req)

	// After calling the handler
	log.Printf("ServerInterceptor B - Response from handler - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	// Postprocessing
	log.Println("ServerInterceptor B - Postprocessing step(s) after calling the handler")

	return h, err
}

type regEx struct {
	RegExp    string
	MinLength int
	MaxLength int
}

// regExInfo ... Wifi Controller parameters to which reggEx checks are needed
type regExInfo struct {
	CWLCUserName      regEx
	CWLCPassword      regEx
	CWLCAlias         regEx
	CWLCUrl           regEx
	TenantName        regEx
	AcctUUID          regEx
	APSoftwareVersion regEx
	CWLCUuid          regEx
	TenantUUID        regEx
}
