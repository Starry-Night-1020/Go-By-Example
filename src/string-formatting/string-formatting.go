package formatting

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func PrintValue() {
	p := point{1, 2}
	fmt.Printf("%v\n", p)

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
