package constant

import (
	"fmt"
	"math"
)

const n = 1000000

func PrintValue() {
	fmt.Println(n)

	const a = 8000000000
	fmt.Println(3e20 / a)

	fmt.Println(math.Sin(n))
}
