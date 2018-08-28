package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/tracer/proto"
)

func TestSimpleFlow(t *testing.T) {
	s := InitTestServer()

	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest", Origin: "testorigin"}, Milestone: &pb.Milestone{Type: pb.Milestone_START, Timestamp: time.Now().Unix()}})
	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest"}, Milestone: &pb.Milestone{Type: pb.Milestone_MARKER, Timestamp: time.Now().Unix() + 2}})
	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest"}, Milestone: &pb.Milestone{Type: pb.Milestone_END, Timestamp: time.Now().Unix() + 3}})

	resp, err := s.Trace(context.Background(), &pb.TraceRequest{Label: "RunTest"})

	if err != nil {
		t.Fatalf("Error running trace: %v", err)
	}

	if len(resp.Calls) != 1 || len(resp.Calls[0].Milestones) != 3 {
		t.Fatalf("Wrong number of calls/milestones: %v", resp)
	}

	if resp.Calls[0].Properties.Origin != "testorigin" {
		t.Errorf("Wrong Origin: %v", resp.Calls[0].Properties)
	}

}
