package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/go-grpc-middleware"

	pb "GRPC-Middleware-Example/releases"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type goReleaseService struct {
	pb.UnimplementedGoReleaseServiceServer
	releases map[string]pb.ReleaseInfo
}

func (g *goReleaseService) GetReleaseInfo(ctx context.Context, req *pb.ReleaseInfoRequest) (*pb.ReleaseInfo, error) {
	ri, ok := g.releases[req.GetVersion()]

	if !ok {
		return nil, status.Errorf(codes.NotFound, "release version %s is not found", req.GetVersion())
	}

	// Success
	return &pb.ReleaseInfo{
		Version:         req.GetVersion(),
		ReleaseDate:     ri.ReleaseDate,
		ReleaseNotesUrl: ri.ReleaseNotesUrl,
	}, nil
}

func (g *goReleaseService) ListReleases(ctx context.Context, r *pb.ListReleasesRequest) (*pb.ListReleasesResponse, error) {
	var releases []*pb.ReleaseInfo

	// build slice with all the available releases
	for k, v := range g.releases {
		ri := &pb.ReleaseInfo{
			Version:         k,
			ReleaseDate:     v.ReleaseDate,
			ReleaseNotesUrl: v.ReleaseNotesUrl,
		}

		releases = append(releases, ri)
	}

	return &pb.ListReleasesResponse{
		Releases: releases,
	}, nil
}

func main() {
	// code redacted

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", fmt.Sprintf("localhost:%d", *port))

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(serverInterceptorA, serverInterceptorB)))

	server := grpc.NewServer(opts...)
	svc := &goReleaseService{releases: make(map[string]pb.ReleaseInfo)}
	svc.releases["1.1"] = pb.ReleaseInfo{
		Version:         "1.1",
		ReleaseDate:     "22.10.2019",
		ReleaseNotesUrl: "First release",
	}
	pb.RegisterGoReleaseServiceServer(server, svc)
	// pb.RegisterGoReleaseServiceServer(server, newServer())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Example server unary interceptor function
func serverInterceptorA(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Println("A - start: ", start)

	// Calls the handler
	h, err := handler(ctx, req)

	log.Println("A - End: ", time.Now())
	// Logging with grpclog (grpclog.LoggerV2)
	log.Printf("A - Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}

func serverInterceptorB(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	log.Println("B - start: ", start)

	// Calls the handler
	h, err := handler(ctx, req)

	log.Println("B - End: ", time.Now())
	// Logging with grpclog (grpclog.LoggerV2)
	log.Printf("B - Request - Method:%s\tDuration:%s\tError:%v\n",
		info.FullMethod,
		time.Since(start),
		err)

	return h, err
}
