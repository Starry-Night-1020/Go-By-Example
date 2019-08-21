package sorting

import (
	"fmt"
	"sort"
)

func PrintValue() {
	strs := []string{"c", "b", "a"}
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{3, 2, 1}
	fmt.Println(ints)
	sort.Ints(ints)
	fmt.Println(ints)
}
