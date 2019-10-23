package interceptors

/*
 * Follow the link below for more details.
 * https://blog.gopheracademy.com/advent-2017/go-grpc-beyond-basics/
 */

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// ServerInterceptorA is an example of unary server interceptor
func ServerInterceptorA(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Before calling the handler
	start := time.Now()
	log.Println("A - start: ", start)

	// Calls the handler
	h, err := handler(ctx, req)

	// After calling the handler
	log.Printf("A - Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	log.Println("A - End: ", time.Now())

	return h, err
}

// ServerInterceptorB is another example of unary server interceptor
func ServerInterceptorB(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Before calling the handler
	start := time.Now()
	log.Println("B - start: ", start)

	// Calls the handler
	h, err := handler(ctx, req)

	// After calling the handler
	log.Printf("B - Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)
	log.Println("B - End: ", time.Now())

	return h, err
}
