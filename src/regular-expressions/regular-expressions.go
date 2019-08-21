package expressions

import (
	"bytes"
	"fmt"
	"regexp"
)

func PrintValue() {
	match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p[a-z]+ch")

	fmt.Println(r.MatchString("peach"))
	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach"))
	fmt.Println(r.FindAllStringIndex("peach punch", 2))

	fmt.Println(r.FindAllStringSubmatch("peach punch pinch apabcdch", -1))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a punch", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
