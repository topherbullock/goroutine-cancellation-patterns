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

	go consumeWithTimeout(messages, 3*time.Second)

	<-helpers.WaitForKeypress()
	close(messages)
	<-helpers.WaitForKeypress()
}

func consumeWithTimeout(messages <-chan string, t time.Duration) {
	timer := time.NewTimer(t)
	for {
		select {
		case message, ok := <-messages:
			if ok {
				timer.Reset(t)
				fmt.Printf("- %s \n", message)
				continue
			}
			fmt.Println("Channel closed!")
			return
		case <-timer.C:
			fmt.Println("TIMEOUT!")
			return
		}
	}
}

func produce(content []string) chan string {
	messages := make(chan string)
	var sleep time.Duration = 1

	go func() {
		defer func() {
			if r := recover(); r != nil {
				return
			}
		}()
		for _, v := range content {
			messages <- v
			time.Sleep(sleep * time.Second)
			sleep = sleep + 1
		}
		close(messages)
	}()

	return messages
}
