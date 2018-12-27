package main

import (
	"testing"

	pb "github.com/brotherlogic/tracer/proto"
	"golang.org/x/net/context"
)

func TestLongContextCall(t *testing.T) {
	s := InitTestServer()
	s.whitelist = append(s.whitelist, "madeup")
	s.calls["madeup"] = &pb.ContextCall{Properties: &pb.ContextProperties{Created: 10, Died: 20, Origin: "madeup"}}
	l := s.getLongContextCall(context.Background())

	if l.GetProperties().Created != 10 {
		t.Errorf("Wrong long call: %v", l)
	}
}

func TestLongContextCallWithFailure(t *testing.T) {
	s := InitTestServer()
	s.calls["madeup"] = &pb.ContextCall{Properties: &pb.ContextProperties{Created: 10, Died: 20, Origin: "madeup"}}
	l := s.getLongContextCall(context.Background())

	if l != nil {
		t.Errorf("Found a context call despite whitelist")
	}
}

func TestLongContextCallWithBuild(t *testing.T) {
	s := InitTestServer()
	s.whitelist = append(s.whitelist, "madeup")
	s.calls["madeup"] = &pb.ContextCall{Milestones: []*pb.Milestone{&pb.Milestone{Label: "blah", Type: pb.Milestone_START, Timestamp: 10}, &pb.Milestone{Label: "blah", Type: pb.Milestone_END, Timestamp: 20}}, Properties: &pb.ContextProperties{Origin: "madeup"}}
	l := s.getLongContextCall(context.Background())

	if l.GetProperties().Created != 10 {
		t.Errorf("Wrong long call: %v", l)
	}
}

func TestLongContextCallWithBuildNoFinish(t *testing.T) {
	s := InitTestServer()
	s.whitelist = append(s.whitelist, "madeup")
	s.calls["madeup"] = &pb.ContextCall{Milestones: []*pb.Milestone{&pb.Milestone{Label: "blah", Type: pb.Milestone_START, Timestamp: 10}}, Properties: &pb.ContextProperties{}}
	l := s.getLongContextCall(context.Background())

	if l != nil {
		t.Errorf("Wrong long call: %v", l)
	}
}

func TestLongContextCallSkipsJumps(t *testing.T) {
	s := InitTestServer()
	s.whitelist = append(s.whitelist, "madeup")
	s.calls["madeup"] = &pb.ContextCall{Properties: &pb.ContextProperties{Created: 10, Died: 20, Origin: "madeup"}}
	s.calls["madeup2"] = &pb.ContextCall{Properties: &pb.ContextProperties{Created: 15, Died: 200, Origin: "madeup"}, Milestones: []*pb.Milestone{&pb.Milestone{Type: pb.Milestone_START_EXTERNAL, Timestamp: 16}, &pb.Milestone{Type: pb.Milestone_END_EXTERNAL, Timestamp: 199}}}
	l := s.getLongContextCall(context.Background())

	if l.GetProperties().Created != 10 {
		t.Errorf("Wrong long call: %v", l)
	}
}
