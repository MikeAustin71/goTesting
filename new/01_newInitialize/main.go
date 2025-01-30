package main

import (
	"fmt"
)

type RockWald struct {
	Lat float64
	Lon float64
}

func main() {
	x := new(RockWald)
	y := &RockWald{}

	fmt.Println("*x==*y", *x == *y)
}
