package async

import (
	"crypto/tls"
	"net"
	"net/http"
)

// Asynchronous HTTP server that can be started and stopped asynchronously.
type Server struct {
	http.Server
	listener net.Listener
	stopped  chan bool
}

// Create a new server instance. Note that Start() must be called before the
// server will begin accepting new connections.
func New(addr string) *Server {
	s := &Server{
		stopped: make(chan bool),
	}
	s.Addr = addr
	return s
}

// Start the server.
func (s *Server) Start() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.Addr = l.Addr().String()
	if s.TLSConfig != nil {
		l = tls.NewListener(l, s.TLSConfig)
	}
	s.listener = l
	go func() {
		s.Serve(s.listener)
		s.stopped <- true
	}()
	return nil
}

// Stop the server. This method blocks until the server is stopped.
func (s *Server) Stop() {
	s.listener.Close()
	<-s.stopped
}
