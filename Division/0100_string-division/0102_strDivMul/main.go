package main

import (
	"fmt"
	"math"
)


func main()  {
/*
	dividendAry := []int{4,2,5}
	divisorAry  := []int{0,2,5}
	zeroAry     := []int{0,0,0,0,0,0,0}
	tenAry      := []int{0,1,0}
	dividendAryFirstDigitIdx := findFirstSignificantDigitIdx(dividendAry)
	divisorAryFirstDigitIdx := findFirstSignificantDigitIdx(dividendAry)

	lenVariance := dividendAryFirstDigitIdx - divisorAryFirstDigitIdx
*/


	testRaiseToPower()

}

func testRaiseToPower(){
	tenAry      := []int{0,1,0}
	result := raiseToPower(tenAry, 1)

	fmt.Println("Result ", result)

}

func raiseToPower(n1Ary []int, power int) []int {

	lenN1Ary := len(n1Ary)

	n2Ary := make([]int, lenN1Ary)

	for i:= 0; i < lenN1Ary; i++{
		n2Ary[i] = n1Ary[i]
	}

	power--

	if power == 0 {
		return n1Ary
	}

	var result [] int
	for i:=0; i < power; i++ {

		result = multiply(n1Ary, n2Ary)

		n1Ary = result

	}

	return result
}

func findFirstSignificantDigitIdx(n1Ary []int) int {

	aryLen := len(n1Ary)

	for i:= 0; i < aryLen; i++ {

		if n1Ary[i] > 0 {
			return i
		}
	}

	return -1
}

func multiply(n1Ary []int, n2Ary[]int) []int {

	lenN1Ary := len(n1Ary)
	lenN2Ary := len(n2Ary)
	lenLevels := lenN2Ary
	lenNumPlaces := (lenN1Ary + lenN2Ary) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	intFinalAry := make([]int, lenNumPlaces+1)

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0

	for i := lenN2Ary - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := lenN1Ary - 1; j >= 0; j-- {

			n1 = n1Ary[j]
			n2 = n2Ary[i]
			n3 = (n1 * n2) + carry
			n4 = int(math.Mod(float64(n3), float64(10.00)))

			intMAry[levels][place] = n4

			carry = int(n3 / 10)

			place--
		}

		intMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0
	firstDigitIdx := -1

	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = intFinalAry[j+1]
			n2 = intMAry[i][j]
			n3 = n1 + n2 + carry
			n4 = 0

			if n3 > 9 {
				n4 = int(math.Mod(float64(n3), float64(10.0)))
				carry = n3 / 10

			} else {
				n4 = n3
				carry = 0
			}

			intFinalAry[j+1] = n4
			if n4 != 0 {
				firstDigitIdx = j+1
			}
		}

		if carry > 0 {
			intFinalAry[0] = carry
			firstDigitIdx = 0
		}

	}

	if firstDigitIdx == -1 {
		firstDigitIdx = len(intFinalAry) - 1
	}

	return intFinalAry[firstDigitIdx:]
}

func subtractArys(dividendAry []int, divisorAry []int) (subResult []int, sigVal int) {

	lenDivisorAry := len(divisorAry)

	subResult = make([]int, len(divisorAry))

	n1:=0
	n2:=0
	n3:=0
	carry:=0
	signVal := 1

	for i:=lenDivisorAry -1 ; i >=0; i-- {
		n1 = dividendAry[i]
		n2 = divisorAry[i]

		if n1 < n2 + carry {
			n1 = n1 + 10
			n3 = n1 - (n2 + carry)
			carry = 1
		} else {
			n3 = n1 - (n2 + carry)
			carry = 0
		}

		subResult[i] = n3
	}

	if carry > 0 {
		signVal = -1
	}

	return subResult, signVal
}


func isZeroValAry(n1Ary []int) bool {
	lenAry := len(n1Ary)

	isZeroValue := true

	for i:=0; i < lenAry; i++{
		if n1Ary[i] > 0 {
			isZeroValue = false
			break
		}
	}

	return isZeroValue
}

func compareIntArys(n1Ary []int, n2Ary[]int) int {

	lenAry := len(n1Ary)

	for i:= 0; i < lenAry; i++ {
		if n1Ary[i] > n2Ary[i] {
			return 1
		}

		if n1Ary[i] < n2Ary[i] {
			return -1
		}
	}

	return 0
}

