package main

import (
  "fmt"
)

func main() {

  //testSomeFunc1()
  //testSomeFunc2()
  //testSomeFunc3()
  testSomeFunc4()

}

func testSomeFunc1() {

  originalTestStr1 := "oldString"
  // Strings are immutable
  testStr1 := "oldString"

  testStr2 := someFunc1(testStr1)

  fmt.Println()
  fmt.Println("            Testing someFunc1()")
  fmt.Println("-------------------------------------------")
  fmt.Println("original test 1 str: ", originalTestStr1)
  fmt.Println("Testing immutable strings!")
  fmt.Println("testStr1: ", testStr1)
  fmt.Println("testStr2: ", testStr2)
  fmt.Println("-------------------------------------------")

}

func testSomeFunc2() {
  // Strings are immutable
  originalTestStr1 := "oldString"
  testStr1 := "oldString"

  testStr2 := someFunc2(&testStr1)

  fmt.Println()
  fmt.Println("            Testing someFunc2()")
  fmt.Println("===========================================")
  fmt.Println("Testing string pointers!")
  fmt.Println("original test 1 str: ", originalTestStr1)
  fmt.Println("           testStr1: ", testStr1)
  fmt.Println("           testStr2: ", fmt.Sprintf("%v", testStr2))
  fmt.Println("===========================================")

}

func testSomeFunc3() {
  // Strings are immutable
  originalTestStr1 := "oldString"
  testStr1 := "oldString"

  var newByteArray []byte
  newByteArray = append(newByteArray, testStr1... )

  //testStr2 := someFunc3(&newByteArray)

  fmt.Println()
  fmt.Println("            Testing someFunc3()")
  fmt.Println("===========================================")
  fmt.Println("Testing string pointers!")
  fmt.Println("original test 1 str: ", originalTestStr1)
  fmt.Println("           testStr1: ", testStr1)
  //fmt.Println("           testStr2: ", fmt.Sprintf("%s", string(testStr2[:])))
  // This doesn't work it only takes []byte not *[]byte.
  fmt.Println("===========================================")

}

func testSomeFunc4() {

  originalTestStr1 := "oldString"
  testStr1 := "oldString"
  b := []byte(testStr1)

  testStr2 := someFunc4(b)


  fmt.Println()
  fmt.Println("           Testing someFunc4()")
  fmt.Println("===========================================")
  fmt.Println("Testing byte arrays!")
  fmt.Println("  original test 1 str: ", originalTestStr1)
  fmt.Println("  testStr1 byte array: ", string(b))
  fmt.Println("             testStr2: ", testStr2)
  // This doesn't work it only takes []byte not *[]byte.
  fmt.Println("===========================================")

}

func someFunc1(tstStr string) string {
  // tstStr[0] = 'a' // generates error
  tstStr =  "new string"

  return tstStr
}

func someFunc2(tstStr *string) *string {

  tstStr2 := "new string"

  tstStr = &tstStr2

  return tstStr
}

func someFunc3(tstStr *[]byte) *[]byte {

  newStr := "new string"
  var newByteArray []byte
  newByteArray = append(newByteArray, newStr... )
  tstStr = &newByteArray
  return tstStr
}

func someFunc4(b []byte) string {

  bLen := len(b)
  sameOldString := string(b[:bLen])

  b = []byte("Really New String")

  return sameOldString
}