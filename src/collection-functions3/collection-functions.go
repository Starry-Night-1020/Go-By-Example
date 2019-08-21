package functions3

import (
	"fmt"
	"strings"
)

func Index(s []string, t string) int {
	for i, str := range s {
		if str == t {
			return i
		}
	}
	return -1
}

func Include(s []string, t string) bool {
	return Index(s, t) >= 0
}

func Any(s []string, f func(string) bool) bool {
	for _, str := range s {
		if f(str) {
			return true
		}
	}
	return false
}

func All(s []string, f func(string) bool) bool {
	for _, str := range s {
		if !f(str) {
			return false
		}
	}
	return true
}

func Filter(s []string, f func(string) bool) []string {
	result := make([]string, 0)
	for _, str := range s {
		if f(str) {
			result = append(result, str)
		}
	}
	return result
}

func Map(s []string, f func(string) string) []string {
	vs := make([]string, len(s))
	for i, str := range s {
		vs[i] = f(str)
	}
	return vs
}

func PrintValue() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}
