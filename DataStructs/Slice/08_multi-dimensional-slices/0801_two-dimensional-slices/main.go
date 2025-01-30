package main

import (
	"fmt"
)

func main() {

	d := make([][]string, 0)
	e := []string{"h1", "h2", "h3"}

	d = append(d, e)

	f := []string{"x1", "x2", "x3", "x4"}

	d = append(d, f)

	fmt.Println("d[0]", d[0])
	fmt.Println("d[1]", d[1])

	fmt.Println("Printing d[0]")
	for i := 0; i < 3; i++ {
		fmt.Printf("d[0][%v]= %v \n", i, d[0][i])
	}
	fmt.Println()
	fmt.Println("Printing d[1]")
	for i := 0; i < 4; i++ {
		fmt.Printf("d[1][%v]= %v \n", i, d[1][i])
	}

}

/*
$ go run main.go
d[0] [h1 h2 h3]
d[1] [x1 x2 x3 x4]
Printing d[0]
d[0][0]= h1
d[0][1]= h2
d[0][2]= h3

Printing d[1]
d[1][0]= x1
d[1][1]= x2
d[1][2]= x3
d[1][3]= x4

*/
