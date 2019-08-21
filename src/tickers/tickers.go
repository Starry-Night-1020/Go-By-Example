package tickers

import (
	"fmt"
	"time"
)

func PrintValue() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("time: ", t)
		}
	}()

	time.Sleep(time.Second * 2)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
