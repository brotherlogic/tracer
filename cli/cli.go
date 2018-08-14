package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	switch os.Args[1] {
	case "trace":
		traceFlags := flag.NewFlagSet("List", flag.ExitOnError)
		var label = traceFlags.String("label", "", "The label to trace")
		if err := traceFlags.Parse(os.Args[2:]); err == nil {
			r := &pb.TraceRequest{Label: *label}
			list, err := client.Trace(ctx, r)
			if err == nil {
				if len(list.Calls) == 0 {
					fmt.Printf("No traces found!\n")
				}
				for _, call := range list.Calls {
					log.Printf("%v", call.Properties)
					fmt.Printf("%v [%v] \n", call.Properties.Label, (call.Properties.Died-call.Properties.Created)/1000000)
					for _, m := range call.Milestones {
						fmt.Printf("  [%v] - %v (%v)\n", (m.Timestamp-call.Properties.Created)/1000000, m.Label, m.Type)
					}
				}
			} else {
				fmt.Printf("ERROR: %v\n", err)
			}
		}
	}

}
