package branch

import (
	"fmt"
)

func PrintValue() {
	if 7%2 == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
