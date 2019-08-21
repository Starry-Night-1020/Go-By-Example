package main

import (
	"fmt"
	"os"
)

func main() {
	argWithPath := os.Args
	argWithoutPath := os.Args[1:]

	params := os.Args[3]
	fmt.Println(argWithPath)
	fmt.Println(argWithoutPath)
	fmt.Println(params)
}
