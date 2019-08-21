package time2

import (
	"fmt"
	"time"
)

func PrintValue() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(now.Year())
	p(now.Location())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(then.Add(diff))
	p(then.Add(-diff))
}
