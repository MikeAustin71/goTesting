package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Println("Test For *big.Int is Even!")


	bigNum := big.NewInt(24)
	bigOne := big.NewInt(1)

	fmt.Println("        Test Value: ", bigNum.Text(10))

	isEven := big.NewInt(0).And(bigNum, bigOne)

	isEvenResult := true

	if  isEven.Cmp(bigOne) == 0 {
		isEvenResult = false
	}

	fmt.Println("Test Value Is Even: ", isEvenResult)

}
