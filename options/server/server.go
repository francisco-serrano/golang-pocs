package main

import (
	"fmt"
	"time"
)

type server struct {
	addr    string
	timeout time.Duration
}

type option func(s *server)

func newServer(addr string, opts ...option) *server {
	srv := &server{
		addr:    addr,
		timeout: 10 * time.Second, // value by default
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func timeout(t time.Duration) option {
	return func(s *server) {
		s.timeout = t
	}
}

func main() {
	fmt.Printf("%+v\n", newServer(":8080"))
	fmt.Printf("%+v\n", newServer(":8080", timeout(5*time.Second)))
}
