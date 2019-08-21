package epoch

import (
	"fmt"
	"time"
)

func PrintValue() {
	now := time.Now()
	fmt.Println(now)
	sec := now.Unix()
	fmt.Println(sec)
	nano := now.UnixNano()
	fmt.Println(nano)

	mills := nano / 1000000
	fmt.Println(mills)

	fmt.Println(time.Unix(sec, 0))
	fmt.Println(time.Unix(0, nano))
	fmt.Println(time.Unix(sec, nano))
}
