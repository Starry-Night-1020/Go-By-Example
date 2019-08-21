package directions

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func PrintValue() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, "receives message")
	go pong(pings, pongs)

	msg := <-pongs
	fmt.Println(msg)
}
