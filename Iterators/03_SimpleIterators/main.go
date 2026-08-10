package main

import (
	"fmt"
	"math"
)

func main() {

	test01()
	//test02()
	//test03()

}

func test01() {

	var limit uint64 = 6
	var count uint64 = 0
	var initialCountNum uint64 = 0

	fmt.Printf("\ntest01 - Start Loop\n")
	fmt.Printf("Initial 'limit' value = %d\n\n", limit)

	for i := range limit {

		count++

		if count == 1 {
			initialCountNum = i
		}

		fmt.Printf("i: %d\n", i)

	}

	fmt.Printf("\nExpected 6-Iterations\n")
	fmt.Printf("Actual Count %d through %d\n\n", initialCountNum, limit-1)

	/*

		test01 - Start Loop
		Initial 'limit' value = 6

		i: 0
		i: 1
		i: 2
		i: 3
		i: 4
		i: 5

		Expected 6-Iterations
		Actual Count 0 through 5

	*/
}

func test02() {

	var limit uint64 = 6
	var count uint64 = 0
	var initialCountNum uint64 = 0

	fmt.Printf("\ntest02 - Start Loop\n")
	fmt.Printf("Initial 'limit' value = %d\n\n", limit)

	for range limit {

		count++

		if count == 1 {
			initialCountNum = 0
		}

		fmt.Printf("index: %d\n", count-1)

	}

	fmt.Printf("\nExpected 6-Iterations\n")
	fmt.Printf("Actual Iterations %d through %d\n", initialCountNum, count-1)
	fmt.Printf("Actual Loop Count: %d\n", count)
	fmt.Printf("Ending Value of 'limit': %d\n\n", limit)

	/* -- OUTPUT --
	test02 - Start Loop
	Initial 'limit' value = 6

	index: 0
	index: 1
	index: 2
	index: 3
	index: 4
	index: 5

	Expected 6-Iterations
	Actual Iterations 0 through 5
	Actual Loop Count: 6
	Ending Value of 'limit': 6

	*/

}

func test03() {

	var limit uint64 = 4048
	var count uint64 = 0
	var initialCountNum uint64 = 0

	fmt.Printf("\ntest03 - Start Loop\n")
	fmt.Printf("Initial 'limit' value = %d\n\n", limit)

	for range limit {

		count++

		if count == 1 {
			initialCountNum = 0
		}

		fmt.Printf("index: %d\n", count-1)

		if count == 12 {
			break
		}

	}

	fmt.Printf("\nExpected 6-Iterations\n")
	fmt.Printf("Actual Iterations %d through %d\n", initialCountNum, count-1)
	fmt.Printf("Actual Loop Count: %d\n", count)
	fmt.Printf("Ending Value of 'limit': %d\n\n", limit)

	/* -- OUTPUT --
	test03 - Start Loop
	Initial 'limit' value = 4048

	index: 0
	index: 1
	index: 2
	index: 3
	index: 4
	index: 5
	index: 6
	index: 7
	index: 8
	index: 9
	index: 10
	index: 11

	Expected 6-Iterations
	Actual Iterations 0 through 11
	Actual Loop Count: 12
	Ending Value of 'limit': 4048

	*/

}

func test04() {

	var count uint64 = 0
	var initialCountNum uint64 = 0

	fmt.Printf("\ntest03 - Start Loop\n")
	fmt.Printf("Initial 'limit' value = %d\n\n", math.MaxUint32)

	for range math.MaxUint32 {

		count++

		if count == 1 {
			initialCountNum = 0
		}

		fmt.Printf("index: %d\n", count-1)

		if count == 12 {
			break
		}

	}

	fmt.Printf("\nExpected 6-Iterations\n")
	fmt.Printf("Actual Iterations %d through %d\n", initialCountNum, count-1)
	fmt.Printf("Actual Loop Count: %d\n", count)

}
