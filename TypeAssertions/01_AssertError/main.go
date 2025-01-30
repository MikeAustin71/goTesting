package main

import (
	"io"
	"os"
	"bytes"
	"fmt"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)      // success: f == os.Stdout
	c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	fmt.Printf("f is of type %T \n",f)
	fmt.Printf("c is of type %T \n", c)
}
