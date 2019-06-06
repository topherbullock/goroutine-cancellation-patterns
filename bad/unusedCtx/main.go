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

	go say(ctx, "hello")
	go say(ctx, "gophers")

	<-helpers.WaitForKeypress()
	cancel()
	fmt.Println("context cancelled")
	<-helpers.WaitForKeypress()
}

func say(ctx context.Context, message string) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(message)
		}
	}
}
