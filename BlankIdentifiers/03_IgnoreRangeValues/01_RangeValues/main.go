package main

import "fmt"

func main() {
	// Create a string to string map.
	animals := map[string]string{}
	animals["cat"] = "Mittens"
	animals["dog"] = "Spot"

	// Loop over the map.
	for _, value := range animals {
		fmt.Println(value)
	}
}
