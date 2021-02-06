package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	address := "localhost:8080"

	s, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")

		text, _ := reader.ReadString('\n')

		data := []byte(text + "\n")

		c.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP Client!")
			return
		}

		b := make([]byte, 1024)

		n, _, err := c.ReadFromUDP(b)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Reply: %s\n", string(b[0:n]))
	}
}
