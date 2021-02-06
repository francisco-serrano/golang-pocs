package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func main() {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_RAW, syscall.IPPROTO_ICMP)
	if err != nil {
		log.Fatal(err)
	}

	f := os.NewFile(uintptr(fd), "captureICMP")
	if f == nil {
		log.Fatal("error in os.NewFile:", err)
	}

	if err := syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_RCVBUF, 256); err != nil {
		log.Fatal(err)
	}

	for {
		buf := make([]byte, 1024)
		numRead, err := f.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("% X\n", buf[:numRead])
	}
}
