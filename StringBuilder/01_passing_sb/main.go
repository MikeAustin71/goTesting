package main

import (
	"fmt"
	"strings"
)

// main -
// Conclusion from testing.
// The only way to use strings.Builder in calls to subsidiary methods is to
// pass a pointer to the strings.Builder.
func main() {

	sb := strings.Builder{}

	err := subMethod01(&sb)

	if err != nil {
		fmt.Printf("%v",
			err.Error())

		return
	}

	err = subMethod02(&sb)

	if err != nil {
		fmt.Printf("%v",
			err.Error())

		return
	}

	fmt.Printf("#1 - Writing Contents of 'sb-With Pointer'\n\n")

	fmt.Printf(sb.String() + "\n")

	err = subMethod03(&sb)

	if err != nil {
		fmt.Printf("%v",
			err.Error())

		return
	}

	err = subMethod04(&sb)

	if err != nil {
		fmt.Printf("%v",
			err.Error())

		return
	}

	fmt.Printf("#2 - Writing Contents of 'sb-With Pointer'\n\n")

	fmt.Printf(sb.String() + "\n")

	fmt.Printf("01_passing_sb-main()\n" +
		"Successful Completion!!!\n")

	return
}

func subMethod01(
	strBuilder *strings.Builder) error {

	strBuilder.WriteString("Hello from subMethod01\n")

	return nil
}

func subMethod02(
	strBuilder *strings.Builder) error {

	strBuilder.WriteString("Hello from subMethod02\n")

	return nil
}

func subMethod03(
	strBuilder *strings.Builder) error {

	strBuilder.WriteString("Hello from subMethod03\n")

	return nil
}

func subMethod04(
	strBuilder *strings.Builder) error {

	strBuilder.WriteString("Hello from subMethod04\n")

	return nil
}
