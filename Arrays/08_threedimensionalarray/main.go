package main

import "fmt"

func main() {

		var a = [7][2][3]int{}

		for i:=0; i < 7; i++ {
			a[i][0][0] = i
			a[i][0][1] = 2
			a[i][0][2] = 3
			a[i][1][0] = 1
			a[i][1][1] = 2
			a[i][1][2] = 3
		}

		for k:=0; k < len(a); k++ {
			fmt.Println("k= ", a[k][0], "  3rd-Dimension= ", a[k][1][1])
		}

}
