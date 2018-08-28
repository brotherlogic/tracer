package main

import (
	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	found := false

	for _, call := range s.calls {
		if call.Properties.Id == req.Properties.Id {
			found = true
			call.Milestones = append(call.Milestones, req.Milestone)
			if req.Milestone.Type == pb.Milestone_END {
				call.Properties.Died = req.Milestone.Timestamp
			}
			if req.Milestone.Type == pb.Milestone_START {
				call.Properties.Created = req.Milestone.Timestamp
			}
		}
	}

	if !found {
		call := &pb.ContextCall{Properties: req.Properties, Milestones: []*pb.Milestone{req.Milestone}}
		if req.Milestone.Type == pb.Milestone_START {
			call.Properties.Created = req.Milestone.Timestamp
			call.Properties.Label = req.Milestone.Origin
		}

		s.calls = append(s.calls, call)
	}

	return &pb.RecordResponse{}, nil
}

func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	resp := &pb.TraceResponse{Calls: make([]*pb.ContextCall, 0)}
	for _, call := range s.calls {
		resp.Calls = append(resp.Calls, call)
	}

	return resp, nil
}
