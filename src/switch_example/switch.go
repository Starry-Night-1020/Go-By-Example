package switch_example

import (
	"fmt"
)

func PrintValue() {
	i := 2
	fmt.Print("2 write as ")
	switch i {
	case 1:
		fmt.Println("one")
		break
	case 2:
		fmt.Println("two")
		break
	case 3:
		fmt.Println("three")
		break
	default:
		break
	}
}
