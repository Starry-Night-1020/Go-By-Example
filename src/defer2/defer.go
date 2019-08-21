package defer2

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("createFile")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	file.Close()
}

func PrintValue() {
	file := createFile("/tmp/file")
	defer closeFile(file)
	writeFile(file)
}
