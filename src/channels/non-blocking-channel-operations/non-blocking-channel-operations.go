package operations

import "fmt"

func PrintValue() {
	messages := make(chan string)
	signals := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("received message: " + msg)
	default:
		fmt.Println("no message")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send message: " + msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message: " + msg)
	case signal := <-signals:
		fmt.Println("received signal: " + signal)
	default:
		fmt.Println("no activity")
	}
}
