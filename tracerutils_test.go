package main

import (
	"testing"

	"golang.org/x/net/context"
)

func TestClean(t *testing.T) {
	s := InitTestServer()
	s.clean(context.Background())
}
