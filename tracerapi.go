package main

import (
	"fmt"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	for _, entry := range s.calls {
		if entry.Events[0].Id == req.Event.Id {
			entry.Events = append(entry.Events, req.Event)
			return &pb.RecordResponse{}, nil
		}
	}

	s.calls = append(s.calls, &pb.Trace{Events: []*pb.Event{req.Event}})
	return &pb.RecordResponse{}, nil
}

//Trace pulls out a trace
func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	for _, entry := range s.calls {
		if entry.Events[0].Id == req.Id {
			return &pb.TraceResponse{Traces: []*pb.Trace{entry}}, nil
		}
	}

	return nil, fmt.Errorf("Unable to find trace with that id: %v", req.Id)
}

//Mark marks a trace
func (s *Server) Mark(ctx context.Context, req *pb.MarkRequest) (*pb.MarkResponse, error) {
	s.markedIds = append(s.markedIds, req.LongRunningId)
	return &pb.MarkResponse{}, nil
}
