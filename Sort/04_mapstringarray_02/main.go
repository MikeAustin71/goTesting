package main

import (
  "fmt"
  "sort"
)

func main() {

  timeSubZones := map[string][]string{
    "America/Indiana" : {
      "America/Indiana/Indianapolis",
      "America/Indiana/Knox",
      "America/Indiana/Marengo",
      "America/Indiana/Petersburg",
      "America/Indiana/Tell_City",
      "America/Indiana/Vevay",
      "America/Indiana/Vincennes",
      "America/Indiana/Winamac"},
    "America/Kentucky" : {
      "America/Kentucky/Louisville",
      "America/Kentucky/Monticello"},
    "America/North_Dakota" : {
      "America/North_Dakota/Beulah",
      "America/North_Dakota/Center",
      "America/North_Dakota/New_Salem"},
    "America/Argentina" : {
      "America/Argentina/Buenos_Aires",
      "America/Argentina/Catamarca",
      "America/Argentina/Cordoba",
      "America/Argentina/La_Rioja",
      "America/Argentina/Mendoza",
      "America/Argentina/Rio_Gallegos",
      "America/Argentina/Salta",
      "America/Argentina/San_Juan",
      "America/Argentina/San_Luis",
      "America/Argentina/Tucuman",
      "America/Argentina/Ushuaia" } }


  keys := make([]string , 0)

  for k := range timeSubZones {

    keys = append(keys, k)

  }

  sort.Strings(keys)

  for i:=0; i < len(keys); i++ {

    subArray, ok := timeSubZones[keys[i]]

    if !ok {
      fmt.Printf("Error: timeSubZones[keys[i]] DOES NOT EXIST!\n" +
        "keys[%v]=%v\n", i+1, keys[i])
      return
    }

    fmt.Println()
    fmt.Println("==============================================================")
    fmt.Printf("SubArray Name: %v\n", keys[i])
    fmt.Println("==============================================================")
    sort.Strings(subArray)

    for j:=0; j < len(subArray); j++ {
      fmt.Printf("%3d. Sub-TimeZone: %v\n", j+1, subArray[j])
    }
    fmt.Println()
  }

}