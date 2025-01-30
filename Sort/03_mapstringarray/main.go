package main

import "fmt"

func main() {

  timesubzones := make(map[string][]string, 1)

  timesubzones["America/Indiana"] = []string{
    "America/Indiana/Indianapolis"}

  array, ok := timesubzones["America/Indiana"]

  if !ok {
    fmt.Printf("ERROR America/Indiana\n")
    return
  }

  array = append(array, "America/Indiana/Knox")
  array = append(array, "America/Indiana/Marengo")

  array2 := timesubzones["America/Indiana"]

  fmt.Println()
  fmt.Println("Array2 Listing")

  for i:=0; i < len(array2); i++ {
    fmt.Printf("%3d. Value='%v'\n", i+1, array2[i])
  }

  fmt.Println()
  fmt.Println("Array2 Listing #2")

  timesubzones["America/Indiana"] = array

  array2 = timesubzones["America/Indiana"]

  fmt.Println()
  fmt.Println("Array2 Listing # 3")

  for i:=0; i < len(array2); i++ {
    fmt.Printf("%3d. Value='%v'\n", i+1, array2[i])
  }




}
