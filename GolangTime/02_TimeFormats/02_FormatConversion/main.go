/*
Format. With Format we convert a Date into a string. As with Parse, we use the special date Jan 2, 2006 to indicate the formatting.
Here:
We convert the result of baseTime to a simple date format. The Format method returns a string.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Reference Time
	// 	Mon Jan 2 15:04:05 -0700 MST 2006
	baseTime, _ := time.Parse("Jan 2 15:04:05 -0700 2006",
		"Aug 17 22:53:05 -0500 2016")

	// This date is used to indicate the layout.
	const layout1 = "Jan 2, 2006"
	res := baseTime.Format(layout1)
	fmt.Println("Format # 1", res)

	const layout2 = "January 2, 2006"
	res = baseTime.Format(layout2)
	fmt.Println("Format # 2", res)

	const layout3 = "Monday January 2, 2006"
	res = baseTime.Format(layout3)
	fmt.Println("Format # 3", res)

}

/*	Output
	$ go run main.go
	Format # 1 Aug 17, 2016
	Format # 2 August 17, 2016
	Format # 3 Wednesday August 17, 2016
*/
