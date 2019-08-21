package limiting

import (
	"fmt"
	"time"
)

func PrintValue() {
	requests := make(chan int, 5)

	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	for r := range requests {
		<-limiter
		fmt.Println("request: ", r, time.Now())
	}
}
