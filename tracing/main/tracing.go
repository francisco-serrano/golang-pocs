package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("-----")
}

func exampleA() {
	f, err := os.Create("/tmp/traceFile.out")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}

	defer trace.Stop()

	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 3; i++ {
		if s := make([]byte, 50000000); s == nil {
			fmt.Println("Operation failed!")
		}
	}

	printStats(mem)

	for i := 0; i < 5; i++ {
		if s := make([]byte, 100000000); s == nil {
			fmt.Println("Operation failed!")
		}

		time.Sleep(10 * time.Millisecond)
	}

	printStats(mem)
}

func exampleB() {
	url := "https://www.google.com"

	client := http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	trace := &httptrace.ClientTrace{
		DNSDone: func(info httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", info)
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("Dial start")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("Dial done")
		},
		TLSHandshakeStart: func() {
			fmt.Println("TLS Handshake Start")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			fmt.Printf("Handshake State: %+v\n", state)
		},
		GotConn: func(info httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", info)
		},
		WroteHeaders: func() {
			fmt.Println("Wrote headers")
		},
		GotFirstResponseByte: func() {
			fmt.Println("First response byte!")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	fmt.Println("Requesting data from server!")

	if _, err := http.DefaultTransport.RoundTrip(req); err != nil {
		fmt.Println(err)
		return
	}

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, response.Body)
}

func main() {
	//exampleA()
	exampleB()
}
