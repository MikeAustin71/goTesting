package main

import (
	"fmt"
	"strconv"
)

var UtcArray = []string {
	"UTC+9",
	"UTC+10",
	"UTC-9",
	"UTC-10",

}

func main() {

	for i:= 0; i < len(UtcArray); i++ {

		utcEquivalent := UtcArray[i]

		utcPrefix := utcEquivalent[0:4]
		var utcHours int
		var err error

		utcHoursStr := utcEquivalent[4:]

		utcHours, err = strconv.Atoi(utcHoursStr)

		if err != nil {
			fmt.Printf("Error returned by strconv.Atoi(utcEquivalent[4:])\n" +
				"utcEquivalent[4:]='%v'\n" +
				"Error='%v'\n", utcEquivalent[4:], err.Error())
		}

		utcOffset := fmt.Sprintf(utcPrefix + "%02d00", utcHours )

		fmt.Printf("Utc Prefix: %v Utc Hours Str: %v  UtcOffset: %v \n",
			utcPrefix, utcHoursStr, utcOffset)

	}



}
