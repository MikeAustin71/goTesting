package main

import (
	"fmt"
	"time"
)

/*	Custom Date Format Example
	https://golang.org/pkg/time/#Time.Format
	Reference Time is for use in format
	Strings is:

	Mon Jan 2 15:04:05 -0700 MST 2006
*/

func main() {
	tDate := "10/15/1983"
	fmt.Println("Target Date is ", tDate)
	test, err := time.Parse("01/02/2006", tDate)
	if err != nil {
		panic(err)
	}

	fmt.Println("Unformatted Test Date is", test)

	outFormat := "Monday January 2, 2006"
	formattedDate := test.Format(outFormat)
	fmt.Println("Formatted Test Date is", formattedDate)

	outFormat = "2006-01-02 Monday"
	fmt.Println("Formatted Test Date #2 is", test.Format(outFormat))

	outFormat = "01/02/2006"
	fmt.Println("Formatted Test Date #3 is", test.Format(outFormat))

	pacific, _ := time.LoadLocation("America/Los_Angeles")
	pt := test.In(pacific)
	fmt.Println("Equivalent Pacific Time", pt)
	fmt.Println("Pacific Time Location", pt.Location())
}

/* Output
	$ go run main.go
	Target Date is  10/15/1983
	Unformatted Test Date is 1983-10-15 00:00:00 +0000 UTC
	Formatted Test Date is Saturday October 15, 1983
	Formatted Test Date #2 is 1983-10-15 Saturday
	Formatted Test Date #3 is 10/15/1983
	Equivalent Pacific Time 1983-10-14 17:00:00 -0700 PDT
	Pacific Time Location America/Los_Angeles
*/
