package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/tracer/proto"
)

func TestRecord(t *testing.T) {
	s := InitTestServer()
	a, err := s.Record(context.Background(), &pb.RecordRequest{Event: &pb.Event{Id: "test"}})
	if err != nil {
		t.Errorf("Full reject was not rejected: %v", a)
	}

	a, err = s.Record(context.Background(), &pb.RecordRequest{Event: &pb.Event{Id: "test"}})
	if err != nil {
		t.Errorf("Full reject was not rejected: %v", a)
	}

	b, err := s.Trace(context.Background(), &pb.TraceRequest{Id: "test"})
	if err != nil {
		t.Errorf("Bad Trace: %v", err)
	}

	if len(b.Traces[0].Events) != 2 {
		t.Errorf("Bad trace: %v", b.Traces[0])
	}
}
