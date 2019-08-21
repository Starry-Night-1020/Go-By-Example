package waitgroups

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d Starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func PrintValue() {
	var wg sync.WaitGroup

	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("All work done")
}
