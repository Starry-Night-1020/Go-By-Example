package timers

import (
	"fmt"
	"time"
)

func PrintValue() {
	time1 := time.NewTimer(time.Second * 2)

	<-time1.C
	fmt.Println("time1 expired")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("time2 expired")
	}()
	time2.Stop()
	fmt.Println("time2 stoped")

}
