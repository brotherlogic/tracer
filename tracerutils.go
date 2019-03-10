package main

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

func (s *Server) clean(ctx context.Context) {
	process := s.calls
	s.calls = nil

	for _, events := range process {
		eventStart := ""
		times := time.Now().Unix()

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
}
