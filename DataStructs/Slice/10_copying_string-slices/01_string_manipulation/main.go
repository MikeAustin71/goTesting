package main

import (
	"fmt"
	"strings"
)

// copy from to index
func main() {
	// last '\' is zero based idx 12


	pathFile := "D:\\xyz\\dirA\\dirB\\fileName.txt"

	srcBaseDir := "D:\\xyz"
	targetDir := "C:\\abc"

	idx := strings.Index(pathFile,srcBaseDir)

	if idx < 0 {
		fmt.Printf("Error: 'idx' is less than zero: idx='%v' ", idx)
		return
	}

	lenSrcBaseDir := len(srcBaseDir)

	newFile := targetDir + pathFile[lenSrcBaseDir:]


	fmt.Println("         pathFile: ", pathFile)
	fmt.Println("  Source Base Dir: ", srcBaseDir)
	fmt.Println("  Target Base Dir: ", targetDir)
	fmt.Println("    new File Name: ", newFile)

}

