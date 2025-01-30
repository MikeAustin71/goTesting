package main

import (
	"fmt"
	"time"
)

func main() {
	const shortForm = "2006-01-02"
	loc, _ := time.LoadLocation("America/Chicago")

	t, _ := time.ParseInLocation(shortForm,
		"2015-01-19", loc)

	fmt.Println("Time Result: ", t)

}

/* Output
	$ go run main.go
	Time Result:  2015-01-19 00:00:00 -0600 CST

 */