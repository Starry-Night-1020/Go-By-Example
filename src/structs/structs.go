package structs

import "fmt"

type person struct {
	name string
	age  int
}

func PrintValue() {
	var one person = person{name: "one", age: 18}
	fmt.Println(one)

	var two person = person{name: "two", age: 16}
	fmt.Println(two)

	var iptr *person = &two
	iptr.name = "iptr"
	fmt.Println(two)
}
