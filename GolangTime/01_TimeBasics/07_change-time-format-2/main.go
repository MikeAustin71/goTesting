package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current time
	t := time.Now()
	fmt.Println("=============================")
	// see http://golang.org/pkg/time/#pkg-constants
	// RFC822      = "02 Jan 06 15:04 MST"
	fmt.Println(t.Format(time.RFC822))

	fmt.Println("=============================")
	// Now change the format
	// RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
	// RFC1123 with numeric zone
	fmt.Println(t.Format(time.RFC1123Z))

}
/*	Output
	$ go run main.go
	=============================
	15 Aug 16 20:38 CDT
	=============================
	Mon, 15 Aug 2016 20:38:02 -0500

*/