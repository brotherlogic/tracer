package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	pb "github.com/brotherlogic/tracer/proto"
)

func (s *Server) inWhitelist(name string) bool {
	for _, n := range s.whitelist {
		if name == n {
			return true
		}
	}

	return false
}

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

		if call.Properties.Died > 0 && call.Properties.Created > 0 {
			if !call.Properties.Delivered {
				finishedCalls[call.Properties.Label]++

				call.Properties.Length = call.Properties.Died - call.Properties.Created
				sTime := int64(0)
				for _, milestone := range call.Milestones {
					if milestone.Type == pb.Milestone_START_EXTERNAL {
						sTime = milestone.Timestamp
					} else if milestone.Type == pb.Milestone_END_EXTERNAL {
						call.Properties.Length -= milestone.Timestamp - sTime
						sTime = 0
					}
				}

				if call.Properties.Length > longest && s.inWhitelist(call.Properties.Origin) {
					rcall = call
					longest = call.Properties.Length
				}
			}
		} else {
			unfinishedCalls[call.Properties.Label]++
		}
	}

	// Look for unfinished calls
	for _, call := range s.calls {
		if unfinishedCalls[call.Properties.Label] > 0 && finishedCalls[call.Properties.Label] == 0 {
			if call.Properties.Died == 0 || call.Properties.Created == 0 {
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
						if m != nil && len(m.Label) > len(betterLabel) {
							betterLabel = m.Label
						}
						milestones += fmt.Sprintf("%v, ", m.Type)
					}
					if call.Properties.Died == 0 || call.Properties.Created == 0 {
						s.RaiseIssue(ctx, "Unfinished call", fmt.Sprintf("The call for %v from %v is unfinished (%v milestones = %v) -> %v (%v and %v)", call.Properties.Label, call.Properties.Origin, len(call.Milestones), milestones, betterLabel, call.Properties.Created, call.Properties.Died), false)
					}
				}
			}
		}
	}

	return rcall
}
