package range_example

import "fmt"

func PrintValue() {
	var nums = []int{2, 3, 4}

	for _, num := range nums {
		fmt.Println(num)
	}

	for i, num := range nums {
		fmt.Println("i, num:", i, num)
	}

	kvs := map[string]string{"apple": "apple", "go": "go"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	kvs2 := make(map[string]int)
	kvs2["k1"] = 1234
	kvs2["k2"] = 123
	fmt.Println(kvs2)
}
