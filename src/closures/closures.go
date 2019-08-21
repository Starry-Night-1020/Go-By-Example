package closures

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func PrintValue() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	nextInt = intSeq()
	fmt.Println(nextInt())
}
