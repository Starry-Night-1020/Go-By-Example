package map_example

import "fmt"

func PrintValue() {
	m := make(map[string]int)
	m["k1"] = 14
	m["k2"] = 13
	fmt.Println(m)

	n := m["k1"]
	fmt.Println(n)

	fmt.Println(len(m))

	delete(m, "k2")
	fmt.Println(m)

	_, ptr := m["k2"]
	fmt.Println(ptr)

	fmt.Println(m)
}
