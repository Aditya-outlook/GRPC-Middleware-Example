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
	start := time.Now()

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
}
}
	return h, err
}
}
