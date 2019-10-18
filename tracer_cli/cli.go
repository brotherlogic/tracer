package main

import (
	"fmt"
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
	host, port, err := utils.Resolve("tracer", "tracer-cli")
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
	case "trace":
		val, err := client.Trace(ctx, &pb.TraceRequest{Id: os.Args[2]})
		if err != nil {
			log.Fatalf("Failed on trace request: %v", err)
		}
		if len(val.Traces) == 0 {
			fmt.Printf("No traces match!")
		}
		if len(val.Traces[0].Events) == 0 {
			fmt.Printf("No events in trace")
		}
		for _, event := range val.Traces[0].Events {
			fmt.Printf("Event: %v\n", event)
		}
	case "default":
		fmt.Printf("Unknown command\n")
	}
}
