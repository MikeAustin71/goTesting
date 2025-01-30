package main

// This code was adapted from an example provided
// by Todd McLeod - Udemy Go Training
import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Todd", 44}
	fmt.Println(&c1.name) // 0x8201e4120

	changeMe(&c1)

	fmt.Println(c1)       // {Rocky 44}
	fmt.Println(&c1.name) // 0x8201e4120
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Todd 44}
	fmt.Println(&z.name) // 0x8201e4120
	z.name = "Rocky"
	fmt.Println(z)       // &{Rocky 44}
	fmt.Println(&z.name) // 0x8201e4120
	fmt.Println("Printing Name",z.name)  // Rocky
	fmt.Println("Printing Age", z.age)   // 44

}
