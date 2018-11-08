package main

import (
	"testing"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

func TestLongContextCall(t *testing.T) {
	s := InitTestServer()
	s.calls["madeup"] = &pb.ContextCall{Properties: &pb.ContextProperties{Created: 10, Died: 20}}
	l := s.getLongContextCall(context.Background())

	if l.GetProperties().Created != 10 {
		t.Errorf("Wrong long call: %v", l)
	}
}
