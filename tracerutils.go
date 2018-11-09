package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/tracer/proto"
)

func (s *Server) getLongContextCall(ctx context.Context) *pb.ContextCall {
	var rcall *pb.ContextCall
	longest := int64(0)
	s.callsMutex.Lock()
	defer s.callsMutex.Unlock()
	for _, call := range s.calls {
		if call.Properties.Died == 0 || call.Properties.Created == 0 {
			for _, m := range call.Milestones {
				if m != nil {
					if m.Type == pb.Milestone_START {
						call.Properties.Created = m.Timestamp
					}
					if m.Type == pb.Milestone_END {
						call.Properties.Died = m.Timestamp
					}
				}
			}
		}

		if call.Properties.Died > 0 && call.Properties.Created > 0 {
			took := call.Properties.Died - call.Properties.Created
			if took > longest {
				rcall = call
				longest = took
			}
		} else {
			minTime := time.Now().Unix()
			for _, m := range call.Milestones {
				if m != nil {
					if m.Timestamp < minTime {
						minTime = m.Timestamp
					}
				}
			}
			if time.Now().Sub(time.Unix(minTime, 0)) > time.Hour {
				s.RaiseIssue(ctx, "Unfinished call", fmt.Sprintf("The call for %v is unfinished", call.Properties.Label), false)
			}
		}
	}
	return rcall
}
