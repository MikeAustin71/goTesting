package main

import "fmt"

const (
	Add = iota + 1
	Subtract
	Multiply
)

// [...] is to tell the Go intrepreter/compiler to figure out the array size
var mathOpStr = [...]string{"", "Add", "Subtract", "Multiply"}

func main() {
	fmt.Println("String Value: '", mathOpStr[Add], "' Integer Value: ", Add)
	fmt.Println("String Value: '", mathOpStr[Subtract], "' Integer Value: ", Subtract)
	fmt.Println("String Value: '", mathOpStr[Multiply], "' Integer Value: ", Multiply)
}
