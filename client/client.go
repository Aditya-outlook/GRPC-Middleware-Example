package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"
	"strconv"

	pb "GRPC-Middleware-Example/releases"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", "127.0.0.1:10000", "The server address in the format of host:port")
)

// ByVersion sort helper for credentials
type ByVersion []*pb.ReleaseInfo

func (a ByVersion) Len() int      { return len(a) }
func (a ByVersion) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a ByVersion) Less(i, j int) bool {
	aiv, _ := strconv.Atoi(a[i].Version)
	ajv, _ := strconv.Atoi(a[j].Version)
	return aiv < ajv
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
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
		sort.Sort(ByVersion(releases))
		// sort.Sort(releases)

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
