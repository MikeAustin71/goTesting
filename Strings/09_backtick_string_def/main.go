package main

import (
	"fmt"
	"regexp"
)

/*
Tebeka Shorts: Construct a Raw String
https://www.youtube.com/watch?v=cm2AEALBBo8&t=2s

Shows examples of string definitions using the
backtick symbol (`)
*/

func main() {

	s := `C:\to\be\or\not\to.bee`

	fmt.Println(s)

	// Used with regex configuration
	re := regexp.MustCompile(`\d+`)
	fmt.Println(re)

	poem := `
	The Road goes ever on and on,
	Down from the door where it began.
	`

	fmt.Println(poem)
}
