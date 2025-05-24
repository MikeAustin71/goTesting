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

	newSlice = getRange1To3Cap10(mSlice1)

	printSlice("range1To3Cap10", newSlice)

	newSlice2 := getRange1To3Cap20(&mSlice1)

	printSlice("range1To3Cap10", newSlice2)

	modifySlice(mSlice1, "Mike", 2)

	printSlice("Modify 2 To Mike newSlice", newSlice2)

	printSlice("After Modify 2 To Mike Original Slice", mSlice1)

	mSlice2 := createBaseSlice()

	printSlice("Original Basic Slice #2", mSlice2)

	newSlice3 := modifySliceDontChangeOriginal(mSlice2, "Mike", 2)

	printSlice("Modified Slice #3", newSlice3)

	printSlice("Original Basic Slice #2", mSlice2)

}

func printSlice(title string, xSlice []string) {

	fmt.Printf("\n\n%v\n"+
		"XSlice = \n%v\n"+
		"Length of XSlice = %v\n"+
		"Capacity of XSlice = %v\n",
		title,
		xSlice, len(xSlice), cap(xSlice))

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

func getRange1To3Cap10(xSlice []string) []string {

	newSlice1 := xSlice[1:3:10]

	return newSlice1
}

func getRange1To3Cap20(xSlice *[]string) []string {

	ySlice := *xSlice
	newSlice1 := ySlice[1:3:cap(ySlice)]

	return newSlice1
}

func modifySlice(xSlice []string, name string, targetElement int) []string {

	xSlice[targetElement] = name

	return xSlice
}

func modifySliceDontChangeOriginal(xSlice []string, name string, targetElement int) []string {

	newTarget := targetElement + 1

	lenOfSlice := len(xSlice)

	var newXSlice []string

	if lenOfSlice <= newTarget {
		newXSlice = append(newXSlice, xSlice[:]...)
		newXSlice = append(newXSlice, name)
	} else {
		newXSlice = append(newXSlice, xSlice[:targetElement]...)
		newXSlice = append(newXSlice, name)
		newXSlice = append(newXSlice, xSlice[targetElement:]...)

	}
	return newXSlice
}
