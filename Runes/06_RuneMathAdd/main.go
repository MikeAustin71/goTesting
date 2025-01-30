package main

import "fmt"

func main() {

	OneSubOne(
		'5',
		'8')

	//add1 := []rune("9999")
	//add2 := []rune("90000")
	//
	//AllByAllMessy(
	//	add1,
	//	add2)
}

func AllByAll(
	add1 []rune,
	add2 []rune) {

	fmt.Println("AllByAll()")
	fmt.Printf("add1  = %v\n", string(add1))
	fmt.Printf("add2  = %v\n", string(add2))

	lenAddends := len(add1)

	total := make([]rune, lenAddends)

	isCarry := false

	var factor rune

	for i := lenAddends - 1; i > -1; i-- {

		if isCarry {
			factor = 47
			isCarry = false
		} else {
			factor = 48
		}

		total[i] = add1[i] + add2[i] - factor

		if total[i] > '9' {
			total[i] -= 10
			isCarry = true
		}

	}

	if isCarry {
		total = append(
			[]rune{'1'},
			total...)
	}

	fmt.Printf("total = %v\n", string(total))

	return
}

func Magnitude(
	num1 []rune,
	num2 []rune) {

	// Both rune arrays are assumed to contain
	// numeric digits 0-9 inclusive

	lenArray := len(num1)
	num1Magnitude := 0
	var numAdj rune
	numAdj = 48
	factor := 1

	for i := lenArray - 1; i > -1; i-- {

		factor = factor * factor

		num1Magnitude += int(num1[i]-numAdj) * factor
	}

}

func AllByAllMessy(
	add1 []rune,
	add2 []rune) {

	fmt.Println("AllByAllMessy()")
	fmt.Printf("add1  = %v\n", string(add1))
	fmt.Printf("add2  = %v\n", string(add2))

	idxAdd1 := len(add1) - 1
	idxAdd2 := len(add2) - 1

	maxLastIdx := idxAdd1

	if idxAdd2 > maxLastIdx {
		maxLastIdx = idxAdd2
	}

	var num1, num2, factor rune

	isCarry := false

	total := make([]rune, maxLastIdx+1)

	for i := maxLastIdx; i > -1; i-- {

		if idxAdd1 < 0 {
			num1 = '0'
		} else {
			num1 = add1[idxAdd1]
		}

		idxAdd1--

		if idxAdd2 < 0 {
			num2 = '0'
		} else {
			num2 = add2[idxAdd2]
		}

		idxAdd2--

		if isCarry {
			factor = 47
			isCarry = false
		} else {
			factor = 48
		}

		total[i] = num1 + num2 - factor

		if total[i] > '9' {
			total[i] -= 10
			isCarry = true
		}

	}

	if isCarry {
		total = append(
			[]rune{'1'},
			total...)
	}

	fmt.Printf("total = %v\n", string(total))

}

func OneSubOne(
	minuend rune,
	subtrahend rune) {

	difference := minuend - subtrahend + 48

	var factor rune

	if difference < '0' {
		factor = 10
	} else {
		factor = 0
	}

	fmt.Printf("OneSubOne\n")
	fmt.Printf("minuend = %v\n", string(minuend))
	fmt.Printf("subtrahend = %v\n", string(subtrahend))

	fmt.Printf("difference = %d\n", difference)

	fmt.Printf("difference = %v\n", string(difference))
	fmt.Println()

	fmt.Printf("difference + %d = %d\n",
		factor,
		difference+factor)

	fmt.Printf("difference + %d = %v\n",
		factor,
		string(difference+factor))

	fmt.Println()

}

func OneByOne() {

	num1 := '5'
	num2 := '4'
	num3 := num1 + num2

	fmt.Printf("OneByOne\n")
	fmt.Printf("num3 = %d\n", num3)

	fmt.Printf("num3 = %v\n", string(num3))
	fmt.Println()

	fmt.Printf("num3 - 48 = %d\n", num3-48)

	fmt.Printf("num3 - 48 = %v\n", string(num3-48))
	fmt.Println()

}

func OneByOne2() {

	num1 := '0'
	num2 := '1'
	var num3 rune
	num3 = num1 + num2 - 48
	isCarry := false

	if num3 > '9' {
		num3 -= 10
		isCarry = true
	}

	fmt.Printf("OneByOne\n")
	fmt.Printf("num1 = %v\n", string(num1))
	fmt.Printf("num2 = %v\n", string(num2))
	fmt.Printf("-----------\n")
	fmt.Printf("num3 = %d\n", num3)

	fmt.Printf("num3 = %v\n", string(num3))
	if isCarry {
		fmt.Printf("Carry == 1\n")
	}
	fmt.Println()

}

func ThreeByThree() {

	nums1 := make([]rune, 3)
	nums1[0] = '4'
	nums1[1] = '4'
	nums1[2] = '5'

	nums2 := make([]rune, 3)
	nums2[0] = '3'
	nums2[1] = '3'
	nums2[2] = '5'

	results := make([]rune, 4)
	results[0] = '0'
	results[1] = '0'
	results[2] = '0'
	results[3] = '0'

	carry := make([]rune, 1)
	carry[0] = '0'

	fmt.Printf("Starting nums= %v\n",
		string(nums1))

	for i := 2; i > -1; i-- {

		if carry[0] == '1' {
			carry[0] = '0'
			nums2[i] = nums1[i] + 1
		}

		if nums1[i] > '9' {
			nums1[i] = '0'

			carry[0] = carry[0] + 1

		}
	}

	if carry[0] == '1' {
		nums1 = append(carry, nums1...)
	}

	fmt.Printf("Ending nums= %v\n",
		string(nums1))

}
