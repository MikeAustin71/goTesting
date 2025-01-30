package main

import (
	"fmt"
	"time"
)

// Adapted from https://golang.org/pkg/time/#ParseInLocation

func main() {
	// Reference Time
	// 	Mon Jan 2 15:04:05 -0700 MST 2006

	loc, _ := time.LoadLocation("America/Chicago")

	const longForm = "Jan 2, 2006 15:04:00"
	baseTime, _ := time.ParseInLocation(longForm, "Aug 18, 2016 15:05:00", loc)
	fmt.Println(baseTime) // 2016-08-18 15:05:00 -0500 CDT

	// Pre-defined formats
	// https://golang.org/pkg/time/#pkg-constants
	fmt.Println(baseTime.Format(time.RFC822)) // 18 Aug 16 15:05 CDT

	fmt.Println(baseTime.Format(time.Kitchen)) // 3:05PM

}

/*	Output
	$ go run main.go
	2016-08-18 15:05:00 -0500 CDT
	18 Aug 16 15:05 CDT
	3:05PM
*/
