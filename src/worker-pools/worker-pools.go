package pools

import (
	"fmt"
	"time"
)

func worker(id int, jobs chan int, result chan int) {
	for j := range jobs {
		fmt.Println("worker", id, "process job", j)
		time.Sleep(time.Second)
		result <- 2 * j
		fmt.Println("job finished", j)
	}
}

func PrintValue() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, result)
	}

	for i := 1; i <= 9; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= 9; i++ {
		<-result
	}
	fmt.Println("all jobs finished")
}
