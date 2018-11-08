package main

import (
	"golang.org/x/net/context"

	pb "github.com/brotherlogic/tracer/proto"
)

func (s *Server) getLongContextCall(ctx context.Context) *pb.ContextCall {
	var rcall *pb.ContextCall
	longest := int64(0)
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()
	for _, call := range s.calls {
		if call.Properties.Died > 0 && call.Properties.Created > 0 {
			took := call.Properties.Died - call.Properties.Created
			if took > longest {
				rcall = call
				longest = took
			}
		} else {
			s.RaiseIssue(ctx, "Unfinished call", "This call is unfinished", false)
		}
	}
	return rcall
}
