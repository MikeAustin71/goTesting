package main

/*
Sleep. This func receives a Duration. It then pauses execution on the current goroutine for that duration of time. Here we call sleep four times for 100 ms each time.
Duration:
We construct the Duration from nanoseconds. One million nanoseconds is one millisecond. We multiply that by 100 for 100 ms.

*/

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 4; i++ {
		// Duration receives nanoseconds.
		// ... 100 million nanoseconds is 100 milliseconds.
		d := time.Duration(1000 * 1000 * 100)
		// Sleep for 100 ms.
		time.Sleep(d)
		fmt.Println(time.Now())
	}
}

/*	Output - Note: Output will vary.

	$ go run main.go
	2016-08-16 13:45:00.9991186 -0500 CDT
	2016-08-16 13:45:01.0996029 -0500 CDT
	2016-08-16 13:45:01.2000558 -0500 CDT
	2016-08-16 13:45:01.3001332 -0500 CDT
*/
