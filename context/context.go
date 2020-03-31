package context

import (
	"context"
	"fmt"
	"log"
	"time"
)

func Run() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	mySleepAndTalk(ctx, 5*time.Second, "hello")
}

func mySleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
