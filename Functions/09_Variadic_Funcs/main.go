package main

import (
	"fmt"
	"strings"
)

func main() {
	T02Caller()
}

func T01Caller() {

	ePrefix := "T01Caller"
	title := "Test of Variadic Function"
	var total int
	var err error
	total, err = T01VariadicFunc(title, 1, 2, 3, 4, 5, 6, 8, 9, 10, 100)
	if err != nil {
		fmt.Printf("\n%v\n"+
			"Error returned from T01_VariadicFunc:\n"+
			"%v", ePrefix, err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"          SUCCESS!\n"+
		"Results of T01_VariadicFunc:\n"+
		"Total: %d\n\n", ePrefix, total)

}

func T02Caller() {

	ePrefix := "T01Caller"
	title := "Test of Variadic Function"
	var total int
	var err error
	total, err = T01VariadicFunc(title)
	if err != nil {
		fmt.Printf("\n%v\n"+
			"Error returned from T01_VariadicFunc:\n"+
			"%v", ePrefix, err.Error())

		return
	}

	fmt.Printf("\n%v\n"+
		"          SUCCESS!\n"+
		"Results of T01_VariadicFunc:\n"+
		"Total: %d\n\n", ePrefix, total)

}

func T01VariadicFunc(title string, numbers ...int) (total int, err error) {

	ePrefix := "T01_VariadicFunc"

	if len(title) < 1 {
		return 0, fmt.Errorf("\n%v\nInput Parameter Title is empty!\n\n", ePrefix)
	}

	titleLine := strings.Repeat("=", len(title)+7+4)

	fmt.Printf("\n\n%s\n", titleLine)
	fmt.Printf(" Title: %s\n", title)
	fmt.Printf("%s\n", titleLine)

	lenNums := len(numbers)

	fmt.Printf("\nLength of numbers is %d \n", lenNums)

	for _, number := range numbers {

		total += number

	}

	return total, err
}
