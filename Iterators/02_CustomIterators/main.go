package main

import (
	"fmt"
	"iter"
)

type Slice []int

// Version 1
//func (s Slice) Map(
//	transform func(int) int,
//) iter.Seq[int] {
//	return func(yield func(int) bool) {
//		for _, v := range s {
//			if !yield(transform(v)) {
//				return
//			}
//		}
//	}
//}

// Version 2
//func (s Slice) Map(
//  transform func(int) int,
//) iter.Seq[int] {
//  return func(yield func(int) bool) {
//    for _, v := range s {
//
//      transformedValue := transform(v)
//
//      stopIteration := yield(transformedValue)
//
//      if !stopIteration {
//        return
//      }
//    }
//  }
//}

// Version 3
func (s Slice) Map(
	transform func(int) int,
) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for i, v := range s {

			transformedValue := transform(v)

			stopIteration := yield(i, transformedValue)

			if !stopIteration {
				return
			}
		}
	}
}

func main() {

	Test01()
}

func Test01() {

	fmt.Printf("\n\n========================\n" +
		"         Test01\n")

	numbers := Slice{0, 1, 2, 3, 4, 5}

	doubled := numbers.Map(func(i int) int {
		return i * 2
	})

	for i, v := range doubled {
		fmt.Println(i, v)
	}

	fmt.Printf("\n\n -- Running 2nd Version 'doubled' --\n\n")

	// range syntax does the same thing as:
	doubled(func(i int, v int) bool {
		fmt.Println(i, v)
		return true
	})

}
