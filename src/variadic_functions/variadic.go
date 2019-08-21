package variadic_functions

import "fmt"

func sum(nums ...int) int {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func PrintValue() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
}
