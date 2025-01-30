package main

import (
  "fmt"
  "sort"
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



type TimeZoneDataDto struct {
  MajorGroup string
  TzName string
  TzValue string
  TzClass int       // 0 = Unknown
                    // 1 = Canonical
                    // 2 = Alias
                    // 3 = Sub-Group
}

func (tzDataDto TimeZoneDataDto) New(
  majorGroup,
  tzName,
  tzValue string,
  tzClass int) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataDto.New() - ERROR:\n"

  if tzClass < 1 || tzClass > 3 {
    return TimeZoneDataDto{},
    fmt.Errorf(ePrefix + "Input Parameter tzClass is out of bounds and INVALID!\n" +
      "Valid values are 1-3!\ntzClass='%v'", tzClass)
  }

  tzDto := TimeZoneDataDto{}
  tzDto.MajorGroup = majorGroup
  tzDto.TzName = tzName
  tzDto.TzValue = tzValue
  tzDto.TzClass = tzClass

  return tzDto, nil
}

type byTzDtoName []TimeZoneDataDto

func (byTzDtoName byTzDtoName) Len() int {
  return len(byTzDtoName)
}

func (byTzDtoName byTzDtoName) Swap(i, j int) {
  byTzDtoName[i], byTzDtoName[j] = byTzDtoName[j], byTzDtoName[i]
}

func (byTzDtoName byTzDtoName) Less(i, j int) bool {

  if byTzDtoName[i].MajorGroup == byTzDtoName[j].MajorGroup {
    return byTzDtoName[i].TzName < byTzDtoName[j].TzName
  }

  return byTzDtoName[i].MajorGroup < byTzDtoName[j].MajorGroup
}



func main() {

  tzDtoArray := make([]TimeZoneDataDto,0, 20)

  for i:=0; i < len(timeZones); i++ {

    name := timeZones[i]

    sArray := strings.Split(name, "/")

    if len(sArray) < 2 {
      fmt.Printf("ERROR: Split returned a length of %v\n" +
        "name='%v'\n", len(sArray), name)
      return
    }

    tzDto, err := TimeZoneDataDto{}.New(sArray[0], name, name, 1)

    if err != nil {

        fmt.Printf("Error returned by TimeZoneDataDto{}.New(sArray[0], name, name, 1)\n" +
          "sArray[0]='%v'\nname='%v'\nError='%v'\n",
          sArray[0], name, err.Error())

        return
    }

    tzDtoArray = append(tzDtoArray, tzDto)

  }

  fmt.Println("Custom Sort Result")
  fmt.Println()

  sort.Sort(byTzDtoName(tzDtoArray))

  for j:=0; j < len(tzDtoArray); j++ {

    fmt.Printf("%3d. Group: %v   \tName: %v\tValue: %v\tClass: %v\n",
      j+1, tzDtoArray[j].MajorGroup, tzDtoArray[j].TzName, tzDtoArray[j].TzValue, tzDtoArray[j].TzClass )

  }
}
