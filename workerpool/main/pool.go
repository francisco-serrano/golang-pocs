package main

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(wg *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{c, square}

		data <- output

		time.Sleep(time.Second)
	}

	wg.Done()
}

func makeWorkerPool(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(&wg)
	}

	wg.Wait()
	close(data)
}

func create(n int) {
	for i := 0; i < n; i++ {
		clients <- Client{i, i}
	}

	close(clients)
}

func main() {
	fmt.Println("Capacity of clients:", cap(clients))
	fmt.Println("Capacity of data:", cap(data))

	nJobs := 20
	nWorkers := 3

	go create(nJobs)

	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}

		finished <- true
	}()

	makeWorkerPool(nWorkers)

	fmt.Printf(": %v\n", <-finished)
}
