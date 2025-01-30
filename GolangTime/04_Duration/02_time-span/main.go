package main

import (
	"fmt"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func main() {
	start, _ := time.Parse(time.RFC822, "16 Jan 15 10:00 UTC")
	end, _ := time.Parse(time.RFC822, "16 Jan 16 10:00 UTC")

	in, _ := time.Parse(time.RFC822, "16 Jan 15 20:00 UTC")
	out, _ := time.Parse(time.RFC822, "16 Jan 17 10:00 UTC")

	if inTimeSpan(start, end, in) {
		fmt.Println(in, "is between", start, "and", end, ".")
	}

	if !inTimeSpan(start, end, out) {
		fmt.Println(out, "is not between", start, "and", end, ".")
	}
}

/*	Output
	$ go run main.go
	2015-01-16 20:00:00 +0000 UTC is between 2015-01-16 10:00:00 +0000 UTC and 2016-01-16 10:00:00 +0000 UTC .
	2017-01-16 10:00:00 +0000 UTC is not between 2015-01-16 10:00:00 +0000 UTC and 2016-01-16 10:00:00 +0000 UTC .
 */
