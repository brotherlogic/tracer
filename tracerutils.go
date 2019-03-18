package main

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/net/context"
)

func (s *Server) clean(ctx context.Context) {
	process := s.calls
	s.calls = nil

	for _, events := range process {

		for _, marked := range s.markedIds {
			if events.Events[0].Id == marked.LongRunningId {
				list := events.Events[0].Call
				for _, ev := range events.Events[1:] {
					list += "\n" + fmt.Sprintf("%v %v", ev.Timestamp, ev.Call)
				}

				s.RaiseIssue(ctx, "Long Running Trace", list, false)
				s.config.LastMarkSent = time.Now().Unix()
				s.save(ctx)
			}
		}

		eventStart := ""
		times := time.Now().UnixNano()

		for _, ev := range events.Events {
			if ev.Timestamp < times {
				times = ev.Timestamp
				eventStart = ev.Call
			}
		}

		s.counts[eventStart]++
	}

	most := ""
	mostCalls := 0
	allCalls := 0
	for c, count := range s.counts {
		allCalls += count
		if count > mostCalls {
			most = c
			mostCalls = count
		}
	}

	s.Log(fmt.Sprintf("Most calls: %v", mostCalls))
	s.allCalls = int64(allCalls)
	s.mostCalled = most

	s.Log(fmt.Sprintf("Calls: %v", reflect.ValueOf(s.counts).MapKeys()))
}
