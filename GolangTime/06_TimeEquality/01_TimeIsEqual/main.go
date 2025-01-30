package main

/*
Equal. With the Equal() func we see if two times indicate the same instant. Equal() ignores Locations. With the "==" operator, the Locations must also be equal.
Note:
Equal() is a more conceptual time comparison, which accounts for Locations. The "==" operator just compares two Time instances.

 */
import (
	"fmt"
	"time"
)

func main() {
	// These times are created with equal values.
	t1 := time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("t1=", t1)

	t2 := time.Date(2015, 1, 1, 12, 0, 0, 0, time.UTC)
	fmt.Println("t2=", t2)

	// Compare times.
	if t1 == t2 {			// Yields true
		fmt.Println("Times t1 and t2 are equal")
	}
	// Equal compares date times but ignores locations.
	if t1.Equal(t2) {		// Yields true
		fmt.Println("true - t1 == t2")
	}

	central, _ := time.LoadLocation("America/Chicago")

	t3:= t2.In(central)

	fmt.Println("t3 = ",t3)

	if t1!=t3 {	// t!=t3 yields true
		fmt.Println("t1!=t3 is true")
		fmt.Println("t1 is NOT equal to t3")
	}

	if t1.Equal(t3) { // Yields true
		fmt.Println("Equal() ignores location and thinks")
		fmt.Println("t1 is equal to t3. t1.Equal(t3)=true")
	}

}

/*	Output
	$ go run main.go
	t1= 2015-01-01 12:00:00 +0000 UTC
	t2= 2015-01-01 12:00:00 +0000 UTC
	Times t1 and t2 are equal
	true - t1 == t2
	t3 =  2015-01-01 06:00:00 -0600 CST
	t1!=t3 is true
	t1 is NOT equal to t3
	Equal() ignores location and thinks
	t1 is equal to t3. t1.Equal(t3)=true
*/