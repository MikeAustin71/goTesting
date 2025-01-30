package main

/*	Adapted from:
	https://bl.ocks.org/joyrexus/a56717634a672dcdfd48
*/

import (
	"fmt"
	"time"
)

func main() {
	// Get current local time:
	now := time.Now()
	fmt.Println("Local Time is", now)

	/* 	Taken from:
		https://bl.ocks.org/joyrexus/a56717634a672dcdfd48

	Construct a time with func Date:
	https://golang.org/pkg/time/#Date
	*/
	y := 2009
	mo := time.November
	d := 10
	h := 23
	mi := 0
	s := 0
	t := time.Date(y, mo, d, h, mi, s, 0, time.UTC)

	fmt.Println("Constructed Date Time is", t)
}

/*	Output - Note: Local Time will vary.
	$ go run main.go
	Local Time is 2016-08-17 21:18:03.6136834 -0500 CDT
	Constructed Date Time is 2009-11-10 23:00:00 +0000 UTC
*/
