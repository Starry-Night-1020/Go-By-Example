package for_example

import (
	"fmt"
)

func PrintValue() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	for j := 1; j <= 4; {
		fmt.Println(j)
		j += 1
	}
}
