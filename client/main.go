package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	i "GRPC-Middleware-Example/interceptors"
	pb "GRPC-Middleware-Example/releases"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()

	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	// conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(i.ClientInterceptorA))
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(i.ClientInterceptorA, i.ClientInterceptorB)))

	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	client := pb.NewGoReleaseServiceClient(conn)

	ctx := context.Background()
	rsp, err := client.ListReleases(ctx, &pb.ListReleasesRequest{})

	if err != nil {
		log.Fatalf("ListReleases err: %v", err)
	}

	releases := rsp.GetReleases()
	if len(releases) > 0 {
		fmt.Printf("Version\tRelease Date\tRelease Notes\n")
	} else {
		fmt.Println("No releases found")
	}

	for _, ri := range releases {
		fmt.Printf("%s\t%s\t%s\n",
			ri.GetVersion(),
			ri.GetReleaseDate(),
			ri.GetReleaseNotesUrl())
	}
}
