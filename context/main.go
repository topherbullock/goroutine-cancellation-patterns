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
	<-helpers.WaitForKeypress()
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
