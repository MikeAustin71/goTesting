package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println("Before : ", now)

	// reduce the date format
	// remember NOT to use 2006-01-01 or 02-02 or same digit
	// for month and date. Will cause weird date result
	// fmt.Println(now.Format("2006-01-01")) <--- WRONG

	fmt.Println("After : ", now.Format("2006-01-02")) // <-- CORRECT
}

/*	Output - Note: Output will vary according to time.Now()
	$ go run main.go
	Before :  2016-08-18 21:34:10.2972664 -0500 CDT
	After :  2016-08-18
*/
