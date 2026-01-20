// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mikeaustin71/PackageStruct/01Temp/tempconv"
)

func main() {
	fmt.Printf("Starting cf\\main.go ...\n")

	// Select the first argument
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, err2 := fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			if err2 != nil {
				fmt.Printf("Error #1: %v\n"+
					"Error #2: %v", err.Error(), err2.Error())
			}
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

//!-
