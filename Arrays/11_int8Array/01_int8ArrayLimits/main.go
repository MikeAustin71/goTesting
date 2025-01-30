package main

import "math/big"

func main() {

	_ = Uint64ArrayLimits()
}

func Int64ArrayLimits() []int8 {

	var i8Array []int8

	// Signed Int 64:
	//From:
	// −9,223,372,036,854,775,808
	//            to
	// 9,223,372,036,854,775,807
	//
	// From −(2^63) to +2^63 − 1
	var aryElements int64

	aryElements = 50000

	i8Array = make([]int8, aryElements)

	return i8Array

}

func Uint64ArrayLimits() []int8 {

	var i8Array []int8

	i8Array = make([]int8, 9)

	// Unsigned Int 64:
	// From 0 to 18,446,744,073,709,551,615,
	//which equals 2^64 − 1
	var aryElements uint64

	aryElements = 50000

	i8Array = make([]int8, aryElements)

	return i8Array
}

func BigIntArrayLimits() ([]int8, *big.Int) {

	var bigNum *big.Int

	bigNum = big.NewInt(50000)

	var i8Array []int8

	// This does NOT WORK
	// i8Array = make([]int8, bigNum)

	var aryElements uint64

	aryElements = 50000

	i8Array = make([]int8, aryElements)

	return i8Array, bigNum
}
