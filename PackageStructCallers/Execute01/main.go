package main

import (
	"fmt"

	"github.com/mikeaustin71/PackageStruct02/pkgStr/strStuff"
)

func main() {

	myStrStuffer := new(strStuff.StrStuff)

	err := myStrStuffer.SetTheString("Testing Hello World")

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}

	err = myStrStuffer.PrintTheString()

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}

	err = myStrStuffer.MyEmployee.SetEmployee("Mike", 21)

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}

	err = myStrStuffer.MyEmployee.PrintEmployeeInfo()

	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		return
	}

	return
}
