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
	<-helpers.WaitForSignal()
}

func multiSay(ctx context.Context, messages []string) {
	for _, msg := range messages {
		go say(ctx, msg)
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

func say(ctx context.Context, message string) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(message)
		case <-ctx.Done():
			fmt.Printf("'%s' context cancelled\n", message)
			return
		}
	}
}
