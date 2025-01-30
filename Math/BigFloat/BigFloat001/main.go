package main

import (
	"fmt"
	"math/big"
	"strings"
)

func main() {

	ePrefix := "BigFloat001.main() "

	separator := strings.Repeat("-", 75)

	var floatNum, newFloat *big.Float

	var intVal *big.Int

	var accuracy big.Accuracy

	for i:=0; i < 5; i++ {

		switch i {
		case 0:
			floatNum = big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(12345.999999)

		case 1:
			floatNum =
				big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(12345.599)

		case 2:
			floatNum =
				big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(12345.499)


		case 3:
			floatNum =
				big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(12345.1003)

		case 4:
			floatNum =
				big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(12345.0003)

		default:
			fmt.Printf(ePrefix + "\n" +
				"Error: Iterator i INVALID!\n" +
				"i='%v'\n", i)
			return
		}

		intVal, accuracy = floatNum.Int(nil)

		newFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt(intVal)

		fmt.Println(ePrefix)
		fmt.Println(separator)
		fmt.Printf("        Float Value: %17.8f\n",
			floatNum)

		fmt.Printf("          Int Value: %s\n",
			intVal.Text(10))

		fmt.Printf("           accuracy: %v\n",
			accuracy)

		fmt.Printf("New Int Float Value: %17.8f\n",
			newFloat)
		fmt.Println(separator)

	}

	return
}
