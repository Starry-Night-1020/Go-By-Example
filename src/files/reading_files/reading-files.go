package reading_files

import (
	"fmt"
	"io/ioutil"
)

func check(err error) {
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
}

func PrintValue() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Println(string(dat))
}
