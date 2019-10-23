package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func ClientInterceptorA(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	log.Println("A - start: ", start)

	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Logic after invoking the invoker
	log.Printf("A - Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)
	log.Println("A - End: ", time.Now())
	return err
}

func ClientInterceptorB(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	// Logic before invoking the invoker
	start := time.Now()
	log.Println("B - start: ", start)

	// Calls the invoker to execute RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	// Logic after invoking the invoker
	log.Printf("B - Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)
	log.Println("B - End: ", time.Now())
	return err
}
