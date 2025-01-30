package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference Time
	// 	Mon Jan 2 15:04:05 -0700 MST 2006

	baseTime, _ := time.Parse("Jan 2 15:04:05 -0700 2006",
		"Aug 17 22:53:05 -0500 2016")
	fmt.Println("Base Time:", baseTime)
	// Base Time: 2016-08-17 22:53:05 -0500 CDT

	mins := 10
	later := baseTime.Add(time.Duration(mins) * time.Minute)
	fmt.Println("Base Time + 10-minutes:", later)
	// Base Time + 10-minutes: 2016-08-17 23:03:05 -0500 CDT

}

/*	Output
	$ go run main.go
	Base Time: 2016-08-17 22:53:05 -0500 CDT
	Base Time + 10-minutes: 2016-08-17 23:03:05 -0500 CDT
*/
