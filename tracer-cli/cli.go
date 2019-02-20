package main

import (
	"log"
	"os"
	"strconv"

	"github.com/brotherlogic/goserver/utils"
	"google.golang.org/grpc"

	pb "github.com/brotherlogic/tracer/proto"

	//Needed to pull in gzip encoding init
	_ "google.golang.org/grpc/encoding/gzip"
)

func main() {
	host, port, err := utils.Resolve("tracer")
	if err != nil {
		log.Fatalf("Unable to reach tracer: %v", err)
	}
	conn, err := grpc.Dial(host+":"+strconv.Itoa(int(port)), grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("Unable to dial: %v", err)
	}

	client := pb.NewTracerServiceClient(conn)
	ctx, cancel := utils.BuildContext("tracercli-"+os.Args[1], "tracer")
	defer cancel()

	switch os.Args[1] {
	default:
		client.Trace(ctx, &pb.TraceRequest{})
	}
}
