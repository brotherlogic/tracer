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

}
