package main

import (
	"../common"
	"fmt"
)

func main() {

	result := common.NumTree{}.Add(2, 2)

	fmt.Println("Result 2 + 2 = ", result)

	nTree := common.NumTree{}

	nTree.AddToThis(result)
	nTree.AddToThis(result)
	nTree.AddToThis(result)

	fmt.Println("nTree Num = ", nTree.Num)

	finalResult := nTree.Add(9, 9)

	fmt.Println("finalResult = ", finalResult)

}
