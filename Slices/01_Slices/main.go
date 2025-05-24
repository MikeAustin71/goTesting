package main

import "fmt"

func main() {

	mSlice1 := createBaseSlice()

	printSlice(mSlice1)

	changeASliceElement(mSlice1)

	printSlice(mSlice1)

}

func printSlice(xSlice []string) {

	fmt.Printf("\n\n%v\n"+
		"XSlice = \n%v\n"+
		"Length of XSlice = %v\n\n",
		"Printing Slice xSlice",
		xSlice, len(xSlice))

	return
}

func createBaseSlice() []string {

	mSlice1 := []string{"Sam", "Adrian", "James", "Roy", "Adam"}

	mSlice1 = append(mSlice1, "Mike")

	return mSlice1
}

func changeASliceElement(xSlice []string) []string {

	lenSlice := len(xSlice)

	lastElement := lenSlice - 1

	xSlice[lastElement] = "Miguel"

	return xSlice
}
