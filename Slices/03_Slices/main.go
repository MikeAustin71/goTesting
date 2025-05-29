package main

import (
  "fmt"
  "sort"
)

// Courtesy of
// https://www.youtube.com/@anthonygg_
// https://www.youtube.com/watch?v=AL_C9nF_0ss&t=26s

func addUsers(users []string) {

  for _, user := range users {
    fmt.Println(user)
  }
}

func addUsers2(users ...string) {

  for _, user := range users {
    fmt.Println(user)
  }
}

var users = []string{}

func addUser3(user string) {

  users = append(users, user)

}

func addUser4(user ...string) {

  users = append(users, user...)

}

func main() {

  test10()

}

func test01() {

  fmt.Printf("\n\ntest01\n")

  addUsers([]string{"Alice", "Bob", "Foo"})

}

func test02() {

  fmt.Printf("\n\ntest02\n")

  addUsers2("Alice", "Bob", "Foo")

}

func test03() {

  fmt.Printf("\n\ntest03\n")

  addUser3("Alice")
  addUser3("Bob")
  addUser3("Foo")

  fmt.Println(users)

}

func test04() {

  fmt.Printf("\n\ntest04\n")

  addUser3("Alice")
  addUser3("Bob")
  addUser3("Foo")

  fmt.Println(users)

}

func test05() {

  fmt.Printf("\n\ntest05\n")

  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println("Numbers Before: ", numbers)

  fmt.Println("Numbers After: ", removeFromSlice1(numbers, 1))
}

func removeFromSlice1(slice []int, index int) []int {
  // Removes the correct index, but returns an unorder list!
  slice[index] = slice[len(slice)-1]

  return slice[:len(slice)-1]
}

func removeFromSliceWithOrder2(slice []int, index int) []int {
  fmt.Printf("Removing Index: %v - Value: %v\n", index, slice[index])
  // Returns ordered slice!
  return append(slice[:index], slice[index+1:]...)
}

func test06() {

  fmt.Printf("\n\ntest06\n")

  numbers := []int{1, 2, 3, 4, 5}

  fmt.Println("Numbers Before: ", numbers)

  fmt.Println("Numbers After: ", removeFromSliceWithOrder2(numbers, 1))
}

type MySlice []int

func (s MySlice) Remove(index int) []int {
  fmt.Printf("Removing Index: %v - Value: %v\n", index, s[index])
  return append(s[:index], s[index+1:]...)
}

func test07() {
  fmt.Printf("\n\ntest07\n")
  numbers := MySlice{1, 2, 3, 4, 5}
  fmt.Println("Numbers Before: ", numbers)

  numbers = numbers.Remove(1)

  fmt.Println("Numbers After: ", numbers)
}

type Numbers []int

// Len
// Required for sort interface
func (n Numbers) Len() int {
  return len(n)
}

// Swap
// Required for sort interface
func (n Numbers) Swap(i, j int) {
  n[i], n[j] = n[j], n[i]
}

// Less
// Required for sort interface
func (n Numbers) Less(i, j int) bool {
  return n[i] < n[j]
}

func test08() {
  fmt.Printf("\n\ntest08\n")
  numbers := Numbers{1, 10, 4, 9, 3}
  fmt.Println("Numbers Before: ", numbers)

  sort.Sort(numbers)

  fmt.Println("Numbers After: ", numbers)
}

func test09() {
  fmt.Printf("\n\ntest09 byInc\n")

  numbers := Numbers{1, 10, 4, 9, 3}

  fmt.Println("Numbers Before: ", numbers)

  sort.Sort(byInc{numbers})

  fmt.Println("Numbers After: ", numbers)

}

type byInc struct {
  Numbers
}

// Len
// Required for sort interface
func (n byInc) Len() int {
  return len(n.Numbers)
}

// Swap
// Required for sort interface
func (n byInc) Swap(i, j int) {
  n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}

// Less
// Required for sort interface
func (n byInc) Less(i, j int) bool {
  return n.Numbers[i] < n.Numbers[j]
}

// test10
// Only change from byInc is in the
// 'Less' method.
func test10() {
  fmt.Printf("\n\ntest10 - byDec\n")

  numbers := Numbers{1, 10, 4, 9, 3}

  fmt.Println("Numbers Before: ", numbers)

  sort.Sort(byDec{numbers})

  fmt.Println("Numbers After: ", numbers)

}

type byDec struct {
  Numbers
}

// Len
// Required for sort interface
func (n byDec) Len() int {
  return len(n.Numbers)
}

// Swap
// Required for sort interface
func (n byDec) Swap(i, j int) {
  n.Numbers[i], n.Numbers[j] = n.Numbers[j], n.Numbers[i]
}

// Less
// Required for sort interface
// This is the ONLY change from byInc
func (n byDec) Less(i, j int) bool {
  return n.Numbers[i] > n.Numbers[j]
}
