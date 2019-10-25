package main

import (
	"fmt"
	"reflect"
	"time"

	"golang.org/x/net/context"
)

func (s *Server) clean(ctx context.Context) error {
	process := s.calls
	s.calls = nil

	for _, events := range process {

		for _, marked := range s.markedIds {
			if events.Events[0].Id == marked.LongRunningId {
				s.message = fmt.Sprintf("Mark = %v -> %v", marked.LongRunningId, events.Events[0].Id)
				list := fmt.Sprintf("%v/%v %v %v %v [%v]", events.Events[0].Server, events.Events[0].Binary, events.Events[0].Timestamp, events.Events[0].Call, time.Millisecond*time.Duration(marked.RunningTimeInMs), marked.Request)
				for _, ev := range events.Events[1:] {
					list += "\n" + fmt.Sprintf("%v/%v %v %v", ev.Server, ev.Binary, time.Millisecond*time.Duration(ev.Timestamp-events.Events[0].Timestamp), ev.Call)
				}

				list += fmt.Sprintf("\nGenerated from %v", marked.LongRunningId)

				s.RaiseIssue(ctx, "Long Running Trace", list, false)
				s.config.LastMarkSent = time.Now().Unix()
				s.save(ctx)
			}
		}

		eventStart := ""
		times := time.Now().UnixNano()

		if events.Events != nil {
			for _, ev := range events.Events {
				if ev != nil && ev.Timestamp < times {
					times = ev.Timestamp
					eventStart = ev.Call
				}
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

	return nil
}
