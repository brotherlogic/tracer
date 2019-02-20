package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/tracer/proto"
)

func TestSimpleFlow(t *testing.T) {
	s := InitTestServer()

	s.Record(context.Background(), &pb.RecordRequest{})
	resp, err := s.Trace(context.Background(), &pb.TraceRequest{})

	if err == nil {
		t.Errorf("Error running trace: %v", resp)
	}
}
