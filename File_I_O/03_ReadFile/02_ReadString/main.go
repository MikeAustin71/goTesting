package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	This example reads a text file line by line using
	bufio Readstring. Note that the '/n' new line character
	is included in the returned string.
*/

func checkError(e error, msg string) {

	if e != nil {
		fmt.Println(msg)
		panic(e)
	}

}

func main() {

	f, err := os.Open("./TestText.txt")

	checkError(err, "File failed to open!")

	defer f.Close()

	reader := bufio.NewReader(f)

	for {

		text, err := reader.ReadString('\n')

		if err != nil && err == io.EOF {
			break
		}

		checkError(err, "Read String Error")

		fmt.Printf("%s", text)

	}
}

/* 	Output
	$ go run main.go
	Line # 1 Some Text
	Line # 2 Some Text
	Line # 3 Some Text
	Line # 4 Some Text

 */
