package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	port := ":8080"

	s, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp4", s)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	rand.Seed(time.Now().Unix())

	b := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(b)
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(string(b[0:n])) == "STOP" {
			fmt.Println("Exiting UDP Server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))

		fmt.Printf("data: %s\n", string(data))

		if _, err := conn.WriteToUDP(data, addr); err != nil {
			log.Fatal(err)
		}
	}
}
