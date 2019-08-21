package methods

import "fmt"

type rect struct {
	width  int
	height int
}

func cal(rec rect) int {
	return rec.width * rec.height
}

func PrintValue() {
	var rec rect = rect{width: 10, height: 5}
	fmt.Println(cal(rec))
}
