package main

import "fmt"

func main() {

	var slice []rune

	if slice == nil {
		fmt.Println("'slice' is nil!")
	}

	lenOfSlice := len(slice)

	fmt.Printf("Length of nil slice= '%v'\n",
		lenOfSlice)

}
