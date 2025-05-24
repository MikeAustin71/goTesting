package main

import "fmt"

func main() {

	mSlice1 := createBaseSlice()

	printSlice("Basic Slice", mSlice1)

	mSlice1 = appendASliceElement(mSlice1)

	printSlice("Appended Slice", mSlice1)

	changeASliceElement(mSlice1)

	printSlice("Changed Slice", mSlice1)

	newSlice := getRangeOneToThree(mSlice1)

	printSlice("range1To3", newSlice)

	newSlice = getRangeTwoToAll(mSlice1)

	printSlice("range2ToAll", newSlice)

}

func printSlice(title string, xSlice []string) {

	fmt.Printf("\n\n%v\n"+
		"XSlice = \n%v\n"+
		"Length of XSlice = %v\n\n",
		title,
		xSlice, len(xSlice))

	return
}

func createBaseSlice() []string {

	mSlice1 := []string{"Sam", "Adrian", "James", "Roy", "Adam"}

	return mSlice1
}

func appendASliceElement(xSlice []string) []string {

	xSlice = append(xSlice, "Mike")

	return xSlice
}

func changeASliceElement(xSlice []string) []string {

	lenSlice := len(xSlice)

	lastElement := lenSlice - 1

	xSlice[lastElement] = "Miguel"

	return xSlice
}

func getRangeOneToThree(xSlice []string) []string {

	newSlice1 := xSlice[1:3]

	return newSlice1
}

func getRangeTwoToAll(xSlice []string) []string {

	newSlice1 := xSlice[2:]

	return newSlice1
}
