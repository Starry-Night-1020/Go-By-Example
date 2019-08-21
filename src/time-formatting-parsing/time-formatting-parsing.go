package parsing

import (
	"fmt"
	"time"
)

func PrintValue() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))
	p(t.Format(time.RFC1123))
}
