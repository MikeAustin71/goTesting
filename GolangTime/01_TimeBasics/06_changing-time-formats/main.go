package main

import (
	"fmt"
	"time"
)
/*	Socket Loop Example
	https://www.socketloop.com/tutorials/golang-date-and-time-formatting
	It doesn't see to work.
 */
func main() {

	// get current timestamp
	currentTime := time.Now().Local()

	fmt.Println(currentTime)
	// don't like the standard format ?
	// change it
	newFormat := currentTime.Format("2006-01-02 15:00:00 +0800")
	fmt.Println(newFormat)

	fmt.Println("===============================")

	// Changing time layout(form)
	str := "Dec 29, 2014 at 7:54pm (SGT)"
	newLayout := "Jan 2, 2006 at 3:04pm (MST)"
	newTime, _ := time.Parse(newLayout, str)
	fmt.Println(newTime)
}
/*
	Output :

	2014-12-29 12:04:53.531591733 +0800 SGT
	2014-12-29 12:00:00 +0800
	===============================
	2014-12-29 19:54:00 +0800 SGT
*/