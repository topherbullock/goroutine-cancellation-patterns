package helpers

import (
	"fmt"
	"os"
)

func WaitForSignal() chan bool {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	return done
}

func WaitForKeypress() chan string {
	msg := make(chan string, 1)
	go func() {
		var input string
		for {
			input = ""
			fmt.Scanln(&input)
			msg <- input
		}
	}()
	return msg
}
