package main

import (
  "fmt"
  "sort"
)

var timeZones = []string {
  "Brazil/West",
  "Etc/GMT+11",
  "Europe/Berlin",
  "Pacific/Bougainville",
  "Africa/Abidjan",
  "Europe/Kaliningrad",
  "Africa/Brazzaville",
  "Indian/Christmas",
  "Pacific/Guadalcanal",
  "Africa/Juba",
  "Pacific/Pago_Pago",
  "Africa/Tunis",
  "America/Aruba" }



func main() {

  fmt.Println("02_StringArraySort01")
  fmt.Println("Un-Sorted Time Zones")
  fmt.Println()
  for i:=0; i < len(timeZones); i++ {
    fmt.Printf("%30d. - %v\n", i+1, timeZones[i])
  }

  sort.Strings(timeZones)

  fmt.Println()
  fmt.Println("Sorted Time Zones")
  fmt.Println()
  for i:=0; i < len(timeZones); i++ {
    fmt.Printf("%30d. - %v\n", i+1, timeZones[i])
  }
}
