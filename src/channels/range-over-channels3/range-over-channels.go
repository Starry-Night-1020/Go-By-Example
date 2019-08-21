package channels3

import "fmt"

func PrintValue() {
	queue := make(chan string, 2)

	queue <- "one"
	queue <- "two"
	close(queue)
	for i := range queue {
		fmt.Println("i", i)
	}
}
