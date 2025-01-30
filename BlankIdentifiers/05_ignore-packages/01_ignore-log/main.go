package main

/*
	This example demonstrates the ability to
	ignore an imported package (log) using
	the blank identifier underscore character (_)
*/

import (
"fmt"
_ "log"
"time"
)

func main() {
	a:= time.Now()
	fmt.Println("The current time is:",a)
}

/*	Output
	$ go run main.go
	The current time is: 2016-08-08 20:08:47.8922278 -0500 CDT
	** Note Actual time will vary.
*/

