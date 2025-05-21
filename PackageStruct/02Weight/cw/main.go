package main

import (
	"fmt"
	"golangmikesamples/PackageStruct/02Weight/weightconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cw: %v\n", err)
			os.Exit(1)
		}
		p := weightconv.Pound(t)
		k := weightconv.Kilogram(t)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.PToK(p), k, weightconv.KToP(k))
	}
}
