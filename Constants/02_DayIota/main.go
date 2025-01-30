package main

import "fmt"
const (
	Monday = iota
	Tuesday
	Wednesday
)

func main() {
	fmt.Println("Monday Value", Monday)
	fmt.Println("Tuesday Value", Tuesday)
	fmt.Println("Wednesday Value", Wednesday)
}
