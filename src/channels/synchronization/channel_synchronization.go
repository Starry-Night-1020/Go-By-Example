package synchronization

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("start")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func PrintValue() {
	done := make(chan bool, 1)
	go worker(done)
	<-done
	fmt.Println("finish")
}
