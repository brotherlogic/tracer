package main

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/tracer/proto"
)

func TestClean(t *testing.T) {
	s := InitTestServer()
	s.Record(context.Background(), &pb.RecordRequest{Event: &pb.Event{Id: "blah", Call: "doubleblah", Timestamp: time.Now().Unix() - 1000}})
	s.Record(context.Background(), &pb.RecordRequest{Event: &pb.Event{Id: "blah", Call: "doubleblah2", Timestamp: time.Now().Unix() - 500}})
	s.Mark(context.Background(), &pb.MarkRequest{LongRunningId: "blah", Origin: "recordmatcher"})
	s.clean(context.Background())
}
