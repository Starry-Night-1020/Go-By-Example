package parsing3

import (
	"fmt"
	"net/url"
)

func PrintValue() {
	p := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v&isHost=true"

	u, error := url.Parse(s)
	if error != nil {
		panic(error)
	}
	p(u.Scheme)
	p(u.Hostname())
	p(u.Port())
	p(u.Path)

	p(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	for k, v := range m {
		fmt.Printf("%s %s\n", k, v[0])
	}
}
