package main

import (
	"fmt"
	"strings"
)

func main() {

	// Julian Day Count 1721425.5

	// 2/8/-4713

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
	fmt.Println("              RJC_004/main() ")
	fmt.Println(separator)
	fmt.Printf("               targetYear: %v\n", targetYear)
	fmt.Printf("Last Complete Target Year: %v\n", lastCompleteTargetYear)
	fmt.Printf("              targetMonth: %v\n", targetMonth)
	fmt.Printf("                targetDay: %v\n", targetDay)
	fmt.Printf(" Last Complete Target Day: %v\n", lastCompleteTargetDay)
	fmt.Printf("Expected Interval (Years): %v\n", expectedInterval)
	fmt.Println(separator)

	completeYearsInterval := computeCompleteYearInterval(baseYear, lastCompleteTargetYear)


	//_, cycleDays, remainderYears := computeRevisedJulian900YearCycles(completeYearsInterval, "main() ")

	wholeYearDaysDays := computeRevisedJulianCalendarDaysForWholeYears(completeYearsInterval)

	isLeapYear := revisedJulianCalendarIsLeapYear(targetYear)

	ordinalDays, err := computeOrdinalDayNoFromYrMthDay(
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
		wholeYearDaysDays+ ordinalDays)
	fmt.Printf("Expected Total Days: %v\n",
		expectedTotalDays)
	fmt.Println(separator)
	fmt.Println(separator)

}


// computeCompleteYearInterval - Calculates the number of
// complete years between a target year and a base year.
// The returned result is always a positive value expressed
// as a type 'int64'
//
func computeCompleteYearInterval(
	baseYear,
	targetYear int64) (completedYearsInterval int64) {

	completedYearsInterval = baseYear - targetYear

	if completedYearsInterval < 0 {
		completedYearsInterval *= -1
	}

	separator := strings.Repeat("-", 65)

	fmt.Println()
	fmt.Println("computeCompleteYearInterval")
	fmt.Println(separator)
	fmt.Printf("                 baseYear: %v\n",baseYear)
	fmt.Printf("Last Complete Target Year: %v\n", targetYear)
	fmt.Println(separator)
	fmt.Printf("Complete Years Interval: %v\n", completedYearsInterval)
	fmt.Println(separator)

	return completedYearsInterval
}

func computeOrdinalDayNoFromYrMthDay(
	isLeapYear bool,
	month int,
	day int,
	ePrefix string) (ordinalDate int64, err error) {

	ePrefix += "computeOrdinalDateFromYrMthDay() "
	ordinalDate = -1
	err = nil

	if month < 1 || month > 12 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input Parameter 'month' is INVALID!\n" +
			"month='%v'\n", month)
		return ordinalDate, err
	}


	ordDays := []int64 {
		0,
		31,
		59,
		90,
		120,
		151,
		181,
		212,
		243,
		273,
		304,
		334,
	}

	mthDays := []int {
		-1,
		31,
		28,
		31,
		30,
		31,
		30,
		31,
		31,
		30,
		31,
		30,
		31,
	}

	monthDays := mthDays[month]

	if day > monthDays || day < 1 {
		err = fmt.Errorf("\n" + ePrefix + "Error:\n" +
			"Input parameter 'day' is INVALID!\n" +
			"month='%v'\n day='%v'\n",
			month, day)
		return ordinalDate, err
	}

	if month == 1 {
		ordinalDate = int64(day)
		return ordinalDate, err
	} else {

		ordinalDate = ordDays[month-1] + int64(day)

		if isLeapYear && month > 2 {
			ordinalDate++
		}
	}

	return ordinalDate, err
}

// computeRevisedJulianCalendarDaysForWholeYears - Computes
// the total number of 24-hour days in a period of years
// specified by input parameter 'completeYearsInterval'.
//
// 'completeYearsInterval' is defined as a contiguous series
// of whole or complete years. No partial years should be
// included in this interval.
//
func computeRevisedJulianCalendarDaysForWholeYears(
	completeYearsInterval int64) (totalDays int64) {

	separator := strings.Repeat("*", 65)

	totalDays = 0

	if completeYearsInterval < 0 {
		completeYearsInterval *= -1
	}

	if completeYearsInterval == 0 {
		return 0
	}

	fmt.Println()
	fmt.Println("computeRevisedJulianCalendarDaysForWholeYears() ")
	fmt.Println(separator)
	fmt.Printf("    Complete Years Interval: %v\n", completeYearsInterval)

	if completeYearsInterval >= 900 {

		numOfCycles := completeYearsInterval / 900

		totalDays = numOfCycles * 328718

		fmt.Printf("  Number of 900-Year Cycles: %v\n", numOfCycles)
		fmt.Printf("Number of Days in %v-Cycles: %v\n", numOfCycles, totalDays)

		completeYearsInterval = completeYearsInterval - (numOfCycles * 900)

		fmt.Printf("  Number of Remainder Years: %v\n", completeYearsInterval)
		fmt.Println(separator)
		fmt.Println()

	}

	totalDays += completeYearsInterval * 365

	leapDays := completeYearsInterval / 4

	skipLeapDays := completeYearsInterval / 100

	addLeapDays := int64(0)

	if completeYearsInterval >= 200 {
		addLeapDays++
	}

	if completeYearsInterval >= 600 {
		addLeapDays++
	}

	totalDays += leapDays + addLeapDays - skipLeapDays

	fmt.Println(separator)
	fmt.Printf("Totla Days In completeYearsInterval: %v\n",
		totalDays)
	fmt.Println(separator)
	fmt.Println()

	return totalDays
}


func revisedJulianCalendarIsLeapYear(year int64) bool {

	var isLeapYear bool

	if year < 0 {
		year *= -1
	}

	if year % int64(4) == 0 {

		isLeapYear = true

		if year % 100 == 0 {

			isLeapYear = false

			mod900 := year % int64(900)

			if  mod900 == 200 ||
				mod900 == 600 {
				isLeapYear = true
			}

		}

	} else {
		isLeapYear = false
	}

	return isLeapYear
}
