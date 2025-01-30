package main

import "fmt"

func main() {

	var a = [7][2]int{}

	for i:=0; i < 7; i++ {
		a[i][0] = i
		a[i][1] = 2
	}

	for k:=0; k < len(a); k++ {
		fmt.Println("k= ", a[k][0], "  2nd-Dimension= ", a[k][1])
	}

}
