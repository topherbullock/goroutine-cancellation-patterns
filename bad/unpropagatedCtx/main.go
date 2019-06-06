package main

import (
	"context"
	"fmt"
	"time"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

var rootContext = context.Background()

func main() {
	ctx, cancel := context.WithCancel(rootContext)

	go multiSay(ctx, []string{"hello", "gophers"})

	<-helpers.WaitForKeypress()
	cancel()
	<-helpers.WaitForKeypress()
}

func multiSay(ctx context.Context, messages []string) {
	for _, msg := range messages {
		go say(msg)
	}

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("---")
		case <-ctx.Done():
			fmt.Println("context cancelled")
			return
		}
	}
}

func say(message string) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(message)
		}
	}
}
