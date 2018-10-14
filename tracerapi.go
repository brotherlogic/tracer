package main

import (
	"fmt"
	"time"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	t := time.Now()

	s.callsMutex.Lock()
	call, ok := s.calls[req.Properties.Id]
	if ok {
		s.callsMutex.Unlock()
		call.Milestones = append(call.Milestones, req.Milestone)
		if req.Milestone.Type == pb.Milestone_END {
			call.Properties.Died = req.Milestone.Timestamp
		}
		if req.Milestone.Type == pb.Milestone_START {
			call.Properties.Created = req.Milestone.Timestamp
		}

	} else {
		call := &pb.ContextCall{Properties: req.Properties, Milestones: []*pb.Milestone{req.Milestone}}
		if req.Milestone.Type == pb.Milestone_START {
			call.Properties.Created = req.Milestone.Timestamp
			call.Properties.Label = req.Milestone.Origin
		}

		s.calls[req.Properties.Id] = call
		s.callsMutex.Unlock()
	}

	s.Log(fmt.Sprintf("TOOK %v", time.Now().Sub(t)))
	return &pb.RecordResponse{}, nil
}

//Trace pulls out a trace
func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	resp := &pb.TraceResponse{Calls: make([]*pb.ContextCall, 0)}
	for _, call := range s.calls {
		resp.Calls = append(resp.Calls, call)
	}

	return resp, nil
}
