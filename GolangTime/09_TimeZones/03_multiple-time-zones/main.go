package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a time
	// https://golang.org/pkg/time/#pkg-constants
	baseTime := "17 Aug 16 13:23 UTC"

	t, _ := time.Parse(
		time.RFC822,
		baseTime)
	fmt.Println("Base UTC Time", t)
	fmt.Println("Base Time Location", t.Location())

	pacific, _ := time.LoadLocation("America/Los_Angeles")
	pt := t.In(pacific)
	fmt.Println("Equivalent Pacific Time", pt)
	fmt.Println("Pacific Time Location", pt.Location())

	mountain, _ := time.LoadLocation("America/Denver")
	mt := t.In(mountain)
	fmt.Println("Equivalent Mountain Time", mt)
	fmt.Println("Mountain Time Location", mt.Location())

	central, _ := time.LoadLocation("America/Chicago")
	ct := t.In(central)
	fmt.Println("Equivalent Central Time", ct)
	fmt.Println("Central Time Location", ct.Location())

	eastern, _ := time.LoadLocation("America/New_York")
	et := t.In(eastern)
	fmt.Println("Equivalent Eastern Time", et)
	fmt.Println("Eastern Time Location", et.Location())
}

/*	Output
	$ go run main.go
	Base UTC Time 2016-08-17 13:23:00 +0000 UTC
	Base Time Location UTC
	Equivalent Pacific Time 2016-08-17 06:23:00 -0700 PDT
	Pacific Time Location America/Los_Angeles
	Equivalent Mountain Time 2016-08-17 07:23:00 -0600 MDT
	Mountain Time Location America/Denver
	Equivalent Central Time 2016-08-17 08:23:00 -0500 CDT
	Central Time Location America/Chicago
	Equivalent Eastern Time 2016-08-17 09:23:00 -0400 EDT
	Eastern Time Location America/New_York
*/
