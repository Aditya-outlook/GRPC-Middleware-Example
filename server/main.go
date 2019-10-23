package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"

	i "GRPC-Middleware-Example/interceptors"
	pb "GRPC-Middleware-Example/releases"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

type goReleaseService struct {
	pb.UnimplementedGoReleaseServiceServer
	releases map[string]pb.ReleaseInfo
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
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Listening on ", fmt.Sprintf("localhost:%d", *port))

	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(i.ServerInterceptorA, i.ServerInterceptorB)))

	server := grpc.NewServer(opts...)
	svc := &goReleaseService{releases: make(map[string]pb.ReleaseInfo)}

	svc.releases["1.13"] = pb.ReleaseInfo{
		Version:         "1.13",
		ReleaseDate:     "22.10.2019",
		ReleaseNotesUrl: "Latest release",
	}
	svc.releases["1.1"] = pb.ReleaseInfo{
		Version:         "1.1",
		ReleaseDate:     "21.10.2009",
		ReleaseNotesUrl: "First release",
	}

	pb.RegisterGoReleaseServiceServer(server, svc)
	// pb.RegisterGoReleaseServiceServer(server, newServer())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
