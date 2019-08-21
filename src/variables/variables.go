package variables

import (
	"fmt"
)

func PrintValue() {
	var a, b int = 1, 2
	fmt.Println(a, b)

	var c string = "initial"
	fmt.Println(c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "string"
	fmt.Println(f)
}
