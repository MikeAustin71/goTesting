package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/mikeaustin71/DateTimeTests/01_RevisedJulianCalendar/RJC_Libs02"
)

func main() {

	ePrefix := "RJC_007/main() "

	err := rjMOrdinalDayTest(ePrefix)

	if err != nil {
		fmt.Printf("%v", err.Error())
		separator := strings.Repeat("%", 65)
		fmt.Println(separator)
		fmt.Println("EXECUTION FAILED!!!!!")
		fmt.Println(separator)
		return
	}

	separator := strings.Repeat("!", 75)
	fmt.Println()
	fmt.Println(separator)
	// fmt.Printf("Expected Ordinal Day Number:        730120.5\n")
	fmt.Println("     SUCCESSFUL COMPLETION!!!")
	fmt.Println(separator)

}

func completeYearsIntervalTest(ePrefix string) (err error) {

	ePrefix += "completeYearsIntervalTest() "

	err = nil

	targetYear := int64(1)
	targetMonth := 1
	targetDay := 1
	targetHour := 0
	targetMinute := 0
	targetSecond := 0
	targetNanosecond := 0

	var baseDateTimeDto, targetDateTimeDto RJC_Libs02.DateTimeTransferDto

	separator := strings.Repeat("=", 65)

	fmt.Println(separator)
	fmt.Printf("%v\n",
		ePrefix)
	calRJM2 := RJC_Libs02.CalendarRevisedJulianUtility{}

	isTargetLeapYear := calRJM2.IsLeapYear(targetYear)

	targetDateTimeDto, err = RJC_Libs02.DateTimeTransferDto{}.New(
		isTargetLeapYear,
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

	targetDateTimeDto.SetTag("targetDateTimeDto")

	// Revised Julian Calendar ------
	// JDN 0 = -4713/02/08 12:00 (12-Noon) Revised Julian Calendar

	baseYear := int64(-4713)
	baseMonth := 2
	baseDay := 8
	baseHour := 12
	baseMinute := 0
	baseSecond := 0
	baseNanosecond := 0

	isBaseLeapYear := calRJM2.IsLeapYear(targetYear)

	// Base Date Time:  February 8, -4713 12:00:00 UTC
	baseDateTimeDto, err = RJC_Libs02.DateTimeTransferDto{}.New(
		isBaseLeapYear,
		baseYear,
		baseMonth,
		baseDay,
		baseHour,
		baseMinute,
		baseSecond,
		baseNanosecond,
		ePrefix)

	if err != nil {
		return err
	}

	baseDateTimeDto.SetTag("baseDateTimeDto")

	fmt.Printf("%v", baseDateTimeDto.String())

	fmt.Printf("%v", targetDateTimeDto.String())

	calUtil := RJC_Libs02.CalendarUtility{}
	var WholeYearsInterval int64

	WholeYearsInterval, err =
		calUtil.GetCompleteYearInterval(
			baseDateTimeDto,
			targetDateTimeDto,
			ePrefix)

	if err != nil {
		return err
	}

	calRJM := RJC_Libs02.CalendarRevisedJulianUtility{}

	var totalDays int64

	totalDays, err = calRJM.VerifyDayCountForWholeYearInterval(
		WholeYearsInterval,
		ePrefix)

	fmt.Println(separator)
	fmt.Printf("            Whole Years Interval: %v\n",
		WholeYearsInterval)

	fmt.Printf("  Verify Total Days For Interval: %v\n",
		totalDays)

	fmt.Println("Expected Total Days For Interval: 1721387")
	fmt.Println(separator)
	fmt.Println("     SUCCESSFUL COMPLETION!!!")
	fmt.Println(separator)

	return nil
}

func rjMJulianDayTest(ePrefix string) (err error) {

	// Julian Day Count 1721425.5
	// Revised Julian Calendar ------
	// JDN 0 = -4713/02/08 12:00 (12-Noon) Revised Julian Calendar
	// From:  2/8/-4713 12:00:00 (12-Noon)
	//   To:  1/1/0001  00:00:00 (midnight)
	//
	// Gregorian Calendar -----------
	// JDN 0 = -4712/01/01 12:00:00 (Noon) Gregorian Calendar

	ePrefix += "rjMJulianDayTest() "
	err = nil

	targetYear := int64(1)
	targetMonth := 1
	targetDay := 1
	targetHour := 0
	targetMinute := 0
	targetSecond := 0
	targetNanosecond := 0

	// expectedInterval := 4713

	// Using Revised Julian Calendar Base date/time:
	//  -4713-11-23 12:00:00 UTC (Noon)
	// expectedTotalDayNoTime := 12451544.5 // For 2000/1/1
	// expectedTotalDayNoTime := 1721425.5 // For 0001/1/1
	expectedTotalDayNoTime :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(1721425.5)

	// 	testRJGYear1()
	calRJM := RJC_Libs02.CalendarRevisedJulianUtility{}

	var julianDayNumber int64
	var julianDayNumberTime, julianDayNumberTimeFraction *big.Float

	julianDayNumber,
		julianDayNumberTime,
		julianDayNumberTimeFraction,
		err = calRJM.GetJulianDayNumber(
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

	separator := strings.Repeat("&", 75)
	fmt.Println(separator)
	fmt.Printf("Expected Julian Day Number: %45.30f\n",
		expectedTotalDayNoTime)
	fmt.Println(separator)
	fmt.Printf("         Julian Day Number: %14d\n",
		julianDayNumber)

	fmt.Printf("     Julian DayNumber Time: %45.30f\n",
		julianDayNumberTime)

	fmt.Printf("Julian DayNumber Fraction: %46.30f\n",
		julianDayNumberTimeFraction)
	fmt.Println(separator)

	return err
}

func rjMOrdinalDayTest(ePrefix string) (err error) {

	ePrefix += "rjMOrdinalDayTest() "
	err = nil

	// 730120.5 = Ordinal Day For:
	// January 1, 2000 12:00:00.000000000 UTC
	//
	expectedOrdinalDayNoTime :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(730121.5)

	targetYear := int64(1)
	targetMonth := 1
	targetDay := 2
	targetHour := 12
	targetMinute := 0
	targetSecond := 0
	targetNanosecond := 0

	calRM := RJC_Libs02.CalendarRevisedJulianUtility{}

	var ordinalDayNumber int64
	var ordinalDayNumberTime,
		ordinalDayNumberTimeFraction *big.Float

	ordinalDayNumber,
		ordinalDayNumberTime,
		ordinalDayNumberTimeFraction,
		err = calRM.GetOrdinalDayNumberTime(
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

	separator := strings.Repeat("&", 75)
	fmt.Println(separator)
	fmt.Printf("             Ordinal Day Number: %14d\n",
		ordinalDayNumber)
	fmt.Println(separator)
	fmt.Printf("         Ordinal DayNumber Time: %45.30f\n",
		ordinalDayNumberTime)
	fmt.Printf("Expected Ordinal DayNumber Time: %45.30f\n",
		expectedOrdinalDayNoTime)
	fmt.Println(separator)
	fmt.Printf("     Ordinal DayNumber Fraction: %46.30f\n",
		ordinalDayNumberTimeFraction)
	fmt.Println(separator)

	return err
}
