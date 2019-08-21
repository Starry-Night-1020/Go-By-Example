package counters

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func PrintValue() {
	var uops uint64

	for i := 1; i <= 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&uops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	result := atomic.LoadUint64(&uops)
	fmt.Println("result: ", result)
}
