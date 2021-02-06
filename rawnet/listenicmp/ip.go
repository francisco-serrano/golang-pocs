package main

import (
	"fmt"
	"log"
	"net"
)

func exampleA() {
	addr, err := net.ResolveIPAddr("ip4", "127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenIP("ip4:icmp", addr)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFrom(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("% X\n", buffer[0:n])
}

func main() {
	exampleA()
}
