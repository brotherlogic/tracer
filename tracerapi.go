package main

import (
	"fmt"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()

	val, ok := s.calls[req.Event.Id]
	if !ok {
		s.calls[req.Event.Id] = &pb.Trace{Events: []*pb.Event{req.Event}}
	} else {
		val.Events = append(val.Events, req.Event)
	}

	return &pb.RecordResponse{}, nil
}

//Trace pulls out a trace
func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()

	val, ok := s.calls[req.Id]
	if ok {
		return &pb.TraceResponse{Traces: []*pb.Trace{val}}, nil
	}

	return nil, fmt.Errorf("Unable to find trace with that id: %v", req.Id)
}
