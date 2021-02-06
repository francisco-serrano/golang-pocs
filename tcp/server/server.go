package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func exampleA() {
	port := ":8080"

	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(netData) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("->", netData)

		t := time.Now()

		myTime := t.Format(time.RFC3339) + "\n"

		c.Write([]byte(myTime))

	}
}

func exampleB() {
	address := "localhost:8080"

	s, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		log.Fatal(err)
	}

	b := make([]byte, 1024)

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		n, err := conn.Read(b)
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(string(b[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}

		fmt.Printf("> %s\n", string(b[0:n-1]))

		if _, err := conn.Write(b); err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	//exampleA()
	exampleB()
}
