package main

import "fmt"

func main() {
	/*
		This proves that pointers are passed by value.
		Altering a pointer in a subordinate function
		does NOT alter it in the calling function.
	*/
	str1 := "This is main() String # 1"

	ptrToStr1 := &str1

	err := mainTestFunc01(ptrToStr1)

	if err != nil {
		fmt.Printf("main()\n%v\n",
			err.Error())

		return
	}

	fmt.Println()
	fmt.Println("After mainTestFunc01()")
	fmt.Printf("ptrToStr1= %v\n"+
		"String= %s\n",
		ptrToStr1,
		*ptrToStr1)

}

func mainTestFunc01(locStr *string) error {

	ePrefix := "mainTestFunc01()"

	if locStr == nil {
		return fmt.Errorf("%v\n"+
			"Error: Input paramter 'locStr' is a nil pointer!\n",
			ePrefix)
	}

	fmt.Printf("locStrFirst Pointer: %v\n"+
		"locStrFirst String: %s\n\n",
		locStr,
		*locStr)

	locStrSecond := "This is mainTestFunc01() local string 'Second'."

	locStr = &locStrSecond

	fmt.Printf("locStrSecond Pointer: %v\n"+
		"locStrSecond String: %s\n\n",
		locStr,
		*locStr)

	locStr = nil

	return nil
}
