package main

import (
	"fmt"
	"time"
)

func main() {

	// Reference Time
	// 	Mon Jan 2 15:04:05 -0700 MST 2006

	t, _ := time.Parse("Jan 2 15:04:05 -0700 2006",
		"Aug 17 22:53:05 -0500 2016")

	fmt.Println(t)

	// Print year, month and day of the current time.
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
}

/*	Output:
	$ go run main.go
	2016-08-17 22:53:05 -0500 CDT
	2016
	August
	17
*/
