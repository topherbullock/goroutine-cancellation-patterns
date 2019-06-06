package main

import (
	"fmt"
	"time"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

func main() {
	messages := produce([]string{
		"Don't communicate by sharing memory, share memory by communicating.",
		"Concurrency is not parallelism.",
		"Channels orchestrate; mutexes serialize.",
		"The bigger the interface, the weaker the abstraction.",
	})
	go consume(messages)
	<-helpers.WaitForKeypress()
	close(messages)
	<-helpers.WaitForKeypress()
}

func consume(messages <-chan string) {
	for {
		select {
		case message, ok := <-messages:
			if ok {
				fmt.Printf("- %s \n", message)
				continue
			}
			fmt.Println("Channel closed!")
			return
		}
	}

}

func produce(content []string) chan string {
	messages := make(chan string)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		for _, v := range content {
			messages <- v
			time.Sleep(1 * time.Second)
		}
		close(messages)
	}()
	return messages
}
