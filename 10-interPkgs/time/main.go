package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now() //.Format("2006/01/02")
	p(now)

	then := time.Date(2019, 11, 17, 20, 34, 58, 0, time.UTC)
	p(then)

	then = then.Add(time.Hour * 24)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	p(then.Weekday())

	p("Then es antes que Now:", then.Before(now))
	p("Then es despues que Now:", then.After(now))
	p("Then es igual que Now:", then.Equal(now))

	//Restar fechas
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
}
