package main

import (
	"fmt"
	"time"
)

func main() {

	// get current timestamp
	// Local Time on this machine is
	// Central Daylight Time (CDT) UTC-5
	currenttime := time.Now().Local()

	printTime(currenttime)

}

func printTime(t time.Time) {
	// see http://golang.org/pkg/time/#pkg-constants
	fmt.Println(t.Format(time.RFC822))

	// see http://golang.org/pkg/time/#Time.Format
	const layout = "Jan 2, 2006 at 3:04pm (CDT)"
	fmt.Println(t.Format(layout))
}

/*	Output - Current Time will cause variations
	in Output.

	$ go run main.go
	15 Aug 16 16:40 CDT
	Aug 15, 2016 at 4:40pm (CDT)

*/
