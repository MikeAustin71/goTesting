package main

import (
	"bufio"
	"fmt"
	"os"
	)

/* 	This code reads a file line by line using 'scanner'
	Note that scanner will strip off the new line character
	when returning the string from scanner.Text()
*/

func main() {

	inFile, err := os.Open("./TestText.txt")

	defer inFile.Close()

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inFile)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		line := scanner.Text()

		fmt.Printf("%s\n", line)
		// Note the line received from scanner does
		// NOT contain a new line.
	}

}

/* 	Output
	$ go run main.go
	Line # 1 Some Text
	Line # 2 Some Text
	Line # 3 Some Text
	Line # 4 Some Text
*/
