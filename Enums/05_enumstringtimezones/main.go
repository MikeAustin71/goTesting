package main

import (
	"./timezones"
	"fmt"
)

func main() {

	fmt.Println("Testing IANA Time Zones")

	str := timezones.IanaTz.America.Chicago()

	fmt.Println("US-Central: ", str )

	str = timezones.IanaTz.US.Mountain()

	fmt.Println("US-Mountain: ", str)

	fmt.Println("Cuba: ", timezones.IanaTz.Cuba())

}

