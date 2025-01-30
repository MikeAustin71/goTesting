package main

import (
	"fmt"
	"math/big"
)

func main() {

	ByteArrayTest01()
}

func ByteArrayTest01() {

	funcName := "ByteArrayTest01()"

	var bigI01 *big.Int

	var int64Num int64 = 10005

	bigI01 = big.NewInt(int64Num)

	var byteArray []byte

	// Bytes returns the absolute value of
	// x as a big-endian byte slice.
	byteArray = bigI01.Bytes()

	fmt.Printf("Function: %v\n",
		funcName)

	fmt.Printf("Original value %v\n",
		int64Num)

	lenByteArray := len(byteArray)

	fmt.Printf("Length of Byte Array = '%v'\n",
		lenByteArray)

	fmt.Printf("Retrieved big-endian byte array follows!\n\n")

	for i := 0; i < lenByteArray; i++ {

		fmt.Printf("byteArray[%v] = '%v'\n",
			i,
			byteArray[i])

	}

	// Original Number 10,005
	// byteArray[0] = '39' 256^1 = 256 x 39 = 9,984
	// byteArray[1] = '21' 256^0 = 1 x 21 = 21
	// 9,984 + 21 = 10,005
	// 8-bits per byte = 2^8 = 256

	// Output

	/*

		Function: ByteArrayTest01()
		Original value 10,005
		Length of Byte Array = '2'
		Retrieved big-endian byte array follows!

		byteArray[0] = '39'
		byteArray[1] = '21'

	*/

	fmt.Printf("\n\nEnd Of %v\n",
		funcName)
}
