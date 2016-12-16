package server

import (
	"crypto/tls"
	"net"
	"net/http"
)

// Asynchronous HTTP server that can be started and stopped asynchronously.
type AsyncServer struct {
	http.Server
	listener net.Listener
	stopped  chan bool
}

// Create a new server instance. Note that Start() must be called before the
// server will begin accepting new connections.
func New(addr string) *AsyncServer {
	a := &AsyncServer{
		stopped: make(chan bool),
	}
	a.Addr = addr
	return a
}

// Start the server.
func (a *AsyncServer) Start() error {
	l, err := net.Listen("tcp", a.Addr)
	if err != nil {
		return err
	}
	a.Addr = l.Addr().String()
	if a.TLSConfig != nil {
		l = tls.NewListener(l, a.TLSConfig)
	}
	a.listener = l
	go func() {
		a.Serve(a.listener)
		close(a.stopped)
	}()
	return nil
}

// Stop the server. This method blocks until the server is stopped.
func (a *AsyncServer) Stop() {
	a.listener.Close()
	<-a.stopped
}
