package server

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	s := New(":0")
	if err := s.Start(); err != nil {
		t.Fatal(err)
	}
	url := "http://" + s.Addr
	if _, err := http.Get(url); err != nil {
		t.Fatal(err)
	}
	s.Stop()
	if _, err := http.Get(url); err == nil {
		t.Fatal("error expected")
	}
}
