package main

import "fmt"

/*
	This example demonstrates the use of multiple case
	evaluations in a switch statement.
*/

func main() {
	x := "Red Fish"

	switch x {
	case "Bear", "Raccoon", "Possum":
		fmt.Println("The animal is either a Bear, Raccoon or Possum")
	case "Puppy", "Wolf", "Coyote":
		fmt.Println("The animal is either a Puppy, Wolf or Coyote")
	case "Trout", "Red Fish", "Flounder":
		fmt.Println("The fish is either a Trout, Red Fish or Flounder")
	case "Eagle", "Buzzard", "Road Runner":
		fmt.Println("The bird is either an Eagle, Buzzard or Road Runner")
	default:
		fmt.Println("The animal is unknown")
	}

	fmt.Println("End of Execution - Program Terminated!")
}

/*	Output
	$ go run main.go
	The fish is either a Trout, Red Fish or Flounder
	End of Execution - Program Terminated!
*/
