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
	finishedCalls := make(map[string]int)
	unfinishedCalls := make(map[string]int)

	for _, call := range s.calls {
		if call.Properties.Died == 0 || call.Properties.Created == 0 {
			for _, m := range call.Milestones {
				if m != nil {
					if m.Type == pb.Milestone_START {
						call.Properties.Created = m.Timestamp
						if call.Properties.Label == "" {
							call.Properties.Label = m.Label
						}
					}
					if m.Type == pb.Milestone_END {
						call.Properties.Died = m.Timestamp
					}
				}
			}
		}

		if call.Properties.Died > 0 && call.Properties.Created > 0 && !call.Properties.Delivered {
			finishedCalls[call.Properties.Label]++
			took := call.Properties.Died - call.Properties.Created
			if took > longest {
				rcall = call
				longest = took
			}
		} else {
			unfinishedCalls[call.Properties.Label]++
		}
	}

	// Look for unfinished calls
	for _, call := range s.calls {
		if unfinishedCalls[call.Properties.Label] > 0 && finishedCalls[call.Properties.Label] == 0 {
			minTime := time.Now().UnixNano()
			for _, m := range call.Milestones {
				if m != nil {
					if m.Timestamp < minTime {
						minTime = m.Timestamp
					}
				}
			}

			if time.Now().Sub(time.Unix(0, minTime)) > time.Minute*5 {
				betterLabel := ""
				milestones := ""
				for _, m := range call.Milestones {
					if len(m.Label) > len(betterLabel) {
						betterLabel = m.Label
					}
					milestones += fmt.Sprintf("%v, ", m.Type)
				}
				s.RaiseIssue(ctx, "Unfinished call", fmt.Sprintf("The call for %v from %v is unfinished (%v milestones = %v) -> %v", call.Properties.Label, call.Properties.Origin, len(call.Milestones), milestones, betterLabel), false)
			}
		}
	}

	return rcall
}
