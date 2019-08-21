package select_

import (
	"fmt"
	"time"
)

func PrintValue() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1: " + msg1)
		case msg2 := <-c2:
			fmt.Println("msg2: " + msg2)
		}
	}
}
