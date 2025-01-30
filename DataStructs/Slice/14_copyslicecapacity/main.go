package main

import "fmt"

func main() {

	sliceToCopy := []rune("12345678901234567890")

	sliceTarget := make([]rune, 10, 50)

	runesCopied := copy(sliceTarget, sliceToCopy)

	fmt.Println()
	fmt.Println("          main()")
	fmt.Println("  Testing Slice Capacity")
	fmt.Println("------------------------------------------")
	fmt.Printf("       Slice To Copy: '%v'\n",
		string(sliceToCopy))
	fmt.Printf("Slice To Copy Length: %v\n",
		len(sliceToCopy))
	fmt.Printf("        Runes Copied: %v\n",
		runesCopied)
	fmt.Printf("        Target Slice: %v\n",
		string(sliceTarget))
	fmt.Printf(" Target Slice Length: %v\n",
		len(sliceTarget))
	fmt.Println("------------------------------------------")
	fmt.Printf("This proves that setting a larger\n" +
		"capacity WILL NOT allow one to copy\n" +
		"past the length of the target array.\n")
	fmt.Println("------------------------------------------")
	fmt.Println()
}
