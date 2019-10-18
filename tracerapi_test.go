package main

import (
	"testing"

	"github.com/brotherlogic/keystore/client"
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/tracer/proto"
)

func InitTestServer() *Server {
	s := Init()
	s.SkipLog = true
	s.GoServer.KSclient = *keystoreclient.GetTestClient("./testing")
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
	_, err := s.Mark(context.Background(), &pb.MarkRequest{LongRunningId: "blah", Origin: "recordmatcher"})
	if err != nil {
		t.Errorf("Full reject was not rejected: %v", err)
	}
}
