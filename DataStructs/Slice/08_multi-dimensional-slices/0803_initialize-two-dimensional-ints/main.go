package main

import (
	"fmt"
)

func main()  {
	levels :=	10
	places := 5
	d:= make([][]int,levels)


	for i:=0; i < levels ; i++ {
		d[i] = make([] int, places)
	}


	for j:=0; j < levels ; j++ {
		for k:=0; k < places; k++ {
			d[j][k] = k
		}
	}

	for l:=0; l < levels; l++ {
		for m:=0; m < places; m++ {
			fmt.Printf("Level= %v  Place= %v Value= %v ",l, m, d[l][m])
		}
		fmt.Println()
	}


}

