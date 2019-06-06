package main

import (
	"fmt"
	"time"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

func main() {
	done := make(chan bool, 0)
	go waitForClose(done, "hey")

	<-helpers.WaitForKeypress()
	close(done)
	<-helpers.WaitForKeypress()
}

func waitForClose(done chan bool, message string) {
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println(message)
		case <-done:
			fmt.Printf("'%s' channel closed\n", message)
			return
		}
	}
}
