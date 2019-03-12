package main

import (
	"context"
	"testing"

	pb "github.com/brotherlogic/tracer/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	return s
}

func TestTrace(t *testing.T) {
	s := InitTestServer()
	a, err := s.Trace(context.Background(), &pb.TraceRequest{})
	if err == nil {
		t.Errorf("Full reject was not rejected: %v", a)
	}
}

func TestMark(t *testing.T) {
	s := InitTestServer()
	_, err := s.Mark(context.Background(), &pb.MarkRequest{LongRunningId: "blah"})
	if err != nil {
		t.Errorf("Full reject was not rejected: %v", err)
	}
}
