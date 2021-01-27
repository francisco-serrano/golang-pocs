package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Printf("handleSignal() Caught: %s %d\n", signal, signal)
}

func main() {
	fmt.Println("PID", os.Getpid())

	s := make(chan os.Signal, 1)

	signal.Notify(s)

	go func() {
		for {
			sig := <-s

			switch sig {
			case os.Interrupt:
				fmt.Printf("Caught: %s %d\n", sig, sig)
			default:
				handleSignal(sig)
				return
			}
		}
	}()

	for {
		fmt.Printf(".")

		time.Sleep(20 * time.Second)
	}
}
