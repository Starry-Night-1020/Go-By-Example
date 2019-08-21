package multi_return_values

import "fmt"

func multi_return_values() (int, int) {
	return 3, 7
}

func PrintValue() {
	a, b := multi_return_values()
	fmt.Println(a)
	fmt.Println(b)
}
