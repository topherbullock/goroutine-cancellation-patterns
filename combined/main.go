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

	messages := produce([]string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
	})
	go consume(ctx, messages)

	<-helpers.WaitForKeypress()
	cancel()
	time.Sleep(10 * time.Millisecond)
}

func consume(ctx context.Context, messages <-chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Cancelled by user")
			return
		case message, ok := <-messages:
			if ok {
				fmt.Printf("- %s \n", message)
				continue
			}
			fmt.Println("Done!")
			return
		}
	}

}

func produce(content []string) chan string {
	messages := make(chan string)
	go func() {
		for _, v := range content {
			messages <- v
			time.Sleep(5 * time.Second)
		}
		close(messages)
	}()
	return messages
}
