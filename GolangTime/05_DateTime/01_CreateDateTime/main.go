/*
Create Time, time.Date. To create a Time from ints (like the year, month and day) we call the time.Date func. The time.Date func returns a new Time instance.
Tip:
Don't pass nil as the Location (the last argument). This will cause a panic.

This Golang program uses time.Date to create Time.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a Date in 2006.
	t := time.Date(2006, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println(t)

	// Create a Date in 2015
	t2 := time.Date(2015, 4, 12, 7, 7, 0, 0, time.UTC)
	fmt.Println(t2)

}

/*	Output
	$ go run main.go
	2006-01-01 12:00:00 +0000 UTC
	2015-04-12 07:07:00 +0000 UTC
*/
