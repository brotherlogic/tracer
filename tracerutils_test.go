package main

import "testing"
import pb "github.com/brotherlogic/tracer/proto"

func TestLongContextCall(t *testing.T) {
	s := InitTestServer()
	s.calls = append(s.calls, &pb.ContextCall{Properties: &pb.ContextProperties{Created: 10, Died: 20}})
	l := s.getLongContextCall()

	if l.GetProperties().Created != 10 {
		t.Errorf("Wrong long call: %v", l)
	}
}
