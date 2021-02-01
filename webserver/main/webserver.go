package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"runtime"
	"sync/atomic"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)

	body := "The current time is:"
	fmt.Fprintf(w, `<h1 align="center">%s</h1>`, body)
	fmt.Fprintf(w, `<h2 align="center">%s</h2>\n`, t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func exampleA() {
	port := ":8080"

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", myHandler)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

var count int32

func handleAll(w http.ResponseWriter, r *http.Request) {
	atomic.AddInt32(&count, 1)
}

func getCounter(w http.ResponseWriter, r *http.Request) {
	temp := atomic.LoadInt32(&count)
	fmt.Println("Count:", temp)
	fmt.Fprintf(w, `<h1 align="center">%d</h1>`, count)
}

func exampleB() {
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	http.HandleFunc("/getCounter", getCounter)
	http.HandleFunc("/", handleAll)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func exampleC() {
	port := ":8080"

	r := http.NewServeMux()
	r.HandleFunc("/time", timeHandler)
	r.HandleFunc("/", myHandler)
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	//exampleA()
	//exampleB()
	exampleC()
}
