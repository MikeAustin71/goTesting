package main

/*
	This example is designed to demonstrate Exported Identifiers.
	ReverseAString is a function declared in another Package: strUtilTest.
	ReverseAString is visible to the main function in the main package
	because it was exported from the package (strUtilTest) in which it
	was declared. In order to Export an identifier, the first letter
	of the Identifier's name MUST be capitalized.
*/

import (
	"fmt"
	"golangmikesamples/Identifiers/ExportedIdentifiers/strUtilTest"
)

func main() {
	v := "Hello"
	rv := strUtilTest.ReverseAString(v)
	fmt.Println("This is base string", v)
	fmt.Println("This is the reversed string", rv)
}

/*	Output
	$ go run main.go
	This is base string Hello
	This is the reversed string olleH
*/
