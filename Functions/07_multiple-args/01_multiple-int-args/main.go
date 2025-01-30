package main

import (
	"fmt"
)

func main()  {
	n1:= 5
	n2:= 6
	n3:= 7
	n4:= 8

	result := addIntArgs(n1 , n2, n3 ,n4)

	fmt.Println("Result= ", result)
	// expected Result= 26
}

func addIntArgs(num ... int) int {

	sum:= 0

	for _, num := range num {

		sum += num
	}

	return sum

}

