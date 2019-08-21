package functions

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}

func PrintValue() {
	strs := []string{"c", "bb", "aaa"}
	fmt.Println(strs)
	sort.Sort(ByLength(strs))
	fmt.Println(strs)
}
