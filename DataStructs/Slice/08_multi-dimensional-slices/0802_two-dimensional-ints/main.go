package main

import (
	"fmt"
)

func main()  {

	d:= make([][]int,0)
	e:= []int{1, 2, 3}

	d = append(d, e)

	f:= []int{4, 5, 6, 7}

	d = append(d,f)

	fmt.Println("d[0]", d[0])
	fmt.Println("d[1]", d[1])

	fmt.Println("Printing d[0]")
	for i:=0; i < 3; i++ {
		fmt.Printf("d[0][%v]= %v \n",i,d[0][i])
	}
	fmt.Println()
	fmt.Println("Printing d[1]")
	for i:=0; i < 4; i++ {
		fmt.Printf("d[1][%v]= %v \n",i,d[1][i])
	}

}

