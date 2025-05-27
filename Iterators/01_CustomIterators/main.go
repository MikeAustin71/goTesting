package main

import "fmt"

// Courtesy of:
// [Flo Woelki](https://www.youtube.com/@FloWoelki)
// Go 1.23: Custom Iterators Explained - Best feature?!
// https://www.youtube.com/watch?v=iurUVx0Nquc
//
// # Custom Iterators are new to Golang 1.23
//
// "Use the 'range' keyword to iterate over collections
// that are not maps or slices."

func iter1(yield func(i int) bool) {

	for i := range 3 {
		if !yield(i) {
			return
		}
	}

}

func main() {

	// Test01()

	// Test02()

	// Test03()

	//Test04()

	Test05()

	Test06()
}

func Test06() {

	fmt.Printf("\n\n-----------------\n" +
		"    Test06\n\n")

	iter1(func(i int) bool {
		fmt.Println(i)
		return false
	})
}

func Test05() {

	fmt.Printf("\n\n-----------------\n" +
		"    Test05\n\n")

	iter1(func(i int) bool {
		fmt.Println(i)
		return true
	})
}

func Test04() {

	fmt.Printf("\n\n-----------------\n" +
		"    Test04\n\n")

	for i := range iter1 {
		//if i == 1 {
		//  return
		//}

		fmt.Println(i)
	}

}

func Test03() {
	fmt.Printf("\n\n-----------------\n" +
		"    Test03\n\n")

	s := []string{"Hello", "World"}

	for i, v := range s {
		// i=index v=value
		fmt.Println(i, v)
	}

}

func Test02() {
	fmt.Printf("\n\n-----------------\n" +
		"    Test02\n\n")

	s := []string{"Hello", "World"}

	for i := range s {

		fmt.Println(i) // Prints index 0 and 1

	}

}

func Test01() {

	fmt.Printf("\n\n-----------------\n" +
		"    Test01\n\n")

	for i := range 10 {
		fmt.Println(i)
	}
}
