package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func f(n int) int {
	fn := make(map[int]int)

	for i := 0; i <= n; i++ {
		var f int
		if i <= 2 {
			f = 1
		} else {
			f = fn[i-1] + fn[i-2]
		}

		fn[i] = f
	}

	return fn[n]
}

func handleConnection(c net.Conn) {
	for {
		data, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		temp := strings.TrimSpace(data)
		if temp == "STOP" {
			break
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}

		c.Write([]byte(fibo))
	}

	time.Sleep(5 * time.Second)
	c.Close()
}

func main() {
	port := ":8080"

	l, err := net.Listen("tcp4", port)
	if err != nil {
		log.Fatal(err)
	}

	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(c)
	}
}
