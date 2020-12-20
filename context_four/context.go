package context_four

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Run() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("running goroutine", i)

			for {
				select {
				case <-ctx.Done():
					fmt.Printf("context timeout for goroutine %d of %d\n", i, runtime.NumGoroutine())
					return
				default:
					fmt.Printf("still running goroutine %d of %d\n", i, runtime.NumGoroutine())
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(5 * time.Second)

	fmt.Println("calling cancel function")

	cancel()

	time.Sleep(5 * time.Second)
}
