package main

import (
	"fmt"

	"github.com/topherbullock/goroutine-cancellation-patterns/helpers"
)

func main() {
	go Ham()

	<-helpers.WaitForSignal()
}

func Ham() {
	fmt.Print(".")
	go Ham()
	fmt.Print("🐹")
}
