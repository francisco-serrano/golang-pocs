package main

import (
	"fmt"
	"log"
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

func main() {
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
