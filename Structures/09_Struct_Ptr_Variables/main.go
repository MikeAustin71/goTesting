package main

import "fmt"

func main() {
  BaseRun01()
}

func BaseRun01() {
  origA := 1
  origB := 20
  fmt.Printf("\n\n       BaseRun01\n")
  fmt.Printf("-------------------------------\n")
  fmt.Printf("Original Before Modification:\n")
  fmt.Printf("origA-Pointer: %v  origA-Value %v\n", &origA, origA)
  fmt.Printf("origB-Pointer: %v  origB-Value %v\n", &origB, origB)

  Caller02(&origA, &origB)

  fmt.Printf("\n\n        BaseRun01\n")
  fmt.Printf("-------------------------------\n")
  fmt.Printf("Recap - After Modification\n")
  fmt.Printf("origA-Pointer: %v  origA-Value %v\n", &origA, origA)
  fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", &origB, origB)

}

func Caller02(iPtr1 *int, iPtr2 *int) {
  fmt.Printf("\n\n-------------------------------\n")
  fmt.Printf("    Running Caller 02\n")
  fmt.Printf("-------------------------------\n")

  var mix1 *int
  var mix2 *int

  mix1 = iPtr1
  mix2 = iPtr2

  test01(mix1, mix2)
}

func test01(iPtr1 *int, iPtr2 *int) {

  fmt.Printf("\n\n-------------------------------\n")
  fmt.Printf("         Running test01\n")
  fmt.Printf("Initial Values:\n")
  fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", iPtr1, *iPtr1)
  fmt.Printf("iPtr2-Pointer: %v  iPtr2-Value %v\n", iPtr2, *iPtr2)

  xStr := "\n\ntest01 - Modified Values:\n"
  xStr += "*iPtr1+=10\n"
  xStr += "*iPtr2+=20\n\n"
  fmt.Printf("%v", xStr)

  *iPtr1 += 10
  *iPtr2 += 20

  fmt.Printf("\ntest01 Modified Values:\n")
  fmt.Printf("iPtr1-Pointer: %v  iPtr1-Value %v\n", iPtr1, *iPtr1)
  fmt.Printf("iPtr2-Pointer: %v  iPtr2-Value %v\n\n", iPtr2, *iPtr2)

}
