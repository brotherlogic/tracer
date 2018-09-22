package main

import pb "github.com/brotherlogic/tracer/proto"

func (s *Server) getLongContextCall() *pb.ContextCall {
	var rcall *pb.ContextCall
	longest := int64(0)
	for _, call := range s.calls {
		if call.Properties.Died > 0 && call.Properties.Created > 0 {
			took := call.Properties.Died - call.Properties.Created
			if took > longest {
				rcall = call
				longest = took
			}
		}
	}
	return rcall
}
