package main

import (
	"bitbucket.org/AmarilloMike/GolangMikeSamples/DateTimeTests/01_RevisedJulianCalendar/RJC_Libs"
	"fmt"
	"math/big"
	"strings"
)





func main(){

	// Julian Day Count 1721425.5
	// Revised Julian Calendar ------
	// JDN 0 = -4713/02/08 12:00 (12-Noon) Revised Julian Calendar
	// From:  2/8/-4713 12:00:00 (12-Noon)
	//   To:  1/1/0001  00:00:00 (midnight)
	//
	// Gregorian Calendar -----------
	// JDN 0 = -4712/01/01 12:00:00 (Noon) Gregorian Calendar

	ePrefix := "RJC_006/main() "
	var err error


	targetYear := int64(1)
	targetMonth := 1
	targetDay := 1
	targetHour := 0
	targetMinute:=0
	targetSecond:=0
	targetNanosecond:=0


	baseYear := int64(-4713)
	baseMonth := 2
	baseDay := 8
	baseHour := 0
	baseMinute:=0
	baseSecond:=0
	baseNanosecond:=0

	expectedInterval := 4713

	// expectedTotalDays := 730120.5 for 2000/1/1/
	expectedTotalDays := 1721425 // For -4713/2/9

// 	testRJGYear1()

	err = testRJGJulianDayNo(
		baseYear,
		baseMonth,
		baseDay,
		baseHour,
		baseMinute,
		baseSecond,
		baseNanosecond,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		fmt.Printf("ERROR *****\n" +
			"%v\n", err.Error())
		return
	}

	separator := strings.Repeat("&", 65)
	fmt.Println(separator)
	fmt.Printf("  Expected Interval: %v\n", expectedInterval)
	fmt.Printf("Expected Total Days: %v\n", expectedTotalDays)
	fmt.Println(separator)
	fmt.Println("     SUCCESSFUL COMPLETION!!!")

}

func testRJGYear1 () {

	// Julian Day Count 1721425.5
	// Revised Julian Calendar ------
	// JDN 0 = -4713/02/08 12:00 (12-Noon) Revised Julian Calendar
	// From:  2/8/-4713 12:00:00 (12-Noon)
	//   To:  1/1/0001  00:00:00 (midnight)
	//
	// Gregorian Calendar -----------
	// JDN 0 = -4712/01/01 12:00:00 (Noon) Gregorian Calendar

	targetYear := int64(1)
	targetMonth := 1
	targetDay := 1
	targetHour := 0
	targetMinute:=0
	targetSecond:=0
	targetNanosecond:=0

	ePrefix := "RJC_006/main() "

	calMech := RJC_Libs.CalendarMechanics{}

	var err error

	baseYear := int64(-4713)
	baseMonth := 2
	baseDay := 8
	baseHour := 12
	baseMinute:=0
	baseSecond:=0
	baseNanosecond:=0

	expectedInterval := 4713

	expectedTotalDays := 1721425

	var lastCompleteTargetYear int64

	lastCompleteTargetYear, err = calMech.GetLastCompleteYear(
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}


	separator := strings.Repeat("=", 65)
	fmt.Println(separator)
	fmt.Println("              " + ePrefix)
	fmt.Println(separator)
	fmt.Printf("                 baseYear: %v\n", baseYear)
	fmt.Printf("                baseMonth: %v\n", baseMonth)
	fmt.Printf("                  baseDay: %v\n", baseDay)
	fmt.Printf("                 baseHour: %v\n", baseHour)
	fmt.Printf("               baseMinute: %v\n", baseMinute)
	fmt.Printf("               baseSecond: %v\n", baseSecond)
	fmt.Printf("           baseNanosecond: %v\n", baseNanosecond)
	fmt.Println(separator)
	fmt.Printf("               targetYear: %v\n", targetYear)
	fmt.Printf("              targetMonth: %v\n", targetMonth)
	fmt.Printf("                targetDay: %v\n", targetDay)
	fmt.Printf("               targetHour: %v\n", targetHour)
	fmt.Printf("             targetMinute: %v\n", targetMinute)
	fmt.Printf("             targetSecond: %v\n", targetSecond)
	fmt.Printf("         targetNanosecond: %v\n", targetNanosecond)
	fmt.Println(separator)
	fmt.Printf("Last Complete Target Year: %v\n", lastCompleteTargetYear)
	fmt.Printf("Expected Interval (Years): %v\n", expectedInterval)
	fmt.Println(separator)

	var completeYearsInterval int64

	completeYearsInterval, err =
		calMech.GetCompleteYearInterval(
		baseYear,
		baseMonth,
		baseDay,
		baseHour,
		baseMinute,
		baseSecond,
		baseNanosecond,
		targetYear,
		targetMonth,
		targetDay,
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}


	calRJM := RJC_Libs.CalendarRevisedJulianMechanics{}

	wholeYearDaysDays := calRJM.NumCalendarDaysForWholeYearsInterval(completeYearsInterval)

	isLeapYear := calRJM.IsLeapYear(targetYear)

	var targetYearElapsedWholeOrdinalDays, baseYearElapsedWholeOrdinalDays int

	targetYearElapsedWholeOrdinalDays, err = calMech.GetElapsedWholeOrdinalDaysInYear(
		isLeapYear,
		targetMonth,
		targetDay,
		"main() ")

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	isLeapYear = calRJM.IsLeapYear(baseYear)

	baseYearElapsedWholeOrdinalDays, err = calMech.GetElapsedWholeOrdinalDaysInYear(
		isLeapYear,
		baseMonth,
		baseDay,
		"main() ")

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	grandTotalDays := wholeYearDaysDays+ int64(targetYearElapsedWholeOrdinalDays) + int64(baseYearElapsedWholeOrdinalDays)

	fmt.Println(separator)
	fmt.Printf("    Whole Year Days: %v\n", wholeYearDaysDays)
	fmt.Printf("  targetYearElapsedWholeOrdinalDays: %v\n", targetYearElapsedWholeOrdinalDays)
	fmt.Printf("    baseYearElapsedWholeOrdinalDays: %v\n", baseYearElapsedWholeOrdinalDays)
	fmt.Printf("         Total Days: %v\n",
		grandTotalDays)
	fmt.Printf("Expected Total Days: %v\n",
		expectedTotalDays)
	fmt.Println(separator)

	var targetYearJulianFraction, baseYearJulianFraction *big.Float

	baseYearJulianFraction, _, err = calMech.GetJulianDayNoFraction(
		baseHour,
		baseMinute,
		baseSecond,
		baseNanosecond,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmt.Printf("  Base Year JDN Fraction: %31.30f\n",
		baseYearJulianFraction)

	targetYearJulianFraction, _, err = calMech.GetJulianDayNoFraction(
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}

	fmt.Printf("Target Year JDN Fraction: %31.30f\n",
		targetYearJulianFraction)

	fmt.Println(separator)

}

func testRJGJulianDayNo(
	baseYear int64,
	baseMonth,
	baseDay,
	baseHour,
	baseMinute,
	baseSecond,
	baseNanosecond int,
	targetYear int64,
	targetMonth,
	targetDay,
	targetHour,
	targetMinute,
	targetSecond,
	targetNanosecond int,
	ePrefix string) error {



	// Julian Day Count 1721425.5
	// Revised Julian Calendar ------
	// JDN 0 = -4713/02/08 12:00 (12-Noon) Revised Julian Calendar
	// From:  2/8/-4713 12:00:00 (12-Noon)
	//   To:  1/1/0001  00:00:00 (midnight)
	//
	// Gregorian Calendar -----------
	// JDN 0 = -4712/01/01 12:00:00 (Noon) Gregorian Calendar


	ePrefix += "testRJGJulianDayNo() "

	calMech := RJC_Libs.CalendarMechanics{}

	var err error
	//
	//var lastCompleteTargetYear int64
	//
	//lastCompleteTargetYear, err = calMech.GetLastCompleteYear(
	//	targetYear,
	//	targetMonth,
	//	targetDay,
	//	targetHour,
	//	targetMinute,
	//	targetSecond,
	//	targetNanosecond,
	//	ePrefix)
	//
	//if err != nil {
	//	return err
	//}


	separator := strings.Repeat("=", 65)
	fmt.Println(separator)
	fmt.Println("              " + ePrefix)
	fmt.Println(separator)
	fmt.Printf("                 baseYear: %v\n", baseYear)
	fmt.Printf("                baseMonth: %v\n", baseMonth)
	fmt.Printf("                  baseDay: %v\n", baseDay)
	fmt.Printf("                 baseHour: %v\n", baseHour)
	fmt.Printf("               baseMinute: %v\n", baseMinute)
	fmt.Printf("               baseSecond: %v\n", baseSecond)
	fmt.Printf("           baseNanosecond: %v\n", baseNanosecond)
	fmt.Println(separator)
	fmt.Printf("               targetYear: %v\n", targetYear)
	fmt.Printf("              targetMonth: %v\n", targetMonth)
	fmt.Printf("                targetDay: %v\n", targetDay)
	fmt.Printf("               targetHour: %v\n", targetHour)
	fmt.Printf("             targetMinute: %v\n", targetMinute)
	fmt.Printf("             targetSecond: %v\n", targetSecond)
	fmt.Printf("         targetNanosecond: %v\n", targetNanosecond)
	fmt.Println(separator)
//	fmt.Printf("Last Complete Target Year: %v\n", lastCompleteTargetYear)
	fmt.Println(separator)

	var completeYearsInterval int64

	completeYearsInterval, err =
		calMech.GetCompleteYearInterval(
			baseYear,
			baseMonth,
			baseDay,
			baseHour,
			baseMinute,
			baseSecond,
			baseNanosecond,
			targetYear,
			targetMonth,
			targetDay,
			targetHour,
			targetMinute,
			targetSecond,
			targetNanosecond,
			ePrefix)

	if err != nil {
		return err
	}

	calRJM := RJC_Libs.CalendarRevisedJulianMechanics{}

	wholeYearDaysDays := calRJM.NumCalendarDaysForWholeYearsInterval(completeYearsInterval)

	isLeapYear := calRJM.IsLeapYear(targetYear)

	var targetYearOrdinalNumber, baseYearElapsedWholeOrdinalDays int

	targetYearOrdinalNumber, err = calMech.GetElapsedWholeOrdinalDaysInYear(
		isLeapYear,
		targetMonth,
		targetDay,
		ePrefix)

	if err != nil {
		return err
	}

	isLeapYear = calRJM.IsLeapYear(baseYear)

	baseYearElapsedWholeOrdinalDays, err = calMech.GetElapsedWholeOrdinalDaysInYear(
		isLeapYear,
		baseMonth,
		baseDay,
		ePrefix)

	if err != nil {
		return err
	}

	grandTotalDays := wholeYearDaysDays+ int64(targetYearOrdinalNumber) + int64(baseYearElapsedWholeOrdinalDays)

	fmt.Println(separator)
	fmt.Printf("    Whole Year Days: %v\n", wholeYearDaysDays)
	fmt.Printf("  targetYearOrdinalNumber: %v\n", targetYearOrdinalNumber)
	fmt.Printf("    baseYearElapsedWholeOrdinalDays: %v\n", baseYearElapsedWholeOrdinalDays)
	fmt.Printf("         Total Days: %v\n",
		grandTotalDays)
	fmt.Println(separator)

	var targetYearJulianFraction, baseYearJulianFraction *big.Float

	baseYearJulianFraction, _, err = calMech.GetJulianDayNoFraction(
		baseHour,
		baseMinute,
		baseSecond,
		baseNanosecond,
		ePrefix)

	if err != nil {
		return err
	}

	fmt.Printf("  Base Year JDN Fraction: %31.30f\n",
		baseYearJulianFraction)

	targetYearJulianFraction, _, err = calMech.GetJulianDayNoFraction(
		targetHour,
		targetMinute,
		targetSecond,
		targetNanosecond,
		ePrefix)

	if err != nil {
		return err
	}

	fmt.Printf("Target Year JDN Fraction: %31.30f\n",
		targetYearJulianFraction)

	fmt.Println(separator)

	return nil
}