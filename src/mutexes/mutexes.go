package mutexes

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func PrintValue() {
	state := make(map[int]int)
	mutex := &sync.Mutex{}

	var ops int64 = 0

	for i := 0; i < 10; i++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("ops: ", ops)
	fmt.Println("state: ", state)
}
