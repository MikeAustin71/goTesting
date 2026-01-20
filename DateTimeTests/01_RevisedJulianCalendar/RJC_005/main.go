package main

import (
	"fmt"
	"strings"

	"github.com/mikeaustin71/DateTimeTests/01_RevisedJulianCalendar/RJC_005/libs"
)

// "D:/gowork/src/bitbucket.org\\AmarilloMike\\GolangMikeSamples"
func main() {

	// Julian Day Count 1721425.5
	// From:  2/8/-4713 12:00:00 (12-Noon)
	//   To:  1/1/0001  00:00:00 (midnight)

	baseYear := int64(1)
	//baseMonth := 1
	//baseDay := 1
	//baseHour := 0
	//baseMinute:=0
	//baseSecond:=0
	//baseNanosecond:=0

	targetYear := int64(-4713)
	targetMonth := 2
	targetDay := 8
	//targetHour := 12
	//targetMinute:=0
	//targetSecond:=0
	//targetNanosecond:=0

	expectedInterval := 4713

	expectedTotalDays := 1721425

	lastCompleteTargetYear := targetYear + 1

	lastCompleteTargetDay := targetDay - 1

	separator := strings.Repeat("=", 65)
	fmt.Println(separator)
	fmt.Println("              RJC_005/main() ")
	fmt.Println(separator)
	fmt.Printf("               targetYear: %v\n", targetYear)
	fmt.Printf("Last Complete Target Year: %v\n", lastCompleteTargetYear)
	fmt.Printf("              targetMonth: %v\n", targetMonth)
	fmt.Printf("                targetDay: %v\n", targetDay)
	fmt.Printf(" Last Complete Target Day: %v\n", lastCompleteTargetDay)
	fmt.Printf("Expected Interval (Years): %v\n", expectedInterval)
	fmt.Println(separator)

	calMech := libs.CalendarMechanics{}

	completeYearsInterval := calMech.ComputeCompleteYearInterval(
		baseYear, lastCompleteTargetYear)

	calRJM := libs.CalendarRevisedJulianMechanics{}

	wholeYearDaysDays := calRJM.NumCalendarDaysForWholeYearsInterval(completeYearsInterval)

	isLeapYear := calRJM.IsLeapYear(targetYear)

	ordinalDays, err := calMech.GetOrdinalDayNumber(
		isLeapYear,
		targetMonth,
		lastCompleteTargetDay,
		"main() ")

	if err != nil {
		fmt.Println("Error")
		return
	}

	fmt.Println(separator)
	fmt.Printf("    Whole Year Days: %v\n", wholeYearDaysDays)
	fmt.Printf("       Ordinal Days: %v\n", ordinalDays)
	fmt.Printf("         Total Days: %v\n",
		wholeYearDaysDays+ordinalDays)
	fmt.Printf("Expected Total Days: %v\n",
		expectedTotalDays)
	fmt.Println(separator)
	fmt.Println(separator)

}
