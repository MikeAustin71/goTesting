package main

import "fmt"

func main() {

	a := false
	b := false

	if a || b {
		fmt.Println("This will not run")
	} else {
		fmt.Println("This else condition will run!")
	}
}
