package main

import "fmt"

type Day int

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday

)

func main() {
	fmt.Println("Monday Value", Monday)
	fmt.Println("Tuesday Value", Tuesday)
	fmt.Println("Wednesday Value", Wednesday)
}
