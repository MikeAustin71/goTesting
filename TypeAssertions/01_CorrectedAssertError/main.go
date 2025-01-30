package main

import (
	"io"
	"os"
	"bytes"
	"fmt"
)

// This example demonstrates a correction for the
// failed example in directory 01_AssertError above.

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	fmt.Printf("f is of type %T \n",f)

	c,ok := w.(*bytes.Buffer)

	if !ok {
		fmt.Println("c is NOT of Type *bytes.Buffer - Terminating Execution!")
		return
	}

	fmt.Printf("c is of type %T \n", c)

}

/*	Output
	$ go run main.go
	f is of type *os.File
	c is NOT of Type *bytes.Buffer - Terminating Execution!
*/
