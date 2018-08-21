package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/brotherlogic/tracer/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	return s
}

func TestInvertedCalls(t *testing.T) {
	s := InitTestServer()
	sTime := time.Now().Unix()

	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest"}, Milestone: &pb.Milestone{Type: pb.Milestone_MARKER, Timestamp: sTime + 2}})
	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest"}, Milestone: &pb.Milestone{Type: pb.Milestone_START, Timestamp: sTime}})
	s.Record(context.Background(), &pb.RecordRequest{Properties: &pb.ContextProperties{Id: "123", Label: "RunTest"}, Milestone: &pb.Milestone{Type: pb.Milestone_END, Timestamp: sTime + 3}})

	resp, err := s.Trace(context.Background(), &pb.TraceRequest{Label: "RunTest"})

	if err != nil {
		t.Fatalf("Error running trace: %v", err)
	}

	if len(resp.Calls) != 1 || len(resp.Calls[0].Milestones) != 3 {
		t.Fatalf("Wrong number of calls/milestones: %v", resp)
	}

	if resp.Calls[0].Properties.Created != sTime {
		t.Errorf("Bad start time: %v", resp.Calls[0])
	}
}
