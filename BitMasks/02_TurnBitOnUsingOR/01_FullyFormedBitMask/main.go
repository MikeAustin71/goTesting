package main

import "fmt"

/*
	This example is designed to 'Turn On' Bit # 2
	using a pre-fabricated Bit Mask created using
	Constants.
*/
const (
	maskBit0 = 1 << iota
	maskBit1
	maskBit2
	maskBit3
	maskBit4
	maskBit5
	maskBit6
	maskBit7
	maskBit8
	maskBit9
	maskBit10
	maskBit11
	maskBit12
	maskBit13
	maskBit14
	maskBit15
)

func main() {
	fmt.Println("maskBit0 =", maskBit0)
	fmt.Println("maskBit1 =", maskBit1)
	fmt.Println("maskBit2 =", maskBit2)
	fmt.Println("maskBit3 =", maskBit3)
	fmt.Println("maskBit4 =", maskBit4)
	fmt.Println("maskBit5 =", maskBit5)
	fmt.Println("maskBit6 =", maskBit6)
	fmt.Println("maskBit7 =", maskBit7)
	fmt.Println("maskBit8 =", maskBit8)
	fmt.Println("maskBit9 =", maskBit9)
	fmt.Println("maskBit10 =", maskBit10)
	fmt.Println("maskBit11 =", maskBit11)
	fmt.Println("maskBit12 =", maskBit12)
	fmt.Println("maskBit13 =", maskBit13)
	fmt.Println("maskBit14 =", maskBit14)
	fmt.Println("maskBit15 =", maskBit15)

	A := 97
	fmt.Println("------------------------")
	fmt.Println("A =", A)
	fmt.Println("maskBit2 = ", maskBit2)
	fmt.Println("A | maskBit2 =", A|maskBit2)

	/* Output
	$ go run main.go
	maskBit0 = 1
	maskBit1 = 2
	maskBit2 = 4
	maskBit3 = 8
	maskBit4 = 16
	maskBit5 = 32
	maskBit6 = 64
	maskBit7 = 128
	maskBit8 = 256
	maskBit9 = 512
	maskBit10 = 1024
	maskBit11 = 2048
	maskBit12 = 4096
	maskBit13 = 8192
	maskBit14 = 16384
	maskBit15 = 32768
	------------------------
	A = 97
	maskBit2 =  4
	A | maskBit2 = 101
	Note: Bit # 2 is Turned On in the result = 101
	*/
}
