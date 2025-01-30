package main

import (
	"errors"
	"fmt"
)

// arr := []int{1,2,3,4,5}
// arr[:2]         [1,2]
// arr[2:])        [3,4,5]

func main() {

	example03()
}

func example01() {

	t := []int{1, 2, 3, 4, 5}

	index := 2
	value := 9

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("      Running example01()         \n")
	fmt.Printf("----------------------------------\n")
	fmt.Printf("[]int{1, 2, 3, 4, 5}\n")
	fmt.Printf("[]int array length = '%v'\n",
		len(t))

	fmt.Printf("Inserting '%v' at int[%v]\n",
		value,
		index)

	fmt.Println()

	t, err := insert(t, index, value)

	if err == nil {
		fmt.Println(t) // [1 2 9 3 4 5]
	} else {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf(" Successful Completion example01()\n")
	fmt.Printf("----------------------------------\n")
	fmt.Println()
	fmt.Println()

}

func example02() {

	t := []int{1, 2, 3, 4, 5}

	index := 0
	value := 9

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("      Running example02()         \n")
	fmt.Printf("----------------------------------\n")
	fmt.Printf("[]int{1, 2, 3, 4, 5}\n")
	fmt.Printf("[]int array length = '%v'\n",
		len(t))

	fmt.Printf("Inserting '%v' at int[%v]\n",
		value,
		index)

	fmt.Println()

	t, err := insert(t, index, value)

	if err == nil {
		fmt.Println(t) // [9 1 2 3 4 5]
	} else {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf(" Successful Completion example02()\n")
	fmt.Printf("----------------------------------\n")
	fmt.Println()
	fmt.Println()

}

func example03() {

	t := []int{1, 2, 3, 4, 5}

	index := 4
	value := 9

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf("      Running example03()         \n")
	fmt.Printf("----------------------------------\n")
	fmt.Printf("[]int{1, 2, 3, 4, 5}\n")
	fmt.Printf("[]int array length = '%v'\n",
		len(t))

	fmt.Printf("Inserting '%v' at int[%v]\n",
		value,
		index)

	fmt.Println()

	t, err := insert(t, index, value)

	if err == nil {

		fmt.Printf("New Array Length: '%v'\n\n",
			len(t))

		fmt.Println(t) // 1 2 9 3 4 5]

	} else {
		fmt.Println(err)
	}

	fmt.Println()
	fmt.Printf("----------------------------------\n")
	fmt.Printf(" Successful Completion example03()\n")
	fmt.Printf("----------------------------------\n")
	fmt.Println()
	fmt.Println()

}

func insert(
	orig []int,
	index int,
	value int) (
	[]int,
	error) {

	if index < 0 {
		return nil, errors.New("Index cannot be less than 0")
	}

	if index >= len(orig) {
		return append(orig, value), nil
	}

	orig = append(orig[:index+1], orig[index:]...)

	fmt.Printf("Value of target index\n"+
		"before replacement.\n"+
		" orig[%v] = %v\n\n",
		index,
		orig[index])

	orig[index] = value

	return orig, nil
}
