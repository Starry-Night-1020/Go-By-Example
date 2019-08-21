package arrays

import (
	"fmt"
)

func PrintValue() {
	var a [5]int
	fmt.Println(a)

	var b [5]int
	b[4] = 5
	fmt.Println(b)

	var c [5]int
	for i := range c {
		fmt.Println(i)
	}

	d := [5]int{1, 2, 3, 4, 5}
	fmt.Println(d)

	var e [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(e[i][j])
		}
		fmt.Println()
	}
}
