package main

import (
	"fmt"
	"time"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

func main() {
	go sayHello()
	<-helpers.WaitForSignal()
}

func sayHello() {
	ticker := time.NewTicker(1 * time.Second)

	t := time.NewTimer(1 * time.Second)
	t.Reset(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("Hello!")
		}
	}
}
