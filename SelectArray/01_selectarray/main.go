package main

import (
  "fmt"
  "strings"
)

var timeZones = []string {
  "Brazil/West",
  "Asia/Almaty",
  "Etc/GMT+11",
  "Asia/Kolkata",
  "Europe/Berlin",
  "Asia/Yangon",
  "Pacific/Bougainville",
  "Africa/Abidjan",
  "Asia/Oral",
  "Europe/Kaliningrad",
  "Asia/Damascus",
  "Africa/Brazzaville",
  "Indian/Christmas",
  "Asia/Thimphu",
  "Pacific/Guadalcanal",
  "Asia/Qatar",
  "Africa/Juba",
  "Pacific/Pago_Pago",
  "Africa/Tunis",
  "America/Aruba",
  "Asia/Brunei",
  "Indian/Maldives",
  "Asia/Ashgabat"}


type selectTimeZones  []string

func (selTz selectTimeZones) TimeZoneExists(testTz string, useLwrCase bool) bool {

  strLen := len(testTz)

  if strLen == 0 {
    return false
  }

  if useLwrCase{
    testTz = strings.ToLower(testTz)
  }

  for i:=0; i < len(selTz); i++ {

    if useLwrCase {
      if strings.ToLower(selTz[i]) == testTz {
        return true
      }
    } else {
      if selTz[i] == testTz {
        return true
      }
    }
  }

  return false
}


func main() {
  testTz := selectTimeZones(timeZones)

  test1:=  "Europe/Kaliningrad"

  result1 := testTz.TimeZoneExists(test1, false)

  test2 := "Pacific/Bougainville"

  result2 := testTz.TimeZoneExists(test2, false)

  test3 := "pacific/bougainville"

  result3 := testTz.TimeZoneExists(test3, true)

  test4 := "IDoNotExist/TimeZone"

  result4 := testTz.TimeZoneExists(test4, false)

  test5 := "idonotexist/timezone"

  result5 := testTz.TimeZoneExists(test5, true)


  fmt.Println("01_selectarray - selectTimeZones")
  fmt.Println("================================")
  fmt.Println("  test1: ", test1)
  fmt.Println("result1: ", result1)
  fmt.Println("--------------------------------")
  fmt.Println("  test2: ", test2)
  fmt.Println("result2: ", result2)
  fmt.Println("--------------------------------")
  fmt.Println("  test3: ", test3)
  fmt.Println("result3: ", result3)
  fmt.Println("--------------------------------")
  fmt.Println("  test4: ", test4)
  fmt.Println("result4: ", result4)
  fmt.Println("--------------------------------")
  fmt.Println("  test5: ", test5)
  fmt.Println("result5: ", result5)
  fmt.Println("--------------------------------")

  for i:=0; i < len(timeZones); i++ {
    fmt.Printf("\t%3d.\t%v\n", i+1, timeZones[i])
  }

}