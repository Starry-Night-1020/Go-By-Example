package writing_files

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
}

func PrintValue() {
	d1 := []byte("hello\ngo\n")

	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)
	defer f.Close()

	d2 := []byte("a,b,c,d,e\n")
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	f.Sync()

	w := bufio.NewWriter(f)

	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)
	w.Flush()
}
