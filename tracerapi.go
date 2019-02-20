package main

import (
	"fmt"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

// Record record a trace
func (s *Server) Record(ctx context.Context, req *pb.RecordRequest) (*pb.RecordResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}

//Trace pulls out a trace
func (s *Server) Trace(ctx context.Context, req *pb.TraceRequest) (*pb.TraceResponse, error) {
	return nil, fmt.Errorf("Not implemented")
}
