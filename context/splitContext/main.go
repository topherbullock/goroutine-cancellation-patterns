package main

import (
	"context"
	"fmt"
	"time"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

var rootContext = context.Background()

func main() {
	d := time.Now().Add(5 * time.Second)
	deadlinedCtx, cancel := context.WithDeadline(rootContext, d)
	defer cancel()

	cancellableCtx, cancel := context.WithCancel(rootContext)

	go say(deadlinedCtx, "hello")
	go say(cancellableCtx, "gophers")

	<-helpers.WaitForKeypress()
	cancel()
	<-helpers.WaitForSignal()
}

func say(ctx context.Context, message string) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(message)
		case <-ctx.Done():
			fmt.Printf("'%s' context done\n", message)
			return
		}
	}
}
