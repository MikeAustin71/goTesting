package main

import "fmt"

func bunnyStatus(bunnySpeed int, coyoteSpeed int) string {
	if bunnySpeed <= coyoteSpeed {
		return "It's dinner time!"
	}

	return "Don't Worry - Be Happy!"
}


func main() {
	//bunnySpeed := 44
	//coyoteSpeed := 40
	bunnySpeed := 40
	coyoteSpeed := 44

	status := bunnyStatus(bunnySpeed, coyoteSpeed)
	fmt.Println("Bunny Status:", status)

}
