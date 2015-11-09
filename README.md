## go-asyncserver

[![Build Status](https://travis-ci.org/hectane/go-asyncserver.svg?branch=master)](https://travis-ci.org/hectane/go-asyncserver)
[![GoDoc](https://godoc.org/github.com/hectane/go-asyncserver?status.svg)](https://godoc.org/github.com/hectane/go-asyncserver)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package provides an extremely simple HTTP server that runs asynchronously.

Although Go provides `http.ListenAndServe()`, this method blocks indefinitely. Worse yet, `http.Server` doesn't provide a simple method for stopping the server. The need for go-asyncserver arose from these problems.
