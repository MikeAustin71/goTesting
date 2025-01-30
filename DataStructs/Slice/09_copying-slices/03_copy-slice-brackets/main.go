package main

import (
	"fmt"
)


// copy from to index
func main() {
	// last '\' is zero based idx 12
	pathFile := "D:\\dirA\\dirB\\fileName.txt"
	lengthPathFile := len(pathFile)
	// dot (.) is zero based idx 21
	nameStartIdx := 13
	nameEndIdx := 21

	fileName := pathFile[nameStartIdx:nameEndIdx]

	fmt.Println("    pathFile: ", pathFile)
	fmt.Println("Length of pathFile: ", lengthPathFile)
	fmt.Println("nameStartIdx: ", nameStartIdx)
	fmt.Println("  nameEndIdx: ", nameEndIdx)
	fmt.Println("    fileName: ", fileName)

}
