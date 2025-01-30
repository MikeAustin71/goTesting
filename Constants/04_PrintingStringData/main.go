package main

import "fmt"

type Day int

func (day Day) String() string {
	return days[day]
}

const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday

)
// [...] is to tell the Go intrepreter/compiler to figure out the array size
var days = [...]string {"Monday","Tuesday","Wednesday","Thursday","Friday",}


func main() {
	fmt.Println("Monday String Value:", Monday)
	fmt.Println("Tuesday String Value:", Tuesday)
	fmt.Println("Wednesday String Value:", Wednesday)
}

