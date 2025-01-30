package main

import (
	"fmt"
)

// Prints number in a 9-digit field with
// leading zeros

func main()  {
	num := 970

	s := fmt.Sprintf("Number: %09d", num)

	fmt.Println(" Original Number: ", num)
	fmt.Println("Formatted Output: ", s)
}
