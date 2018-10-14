package main

import (
	"fmt"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	s.Log(fmt.Sprintf("RUNNING %v - %v", req.Properties.Origin, req.Properties.Label))

	s.callsMutex.Lock()
	call, ok := s.calls[req.Properties.Id]
	if ok {
		s.callsMutex.Unlock()
		call.Milestones = append(call.Milestones, req.Milestone)
	} else {
		call := &pb.ContextCall{Properties: req.Properties, Milestones: []*pb.Milestone{req.Milestone}}
		s.calls[req.Properties.Id] = call
		s.callsMutex.Unlock()
	}

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
