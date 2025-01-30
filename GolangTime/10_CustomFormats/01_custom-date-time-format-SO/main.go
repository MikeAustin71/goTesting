package main

import (
	"fmt"
	"time"
)

/* Taken from Stack Overflow
http://stackoverflow.com/questions/14106541/go-parsing-date-time-strings-which-are-not-standard-formats

*/

func main() {
	test, err := time.Parse("01/02/2006", "10/15/1983")
	if err != nil {
		panic(err)
	}

	fmt.Println(test)
}

/* Output
	$ go run main.go
	1983-10-15 00:00:00 +0000 UTC

*/
