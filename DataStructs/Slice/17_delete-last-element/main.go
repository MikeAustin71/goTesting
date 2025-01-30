package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("     Initial Value")
	fmt.Println("---------------------------")
	fmt.Printf("  Length: %v\n",
		len(s))
	fmt.Printf("Capacity: %v\n",
		cap(s))
	fmt.Println(" Values: ")
	for i, value := range s {
		fmt.Println("index=", i, " value=", value)
	}

	fmt.Println()

	s = s[:len(s)-1]

	fmt.Println("Value After Delete Last Element")
	fmt.Println("-------------------------------")
	fmt.Printf("  Length: %v\n",
		len(s))
	fmt.Printf("Capacity: %v\n",
		cap(s))
	fmt.Println(" Values: ")
	for i, value := range s {
		fmt.Println("index=", i, " value=", value)
	}

}
