package interceptors

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

// ClientInterceptorA is an example that intercepts the RPC calls
func ClientInterceptorA(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	// Preprocessing
	log.Println("ClientInterceptor A - Preprocessing step(s) before invoking the RPC")

	// Before invoking the RPC
	log.Printf("ClientInterceptor A - Invoking RPC method=%s;", method)

	// Invoking the RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	// After invoking the RPC
	log.Printf("ClientInterceptor A - Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)

	// Postprocessing
	log.Println("ClientInterceptor A - Postprocessing after invoking the RPC")
	return err
}

// ClientInterceptorB is another example that intercepts the RPC calls
func ClientInterceptorB(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	start := time.Now()
	// Preprocessing
	log.Println("ClientInterceptor B - Preprocessing step(s) before invoking the RPC")

	// Before invoking the RPC
	log.Printf("ClientInterceptor B - Invoking RPC method=%s;", method)

	// Invoking the RPC
	err := invoker(ctx, method, req, reply, cc, opts...)

	// After invoking the RPC
	log.Printf("ClientInterceptor B - Invoked RPC method=%s; Duration=%s; Error=%v", method,
		time.Since(start), err)

	// Postprocessing
	log.Println("ClientInterceptor B - Postprocessing after invoking the RPC")
	return err
}
