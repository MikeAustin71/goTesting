package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference Time
	// 	Mon Jan 2 15:04:05 -0700 MST 2006
	loc, _ := time.LoadLocation("Europe/London")
	const longForm = "Jan 2, 2006 15:04:00"
	baseTime, _ := time.ParseInLocation(longForm, "Aug 24, 2016 19:00:00", loc)
	fmt.Println("Webinar starts", baseTime)
	here, _ := time.LoadLocation("America/Chicago")
	tHere := baseTime.In(here)
	fmt.Println("The Webinar starts here:", tHere)
}

/*	Output
	$ go run main.go
	Webinar starts 2016-08-24 19:00:00 +0100 BST
	The Webinar starts here: 2016-08-24 13:00:00 -0500 CDT
*/
