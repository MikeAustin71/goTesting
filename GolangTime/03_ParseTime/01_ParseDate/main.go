package main

import (
	"fmt"
	"time"
)

func main() {
	// This is the value we are trying to parse.
	value := "January 28, 2015"

	// The form must be January 2,2006.
	// Note: The extended reference date-time is:
	// 	Mon Jan 2 15:04:05 -0700 MST 2006
	form := "January 2, 2006"

	// Parse the string according to the form.
	t, _ := time.Parse(form, value)
	fmt.Println(t)
}

/*	Output
	$ go run main.go
	2015-01-28 00:00:00 +0000 UTC
*/
