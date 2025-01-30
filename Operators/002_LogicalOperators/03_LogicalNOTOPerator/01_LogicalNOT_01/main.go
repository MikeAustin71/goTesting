package main

import "fmt"

func main() {

	a := 3

	if !(a==3) {
		fmt.Println("This did not run")
	}

	if !(a==2){
		fmt.Println("This ran")
	}

}
