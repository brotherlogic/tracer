package main

import (
	"fmt"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	if req.Event == nil {
		return nil, fmt.Errorf("You sent an empty event")
	}
	for _, entry := range s.calls {
		if entry != nil && req.Event != nil && entry.Events != nil {
			if len(entry.Events) > 0 &&
				entry.Events[0].Id == req.Event.Id {
				entry.Events = append(entry.Events, req.Event)
				return &pb.RecordResponse{}, nil
			}
		}
	}

	s.calls = append(s.calls, &pb.Trace{Events: []*pb.Event{req.Event}})
	return &pb.RecordResponse{}, nil
}

//Trace pulls out a trace
func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	traces := []*pb.Trace{}
	for _, entry := range s.calls {
		if entry.Events[0].Id == req.Id {
			return &pb.TraceResponse{Traces: []*pb.Trace{entry}}, nil
		}
		found := false
		for _, ev := range entry.Events {
			if ev.GetCall() == req.GetCall() {
				found = true
			}
		}

		if found {
			traces = append(traces, entry)
		}
	}

	return &pb.TraceResponse{Traces: traces}, nil
}

//Mark marks a trace
func (s *Server) Mark(ctx context.Context, req *pb.MarkRequest) (*pb.MarkResponse, error) {
	s.markedIds = append(s.markedIds, req)
	return &pb.MarkResponse{}, nil
}
