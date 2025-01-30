package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	f64ZeroTest01()
}

func f64ZeroTest01() {

	funcName := "f64ZeroTest01"

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\nFunction: %v\n",
		funcName)

	fmt.Printf(breakStr + "\n\n")

	var float32Num float32

	float32Num = -50.125

	var float64Num float64

	float64Num = float64(float32Num)

	numberStr := strconv.FormatFloat(
		float64Num, 'f', -1, 64)

	fmt.Printf("Original Num Str: %v\n",
		numberStr)

	numberRunes := []rune(numberStr)

	lenNumRunes := len(numberRunes)

	var intRunes, fracRunes []rune

	foundRadixPoint := false

	foundMinusSign := false

	for i := 0; i < lenNumRunes; i++ {

		if numberRunes[i] == '-' {

			foundMinusSign = true

			continue
		}

		if numberRunes[i] == '.' {

			foundRadixPoint = true

			continue
		}

		if numberRunes[i] >= '0' &&
			numberRunes[i] <= '9' {

			if !foundRadixPoint {

				intRunes = append(
					intRunes, numberRunes[i])
			} else {

				fracRunes = append(
					fracRunes, numberRunes[i])
			}
		}
	}

	fmt.Printf("Integer Value: %v\n",
		string(intRunes))

	fmt.Printf("Fractional Value: %v\n",
		string(fracRunes))

	var numSign string

	if foundMinusSign {
		numSign = "Negative"
	} else {
		numSign = "Positive"
	}

	fmt.Printf("Number Sign is: %v\n",
		numSign)

	fmt.Printf("\n\n" + breakStr + "\n Successful Completion!\n" +
		"Function: " +
		funcName + "\n\n")

}
