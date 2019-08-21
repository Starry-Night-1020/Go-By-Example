package hashes

import (
	"crypto/sha1"
	"fmt"
)

func PrintValue() {
	s := "hahaha"
	sha1 := sha1.New()
	sha1.Write([]byte(s))
	sum := string(sha1.Sum(nil))
	fmt.Printf("this is a sha1 string: %x\n: ", sum)
}
