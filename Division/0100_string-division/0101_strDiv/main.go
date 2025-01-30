package main

import (
	"fmt"
)

func main()  {

	dividendTop := 425
	divisorBottom := 25

	dividendAry := []int{4,2,5}
	divisorAry  := []int{0,2,5}
	zeroAry     := []int{0,0,0}


	fmt.Printf("Subtract divisor %v from dividend %v \n", divisorBottom, dividendTop)
	for i:=0; i< 18; i++ {

		subResult, signVal := subtractArys(dividendAry, divisorAry)

		fmt.Printf("subResult= %v SignVal= %v Iteration %v \n",subResult, signVal, i+1)
		dividendAry = subResult

		compResult1 := compareIntArys(subResult, zeroAry)

		compResult2 := compareIntArys(subResult, divisorAry)

		if compResult1 == 0 {
			fmt.Println("Result Array = 0")
			break
		}

		if compResult2 == -1 {
			fmt.Println("Result Less Than Divisor")
			break
		}
	}

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Successful Completion")
}


func div01(dividendTop int, divisorBottom int) (quotient int, remainder int) {

	quotient = dividendTop / divisorBottom

	remainder = dividendTop - (quotient * divisorBottom)


	return quotient, remainder
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

