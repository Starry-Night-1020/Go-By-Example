package channels

import "fmt"

func PrintValue() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println("message: " + msg)
}
