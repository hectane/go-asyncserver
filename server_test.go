package server

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	a := New("127.0.0.1:0")
	if err := a.Start(); err != nil {
		t.Fatal(err)
	}
	url := "http://" + a.Addr
	if _, err := http.Get(url); err != nil {
		t.Fatal(err)
	}
	a.Stop()
	if _, err := http.Get(url); err == nil {
		t.Fatal("error expected")
	}
}
