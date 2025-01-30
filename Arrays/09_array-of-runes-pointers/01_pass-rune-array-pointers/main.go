package main

import "fmt"

func main() {

	var runeArray []rune

	sourceArray := []rune("Hello World, how are you!")

	runeArray = make([]rune, len(sourceArray))

	copy(runeArray, sourceArray)

	fmt.Printf("\n\n")
	fmt.Printf("--------------------------------------------\n")
	fmt.Printf("Original 'runeArray' value in main()\n\n")
	fmt.Printf("%v\n",
		string(runeArray))
	fmt.Printf("Original Rune Array Length: %v\n",
		len(runeArray))
	fmt.Printf("Original Rune Array Capacity: %v\n",
		cap(runeArray))
	fmt.Printf("--------------------------------------------\n\n")

	methodName := "sayHelloTwo"

	err := sayHelloTwo(&runeArray)

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	fmt.Println()
	fmt.Printf("main - %v Call\n",
		methodName)

	fmt.Println("SUCCESSFUL EXECUTION!")
	fmt.Println()
	fmt.Printf("New Rune Array Value:\n"+
		"%v\n", string(runeArray))
	fmt.Printf("New Rune Array Length: %v\n",
		len(runeArray))
	fmt.Printf("New Rune Array Capacity: %v\n",
		cap(runeArray))
}

// This does NOT allow you to modify someRuneArray!
func sayHello(someRuneArray []rune) (err error) {

	if someRuneArray == nil {
		err = fmt.Errorf("sayHello()\n" +
			"Error: Input parameter 'someRuneArray' has a value of nil!\n")

		return err
	}

	lenOfArray := len(someRuneArray)

	if lenOfArray == 0 {
		err = fmt.Errorf("sayHello()\n" +
			"Error: Input parameter 'someRuneArray' is a zero length array!\n")

		return err
	}

	for i := 0; i < lenOfArray; i++ {

		if someRuneArray[i] == 0 {
			err = fmt.Errorf("sayHello()\n"+
				"Error: 'someArray' is invalid!\n"+
				"someRuneArray[%v] has a rune value of zero!\n",
				i)

			return err
		}
	}

	someRuneArray = nil

	newArray := []rune("Hello Mike")

	lenOfArray = len(newArray)

	someRuneArray = make([]rune, lenOfArray)

	for i := 0; i < lenOfArray; i++ {
		someRuneArray[i] = newArray[i]
	}

	/*
		charsCopied :=
			copy(
				someRuneArray,
				newArray)

		if charsCopied != lenOfArray {
			err = fmt.Errorf("sayHello()\n" +
				"Error: Characters copied is invalid!\n" +
				"The number of characters copied to 'someRuneArray'\n" +
				"does NOT match the number of characters in 'newArray'.\n" +
				"Number of characters copied to someRuneArray = '%v'\n" +
				"Number of characters in newArray = '%v'\n",
				charsCopied,
				lenOfArray)

		}
	*/
	return err
}

func sayHelloTwo(someRuneArray *[]rune) (err error) {

	if someRuneArray == nil {
		err = fmt.Errorf("sayHelloTwo()\n" +
			"Error: Input parameter '*someRuneArray' is invalid!\n" +
			"The pointer is a 'nil' value!\n")

		return err
	}

	if *someRuneArray == nil {
		err = fmt.Errorf("sayHelloTwo()\n" +
			"Error: Input parameter 'someRuneArray' is invalid!\n" +
			"Concrete 'someRuneArry' is 'nil'!\n")

		return err
	}

	lenOfArray := len(*someRuneArray)

	if lenOfArray == 0 {
		err = fmt.Errorf("sayHelloTwo()\n" +
			"Error: Input parameter 'someRuneArray' is invalid!\n" +
			"Concrete array 'someRuneArray' is a zero length array!\n")

		return err
	}

	var newArray []rune

	newArray = []rune("Hello Mike")

	lenOfArray = len(newArray)

	*someRuneArray = make([]rune, lenOfArray)

	charsCopied :=
		copy(
			*someRuneArray,
			newArray)

	if charsCopied != lenOfArray {
		err = fmt.Errorf("sayHelloTwo()\n"+
			"Error: Characters copied is invalid!\n"+
			"The number of characters copied to 'someRuneArray'\n"+
			"does NOT match the number of characters in 'newArray'.\n"+
			"Number of characters copied to someRuneArray = '%v'\n"+
			"Number of characters in newArray = '%v'\n",
			charsCopied,
			lenOfArray)

	}

	return err
}
