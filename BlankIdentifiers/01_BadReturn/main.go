package main

import (
	"fmt"
)

func returnTwoVars() (string, int) {
	return _, 6
}

func main() {
	a, b := returnTwoVars()
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}
