package main

import "golang.org/x/net/context"

func (s *Server) clean(ctx context.Context) {
	s.calls = nil
}
