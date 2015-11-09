## go-asyncserver

[![Build Status](https://travis-ci.org/hectane/go-asyncserver.svg?branch=master)](https://travis-ci.org/hectane/go-asyncserver)
[![GoDoc](https://godoc.org/github.com/hectane/go-asyncserver?status.svg)](https://godoc.org/github.com/hectane/go-asyncserver)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an extremely simple HTTP server that runs asynchronously.

Although Go provides `http.ListenAndServe()`, this method blocks indefinitely. Worse yet, `http.Server` doesn't provide a simple method for stopping the server. The need for go-asyncserver arose from these problems.

### Example

The following example demonstrates the creation of a simple asynchronous HTTP server:

    import "github.com/hectane/go-asyncserver"

    // "0" instructs the OS to select a free port
    s := server.New(":0")
    
    // The server doesn't actually begin accepting connections
    // until the Start() method is called
    if err := s.Start(); err != nil {
        panic(err)
    }

    // The server will now accept connections at s.Addr
    // ...do stuff here...

    // Stop the server
    s.Stop()
