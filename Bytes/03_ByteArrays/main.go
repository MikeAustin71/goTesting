package main

import "fmt"

func main() {

	b := make([]byte, 5, 9)
	processBytes(b)
}


func processBytes(b []byte) {


	fmt.Println("         Process Byte Array")
	fmt.Println("-------------------------------------")
	fmt.Println("  Length Byte Array: ", len(b))
	fmt.Println("Capacity Byte Array: ", cap(b))
	fmt.Println()

	manipulateBytes2(b)


	fmt.Println("     Process Byte Array2   ")
	fmt.Println("-------------------------------------")
	fmt.Println("  Length Byte Array: ", len(b))
	fmt.Println("Capacity Byte Array: ", cap(b))
	fmt.Println("byte array Output")

	for i:=0; i < len(b); i++ {
		fmt.Printf("%0v. %v \n", i, b[i])
	}

}

func manipulateBytes1(b []byte) {
	fmt.Println("  Manipulate -1- Byte Array")
	fmt.Println("==============================")

	maxLen := len(b)
	fmt.Println("    Original Length: ", maxLen)
	maxLen +=10
	fmt.Println("         New Length: ", maxLen)
	fmt.Printf("\n\n\n")
	for i:=0; i< maxLen; i++ {
		b = append(b, 'X')
	}
	for i:=0; i < len(b); i++ {
		fmt.Printf("%0v. Manipulate bytes copy %v \n", i, b[i])
	}
	fmt.Println("  Actual New Length: ", len(b))
  fmt.Println("Actual New Capacity: ")
}

func manipulateBytes2(b []byte) {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println("  Manipulate -2- Byte Array")
	fmt.Println("==============================")

	maxLen := len(b)
	fmt.Println("    Original Length: ", maxLen)
	maxLen +=10
	b = make([]byte, maxLen, maxLen+10)

	fmt.Println("         New Length: ", maxLen)
	fmt.Printf("\n\n\n")
	for i:=0; i< maxLen; i++ {
		b[i] = 'Z'
	}

	for i:=0; i < len(b); i++ {
		fmt.Printf("%0v. Manipulate-2 bytes copy %v \n", i, b[i])
	}

	fmt.Println("  Actual New Length: ", len(b))
	fmt.Println("Actual New Capacity: ")

}
