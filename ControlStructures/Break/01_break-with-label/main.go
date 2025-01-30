package main

import (
	"fmt"
)

func main()  {
	var jCnt = 0
	var iCnt = 0
	var cycleCnt = 0

REDO:

	for i := 0; i < 20; i++ {

		for j := 0; j < 2; j++ {

			jCnt = j

			if i == 3 {
				break REDO
			}

			fmt.Println("j=", j, "i =", i)
		} // End of j loop

		iCnt = i
		cycleCnt++
	} // End of i loop

	fmt.Println("Value of cycleCnt =", cycleCnt)
	fmt.Println("Value of iCnt     =", iCnt)
	fmt.Println("Value of jCnt     =", jCnt)
}


