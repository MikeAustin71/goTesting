package main

import (
  "errors"
  "fmt"
  "sort"
  "strconv"
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
  "Asia/Baghdad",
  "Etc/GMT+3",
  "GMT+0",
  "Africa/Abidjan",
  "Asia/Oral",
  "Europe/Kaliningrad",
  "Asia/Damascus",
  "Africa/Brazzaville",
  "Etc/GMT+8",
  "Indian/Christmas",
  "Etc/GMT+12",
  "Antarctica/DumontDUrville",
  "Etc/GMT-1",
  "America/Cancun",
  "Asia/Thimphu",
  "Etc/GMT-4",
  "Pacific/Guadalcanal",
  "Asia/Qatar",
  "America/Belize",
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
  TzSortName string
  TzValue string
  TzClass int       // 0 = Unknown
  // 1 = Canonical
  // 2 = Alias
  // 3 = Sub-Group
}

func (tzDataDto TimeZoneDataDto) New(
  majorGroup,
  tzName,
  tzSortName,
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
  tzDto.TzSortName = tzSortName
  tzDto.TzValue = tzValue
  tzDto.TzClass = tzClass

  return tzDto, nil
}

func (tzDataDto *TimeZoneDataDto) CopyOut() TimeZoneDataDto {

  newTzDto := TimeZoneDataDto{}
  newTzDto.MajorGroup = tzDataDto.MajorGroup
  newTzDto.TzName = tzDataDto.TzName
  newTzDto.TzSortName = tzDataDto.TzSortName
  newTzDto.TzValue = tzDataDto.TzValue
  newTzDto.TzClass = tzDataDto.TzClass
  
  return newTzDto
}

type TimeZoneDataCollection struct {
  tzDataDtos  []TimeZoneDataDto
}

func (tzDtoCol *TimeZoneDataCollection) AddNewTzDataDto(
  majorGroup,
  tzName,
  tzSortName,
  tzValue string,
  tzClass int) error {

  tzDto, err := TimeZoneDataDto{}.New(majorGroup,tzName, tzSortName, tzValue, tzClass)

  if err != nil {
    return fmt.Errorf("TimeZoneDataCollection.AddNewTzDataDto()\n" +
      "ERROR returned by TimeZoneDataDto{}.New(majorGroup,tzName, tzSortName, tzValue, tzClass)\n" +
      "majorGroup='%v'\ttzName='%v'\n" +
      "tzSortName='%v'\ttzValue='%v'\ttzClass='%v'\nError='%v'\n",
      majorGroup,tzName, tzSortName, tzValue, tzClass, err.Error())
  }

  if tzDtoCol.tzDataDtos == nil {
    tzDtoCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)
  }

  tzDtoCol.tzDataDtos = append(tzDtoCol.tzDataDtos, tzDto)

  return nil
}

func (tzDtoCol *TimeZoneDataCollection) GetNumberOfDtos() int {
  return len(tzDtoCol.tzDataDtos)
}

func (tzDtoCol *TimeZoneDataCollection) PeekAtIndex(index int) (TimeZoneDataDto, error) {

  ePrefix := "TimeZoneDataCollection.PeekAtIndex() "
  
  if tzDtoCol.tzDataDtos == nil {
    tzDtoCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)
  }

  lenTzDataDtos := len(tzDtoCol.tzDataDtos)
  
  if lenTzDataDtos == 0 {
    return TimeZoneDataDto{}, 
    errors.New(ePrefix + "ERROR: Time Zone Dto Collection is EMPTY!\n")
  }
  
  
  if index < 0 || index > (lenTzDataDtos - 1) {
    return TimeZoneDataDto{}, 
      fmt.Errorf(ePrefix + "ERROR: Input parameter 'index' is out-of-bounds!\n" +
        "index='%v'\t" +
        "Length Of tzDataDtos array is %v\n", index, lenTzDataDtos)
  }
  
  
  return tzDtoCol.tzDataDtos[index].CopyOut(), nil
}

func (tzDtoCol *TimeZoneDataCollection) SortTzDtosByMjrGrpTzName(caseSensitiveSort bool) {

  if tzDtoCol.tzDataDtos == nil {
    tzDtoCol.tzDataDtos = make([]TimeZoneDataDto, 0, 50)
  }

  if len(tzDtoCol.tzDataDtos) < 2 {
    return
  }

  var less func(i, j int) bool

  if !caseSensitiveSort {
    less = func(i, j int) bool {
      if strings.ToLower(tzDtoCol.tzDataDtos[i].MajorGroup) !=
        strings.ToLower(tzDtoCol.tzDataDtos[j].MajorGroup) {

        return strings.ToLower(tzDtoCol.tzDataDtos[i].MajorGroup) <
          strings.ToLower(tzDtoCol.tzDataDtos[j].MajorGroup)
      }

      return strings.ToLower(tzDtoCol.tzDataDtos[i].TzSortName) <
        strings.ToLower(tzDtoCol.tzDataDtos[j].TzSortName)

    }
  } else {
    less = func(i, j int) bool {
      if strings.ToLower(tzDtoCol.tzDataDtos[i].MajorGroup) !=
        strings.ToLower(tzDtoCol.tzDataDtos[j].MajorGroup) {

        return strings.ToLower(tzDtoCol.tzDataDtos[i].MajorGroup) <
          strings.ToLower(tzDtoCol.tzDataDtos[j].MajorGroup)
      }

      return strings.ToLower(tzDtoCol.tzDataDtos[i].TzSortName) <
        strings.ToLower(tzDtoCol.tzDataDtos[j].TzSortName)
    }
  }

  sort.Slice(tzDtoCol.tzDataDtos, less)
}


func main() {

  lenTimeZoneAry := len(timeZones)
  
  tzDataCol := TimeZoneDataCollection{}

  var tzMajorGroup string
  var tzName string
  var tzSortName string
  var tzValue string
  var tzClass int
  var err error


  for i:=0; i < lenTimeZoneAry; i++ {


    tzMajorGroup,
    tzName,
    tzSortName,
    tzValue,
    tzClass,
    err = parseTzValue(timeZones[i])

    if err != nil {
      fmt.Printf("%v\n", err.Error())
      return
    }

    err =  tzDataCol.AddNewTzDataDto(tzMajorGroup, tzName, tzSortName, tzValue, tzClass)

    if err != nil {

      fmt.Printf("Error returned by tzDataCol.AddNewTzDataDto(majorGroup, " +
        "tzName, tzValue, tzClass)\n" +
        "majorGroup='%v'\n" +
        "tzName='%v'\n" +
        "tzSortName='%v'\n" +
        "tzValue='%v'\n" +
        "tzClass='%v'\n" +
        "Error='%v'\n",
        tzMajorGroup,
        tzName,
        tzSortName, tzValue, tzClass, err.Error())

      return
    }

  }

  if tzDataCol.GetNumberOfDtos() != lenTimeZoneAry {
    fmt.Printf("ERROR: tzDataCol.GetNumberOfDtos() != lenTimeZoneAry\n" +
      "lenTimeZoneAry='%v'\n" +
      "tzDataCol.GetNumberOfDtos()='%v'\n",
      lenTimeZoneAry, tzDataCol.GetNumberOfDtos())
    return
  }

  fmt.Println( "================================")
  fmt.Println( "      Unsorted Time Zones       ")
  fmt.Println( "================================")
  fmt.Println()

  for j:=0; j < lenTimeZoneAry; j++ {

    tzDto, err := tzDataCol.PeekAtIndex(j)
    
    if err != nil {
      fmt.Printf("Error returned by tzDataCol.PeekAtIndex(j).\n" +
        "j='%v'\n" +
        "Error='%v'\n", j, err.Error())
      return
    }
    
    fmt.Printf("%3d.\tGroup: %v\t\t\tName: %v\t\tSortName: %v\t\tClass: %v\n",
      j+1, tzDto.MajorGroup, tzDto.TzName, tzDto.TzSortName, tzDto.TzClass )

  }

  fmt.Println()
  fmt.Println( "++++++++++++++++++++++++++++++++")
  fmt.Println( "      Sorted Time Zones         ")
  fmt.Println( "++++++++++++++++++++++++++++++++")
  fmt.Println()

  tzDataCol.SortTzDtosByMjrGrpTzName(false)

  for j:=0; j < lenTimeZoneAry; j++ {

    tzDto, err := tzDataCol.PeekAtIndex(j)

    if err != nil {
      fmt.Printf("Error returned by tzDataCol.PeekAtIndex(j).\n" +
        "j='%v'\n" +
        "Error='%v'\n", j, err.Error())
      return
    }

    fmt.Printf("%3d.\tGroup: %v\t\t\tName: %v\t\tSortName: %v\t\tClass: %v\n",
      j+1, tzDto.MajorGroup, tzDto.TzName, tzDto.TzSortName, tzDto.TzClass )

  }

}

func parseTzValue(
  rawTzValue string) (majorGroup,
                      tzName,
                      tzSortName,
                      tzValue string,
                      tzClass int,
                      err error) {

  majorGroup = ""
  tzName = ""
  tzSortName = ""
  tzValue = ""
  tzClass = 1
  err = nil
  ePrefix := "parseTzValue() "

  var sArray []string

  if strings.Contains(rawTzValue, "/") {

    sArray = strings.Split(rawTzValue, "/")

    if len(sArray) < 2 {
      err = fmt.Errorf(ePrefix + "ERROR: '/' Split returned a length of %v\n" +
        "rawTzValue='%v'\n", len(sArray), rawTzValue)
      return
    }

    majorGroup = sArray[0]
    tzName = sArray[1]
    tzSortName = tzName
    tzValue = rawTzValue
    tzClass = 1

  } else if strings.Contains(rawTzValue, "+") {
    sArray = strings.Split(rawTzValue, "+")

    if len(sArray) < 2 {
      err = fmt.Errorf(ePrefix + "ERROR: '+' Split returned a length of %v\n" +
        "rawTzValue='%v'\n", len(sArray), rawTzValue)
      return
    }

    majorGroup = sArray[0]
    tzName = rawTzValue
    tzSortName = rawTzValue
    tzValue = rawTzValue
    tzClass = 1

  } else if strings.Contains(rawTzValue, "-") {
    sArray = strings.Split(rawTzValue, "-")

    if len(sArray) < 2 {
      err = fmt.Errorf(ePrefix + "ERROR: '-' Split returned a length of %v\n" +
        "rawTzValue='%v'\n", len(sArray), rawTzValue)
      return
    }

    majorGroup = sArray[0]
    tzName = rawTzValue
    tzSortName = rawTzValue
    tzValue = rawTzValue
    tzClass = 1

  } else {
    majorGroup = rawTzValue
    tzName = rawTzValue
    tzSortName = rawTzValue
    tzValue = rawTzValue
    tzClass = 1
  }

  err = nil
  var err2 error

  tzSortName,
  err2 = parseTzSortName(tzName)

  if err2 != nil {
    err = fmt.Errorf(ePrefix + "ERROR returned by parseTzSortName(tzName)")
    tzSortName = ""
  }


  return majorGroup, tzName, tzSortName, tzValue, tzClass, err
}

func parseTzSortName(
  tzName string) (tzSortName string,
                  err error) {

  ePrefix := "parseTzSortName() "
  var sArry []string
  var lenSArry int
  var err2 error
  var num int

  if strings.Contains(tzName, "+") &&
    strings.ContainsAny(tzName, "0123456789") {

    sArry = strings.Split(tzName, "+")

    lenSArry = len(sArry)

    if lenSArry == 0 {
      tzSortName = ""
      err = fmt.Errorf(ePrefix +
        "'+' Split on TzName yielded zero length array.\n" +
        "tzName='%v'\n", tzName)
      return tzSortName, err

    } else if lenSArry == 1 {
      tzSortName = tzName
      err = nil
      return tzSortName, err

    } else if lenSArry == 2 {

      num, err2 =  strconv.Atoi(sArry[1])

      if err2 != nil {
        err = fmt.Errorf(ePrefix +
          "Error returned from '+' length array =2 strconv.Atoi(sArry[1])\n" +
          "sArry[1]=%v\n" +
          "Error='%v'\n", sArry[1], err2.Error())
        tzSortName = ""
        return tzSortName, err
      }

      tzSortName = sArry[0] +
          fmt.Sprintf("+%02d", num )
      err = nil
      return tzSortName, err
    }

    err = fmt.Errorf(ePrefix +
      "ERROR: '+' Split on TzName yielded %v-length array.\n" +
      "tzName='%v'\n" +
      "Split Array Length='%v'\n", lenSArry, tzName, lenSArry)

    tzSortName = ""
    return tzSortName, err
  }

  if strings.Contains(tzName, "-") &&
    strings.ContainsAny(tzName, "0123456789") {

    sArry = strings.Split(tzName, "-")

    lenSArry = len(sArry)

    if lenSArry == 0 {
      tzSortName = ""
      err = fmt.Errorf(ePrefix +
        "'-' Split on TzName yielded zero length array.\n" +
        "tzName='%v'\n", tzName)
      return tzSortName, err

    } else if lenSArry == 1 {
      tzSortName = tzName
      err = nil
      return tzSortName, err

    } else if lenSArry == 2 {

      num, err2 = strconv.Atoi(sArry[1])

      if err2 != nil {
        err = fmt.Errorf(ePrefix +
          "Error returned from '-' length array =2 strconv.Atoi(sArry[1])\n" +
          "sArry[1]=%v\n" +
          "Error='%v'\n", sArry[1], err2.Error())
        tzSortName = ""
        return tzSortName, err
      }

      tzSortName = sArry[0] +
          fmt.Sprintf("-%02d", num )
      err = nil
      return tzSortName, err

    }

    err = fmt.Errorf(ePrefix +
      "ERROR: '-' Split on TzName yielded %v-length array.\n" +
      "tzName='%v'\n" +
      "Split Array Length='%v'", lenSArry, tzName, lenSArry)
    tzSortName = ""
    return tzSortName, err

  }

  tzSortName = tzName

  err = nil
  return tzSortName, err
}



