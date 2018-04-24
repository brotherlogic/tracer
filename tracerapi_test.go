package main

import "testing"

func InitTestServer() *Server {
	s := Init()
	return s
}

func TestBlank(t *testing.T) {
	blank()
}
