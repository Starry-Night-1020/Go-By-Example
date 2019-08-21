package pointers

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeropointer(iptr *int) {
	*iptr = 0
}

func PrintValue() {
	i := 1
	zeroval(i)
	fmt.Println(i)

	i = 2
	zeropointer(&i)
	fmt.Println(i)

	fmt.Println(&i)
}
