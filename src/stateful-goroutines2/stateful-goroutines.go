package goroutines2

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func PrintValue() {

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = state[write.val]
				write.resp <- true
			}
		}
	}()

	var ops int64 = 0

	for i := 0; i < 10; i++ {
		go func() {
			for {
				readOp := &readOp{
					key:  rand.Intn(10),
					resp: make(chan int)}
				reads <- readOp
				<-readOp.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				writeOp := &writeOp{
					key:  rand.Intn(10),
					resp: make(chan bool),
					val:  rand.Intn(100)}
				writes <- writeOp
				<-writeOp.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("operation %d\n", atomic.LoadInt64(&ops))
}
