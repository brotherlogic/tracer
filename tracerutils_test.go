package main

import (
	"testing"
	"time"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

func TestClean(t *testing.T) {
	s := InitTestServer()
	s.Record(context.Background(), &pb.RecordRequest{Event: &pb.Event{Id: "blah", Call: "doubleblah", Timestamp: time.Now().Unix() - 1000}})
	s.clean(context.Background())
}
