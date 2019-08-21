package functions5

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func PrintValue() {
	p("Contains", s.Contains("test", "t"))
	p("Count", s.Count("test", "t"))
	p("Prefix", s.HasPrefix("test", "te"))
	p("Suffix", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "e"))
	p("Join", s.Join([]string{"a", "b", "c"}, "-"))
	p("Replace", s.Replace("foo", "o", "i", -1))
	p("Replace", s.Replace("foo", "o", "i", 1))
	p("Split", s.Split("a-b-c-d-e", "-"))
	p("toLower", s.ToLower("ABC"))
	p("toUpper", s.ToUpper("abc"))
	p("len", len("test"))
	p("Char", "hello"[1])
}
