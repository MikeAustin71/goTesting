package main

import "fmt"

func main() {

	rjcToYear001()
	fmt.Println()
	fmt.Println()
	rjcToYear002()

}

func rjcToYear001() {

	ePrefix := "rjcToYear001()"

	// November 24, 4714 BCE, in the proleptic Gregorian calendar date
	// for Julian Day Number 0 (zero). On the Julian Calendar, Day Zero
	// is January 1, 4713 BCE (-4712).
	//
	// For the revised Julian Calendar, Midnight, February 7, 4714 BCE
	// to Midnight January 1, 0001 yields a Julian Day Number of 1721425.5.
	// Day Zero therefore computes as -4713/2/18

	originalStartYear := int64(-4713)
	originalStartMonth := 2
	originalStartDay := 7
// Assumed to start at midnight 2/7/-4713 (Astronomical Year Numbering)

	beginningOrdinalDays,
	err := computeOrdinalDateFromMthDayYr(
		originalStartYear,
		originalStartMonth,
		originalStartDay,
		ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

startYear := originalStartYear + 1 // -4712

	// Target Year = 1/1/0001: January 1, 1 CE
	targetYear := int64(1)

	totalDays := beginningOrdinalDays
	yearDays := int64(0)

	for i:=startYear; i < targetYear; i++ {

	yearDays = 365

	if rjcIsLeapYear(i) {
		yearDays = 366
	}

		totalDays += yearDays
	}

// 1721425.5
expectedJulianDayNo := int64(1721425)

varianceDays := expectedJulianDayNo - totalDays

	fmt.Println(ePrefix)
	fmt.Println("--------------------------------------")
	fmt.Println("Computing Julian Day Number for")
	fmt.Println("Revised Julian Calendar 1/1/1 CE")
	fmt.Println("--------------------------------------")
	fmt.Println("    THIS CALCULATION IS CORRECT!")
	fmt.Println("--------------------------------------")
	fmt.Printf("Expected Julian Day No: %v\n", expectedJulianDayNo)
	fmt.Printf("Computed Julian Day No: %v\n", totalDays)
	fmt.Println("--------------------------------------")
	fmt.Printf("         Variance Days: %v\n", varianceDays)
	fmt.Println("--------------------------------------")
	fmt.Printf("     Original Start Year: %v\n", originalStartYear)
	fmt.Printf("    Original Start Month: %v\n", originalStartMonth)
	fmt.Printf("      Original Start Day: %v\n", originalStartDay)
	fmt.Printf("            Ordinal Days: %v\n", beginningOrdinalDays)
	fmt.Printf("     First Complete Year: %v\n", startYear)
	fmt.Println("--------------------------------------")
	fmt.Println("--------------------------------------")


}

// Receives Julian Day Number and start date. Then
// counts backwards to compute start date.
func rjcToYear002() {

	ePrefix := "rjcToYear002() "
	// 1721425.5
	julianDayNo := int64(1721425)

	// Start Date January 1, 1, CE Midnight
	// Start Year = 1 - 1 = 0
	year := int64(1)
	isRemainderGreaterThanOrEqualToYear := true

	var yearDays int64
	yearCount:= int64(0)
	dayCount:= int64(0)
	leapDayCount := int64(0)
	daysRemaining := julianDayNo
	firstCompleteYear := year - 1

	for isRemainderGreaterThanOrEqualToYear == true {

		year--
		yearCount++

		yearDays = 365

		if rjcIsLeapYear(year) {
			leapDayCount++
			yearDays = 366
		}

		dayCount+= yearDays

		daysRemaining-= yearDays

		if daysRemaining < 365 {
			isRemainderGreaterThanOrEqualToYear = false
		}
	}


	var nextYear int64

	if year < 0 {
		nextYear = year - 1
	} else {
		nextYear = year + 1
	}

	finalYear,
	month,
	day,
	err := computeMonthDayFromOrdDate(
		daysRemaining,
		nextYear,
		ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	fmt.Println()
	fmt.Println(ePrefix)
	fmt.Println("--------------------------------------")
	fmt.Println("Computing Revised Julian Calendar Date")
	fmt.Printf("for Julian Day Number: %v\n", julianDayNo)
	fmt.Println("--------------------------------------")
	fmt.Println("    THIS CALCULATION IS CORRECT!")
	fmt.Println("--------------------------------------")
	fmt.Printf("     Starting Year: %v\n", "0001 CE")
	fmt.Printf("    Starting Month: %v\n", "January")
	fmt.Printf("      Starting Day: %v\n", "1")
	fmt.Printf("   First Full Year: %v\n", firstCompleteYear)
	fmt.Println("--------------------------------------")
	fmt.Printf("     Elapsed Years: %v\n", yearCount)
	fmt.Printf("    Leap Day Count: %v\n", leapDayCount)
	fmt.Printf("      Elapsed Days: %v\n", dayCount)
	fmt.Println("--------------------------------------")
	fmt.Printf("Last Complete Year: %v\n", year)
	fmt.Printf(" Computed End Year: %v\n", finalYear)
	fmt.Printf("Computed End Month: %v\n", month)
	fmt.Printf("  Computed End Day: %v\n", day)
	fmt.Println("--------------------------------------")
	fmt.Printf("      Initial Year: %v\n", nextYear)
	fmt.Printf("      Ordinal Date: %v\n", daysRemaining)
	fmt.Println("--------------------------------------")

	return
}

func computeMonthDayFromOrdDate(
	ordinalDate int64,
	inputYear int64,
	ePrefix string)(year int64, month, day int, err error) {

	ePrefix += "computeMonthDayFromOrdDate() "
	year = -1
	month = -1
	day = -1
	err = nil

	if ordinalDate < 0 ||
		ordinalDate > 364 {
		err = fmt.Errorf("\n" +ePrefix + "Error:\n" +
			"Input Parameter 'ordinalDate' is INVALID!\n" +
			"ordinalDate='%v'\n",
			ordinalDate)
		return year, month, day, err
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

	year = inputYear

	isLeapYear := rjcIsLeapYear(year)

	for i:=11; i > -1; i-- {

		if ordDays[i] <= ordinalDate {

			ordinalDate -= ordDays[i]

			month = i + 1

			if month > 2 && isLeapYear {
				ordinalDate--
			}

			if ordinalDate == 0 &&
					month==1 {
				year--
				month = 12
				day = 31

				break

			} else if ordinalDate == 0 {
				month--
				day = mthDays[month]
				break
			} else {
				day = int(ordinalDate)
				break
			}

		}
	}

	return year, month, day, err
}

func computeOrdinalDateFromMthDayYr(
	year int64,
	month int,
	day int,
	ePrefix string) (ordinalDate int64, err error) {

	ePrefix += "computeOrdinalDateFromMthDayYr() "
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

	isLeapYear := rjcIsLeapYear(year)

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

func rjcIsLeapYear(year int64) bool {

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