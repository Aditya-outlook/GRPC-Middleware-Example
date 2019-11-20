package main

import (
	"context"
	"flag"
	"log"

	pb "GRPC-Middleware-Example/releases"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(i.ClientInterceptorA))
	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(i.ClientInterceptorA, i.ClientInterceptorB)))

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	client := pb.NewGoReleaseServiceClient(conn)

	ctx := context.Background()

	req := pb.GetReleaseRequest{Version: "@@@@", Param2: 1, Param3: "2"}
	rsp1, err := client.GetRelease(ctx, &req)
	log.Println(rsp1, err)
	if err != nil {
		log.Fatalf("GetRelease err: %v", rsp1)
	} else {
		log.Println(rsp1)
	}
}
