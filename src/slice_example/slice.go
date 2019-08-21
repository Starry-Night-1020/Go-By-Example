package slice_example

import "fmt"

func PrintValue() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println(len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[0:]
	fmt.Println(l)

	t := []string{"g", "f", "d", "s", "a"}
	fmt.Println(t)
	t = t[:2]

	n := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	n = n[:2]
	fmt.Println(n)

	twoD := make([][]int, 3)
	for i := range twoD {
		for j := range twoD {
			fmt.Println("i, j", i, j)
		}
	}
}
