package main

import (
  "fmt"
  "sort"
)

func main() {

  var m = map[int]string{
    0: "str1",
    1: "str2",
    2: "str3",
    3: "str4",
    4: "str5",
    5: "str6",
    6: "str7",
    7: "str8",
    8: "str9",
    9: "str10",
  }

  fmt.Println("Un-Ordered List of Map Entries:")

  for ky, mval := range m {
    fmt.Printf("Key Value: %v\tMap Value: %v\n",
      ky, mval)
  }

  fmt.Println()
  fmt.Println()
  fmt.Println("Ordered List of Map Entries:")
  fmt.Println()
  keys := make([]int, 0)

  for k := range m {

    keys = append(keys, k)

  }

  sort.Ints(keys)

  for i:=0; i < len(keys); i++ {
    mVal, ok := m[keys[i]]

    if !ok {
      fmt.Printf("Error: Could Not Locate keys[%v]: %v\n", i, keys[i])
      return
    }

    fmt.Printf("Key Value: %v \tMap Value: %v \n",
      keys[i], mVal)
  }

}

/* Output

Un-Ordered List of Map Entries:
Key Value: 6    Map Value: str7
Key Value: 7    Map Value: str8
Key Value: 8    Map Value: str9
Key Value: 5    Map Value: str6
Key Value: 9    Map Value: str10
Key Value: 0    Map Value: str1
Key Value: 1    Map Value: str2
Key Value: 2    Map Value: str3
Key Value: 3    Map Value: str4
Key Value: 4    Map Value: str5


Ordered List of Map Entries:

Key Value: 0    Map Value: str1
Key Value: 1    Map Value: str2
Key Value: 2    Map Value: str3
Key Value: 3    Map Value: str4
Key Value: 4    Map Value: str5
Key Value: 5    Map Value: str6
Key Value: 6    Map Value: str7
Key Value: 7    Map Value: str8
Key Value: 8    Map Value: str9
Key Value: 9    Map Value: str10
 */
