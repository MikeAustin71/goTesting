package main

import (
	"fmt"
	"time"
)

// Taken from https://bl.ocks.org/joyrexus/a56717634a672dcdfd48

func main() {
	t0 := time.Now()
	fmt.Println("The time is now ...")
	printTime(t0)

	// To convert an integer number of units to a Duration, just multiply.
	// see `http://golang.org/pkg/time/#Duration`.
	mins := 10
	days := 2

	t1 := t0.Add(time.Duration(mins) * time.Minute)
	fmt.Printf("\n... and in %v minutes the time will be ...\n", mins)
	printTime(t1)

	fmt.Printf("\n... and %v days after that ...\n", days)
	t2 := t1.Add(time.Duration(days) * time.Hour * 24)
	printTime(t2)

	// Use `AddDate(y, m, d)` to add the given number of years, months, days.
	years := 2
	months := 2

	fmt.Printf(
		"\n... and %v years, %v months, and %v days after that ...\n",
		years, months, days,
	)
	t3 := t2.AddDate(years, months, days)
	printTime(t3)

}

func printTime(t time.Time) {
	// see http://golang.org/pkg/time/#pkg-constants
	fmt.Println(t.Format(time.RFC822))

	// see http://golang.org/pkg/time/#Time.Format
	const layout = "Jan 2, 2006 at 3:04pm (MST)"
	fmt.Println(t.Format(layout))
}

/* Output - Note: Output will vary.
	$ go run main.go
	The time is now ...
	17 Aug 16 15:43 CDT
	Aug 17, 2016 at 3:43pm (CDT)

	... and in 10 minutes the time will be ...
	17 Aug 16 15:53 CDT
	Aug 17, 2016 at 3:53pm (CDT)

	... and 2 days after that ...
	19 Aug 16 15:53 CDT
	Aug 19, 2016 at 3:53pm (CDT)

	... and 2 years, 2 months, and 2 days after that ...
	21 Oct 18 15:53 CDT
	Oct 21, 2018 at 3:53pm (CDT)

*/
