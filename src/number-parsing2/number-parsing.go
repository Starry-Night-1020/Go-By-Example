package parsing2

import (
	"fmt"
	"strconv"
)

func PrintValue() {
	p := fmt.Println
	f, _ := strconv.ParseFloat("1.2345", 64)
	p(f)

	g, _ := strconv.ParseInt("123555", 0, 64)
	p(g)

	k, _ := strconv.Atoi("123321")
	p(k)

	h, err := strconv.Atoi("wat")
	if err != nil {
		p(h)
		p(err)
	}
	p(h)
}
