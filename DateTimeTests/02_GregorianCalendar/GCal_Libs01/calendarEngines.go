package GCal_Libs01

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"
)

type CalendarEngines struct {
	lock  *sync.Mutex
}

// JulianDayNoEngine - Method stub for a function to calculate the Julian
// Day Number from a specific date for all supported calendars.
func (calEng *CalendarEngines) JulianDayNoEngine(
	targetDateTimeDto DateTimeTransferDto,
	calCycleCfg CalendarCycleConfiguration,
	ePrefix string) (
	julianDayNumber int64,
	julianDayNumberTime *big.Float,
	julianDayNumberTimeFraction *big.Float,
	err error) {

	if calEng.lock == nil {
		calEng.lock = &sync.Mutex{}
	}

	calEng.lock.Lock()

	defer calEng.lock.Unlock()

	ePrefix += "CalendarEngines.JulianDayNoEngine() "

	julianDayNumber = 0

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	julianDayNumberTimeFraction =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetFloat64(0.0)

	err = nil


	if calCycleCfg.GetNumberOfCalendarCycleConfigDtos() == 0 {
		err = errors.New(ePrefix + "\n" +
			"Error: Input parameter 'calCyclesCfg' is invalid!\n" +
			"The number of calendar cycles is zero!\n")
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	err = targetDateTimeDto.IsValidDateTime(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseDateTimeDto := calCycleCfg.GetJDNBaseStartYearDateTime()

	// Base Date Time for Gregorian Calendar:
	// November 24, -4713 12:00:00.000000000 UTC (Noon)
	var baseTargetComparisonResult, baseRemainingDaysInYear,
	targetRemainingDaysInYear, targetYearOrdinalNumber,
	julianDayNumberSign int

	var mainCycleAdjustmentDays, remainderYears, ordinalDays *big.Int

	baseTargetComparisonResult, err = baseDateTimeDto.Compare(&targetDateTimeDto, ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	baseDateTimeDto.SetTag("baseDateTimeDto")
	targetDateTimeDto.SetTag("targetDateTimeDto")


	targetYearOrdinalNumber, err = targetDateTimeDto.GetOrdinalDayNumber(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	targetRemainingDaysInYear, err = targetDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Target Days Remaining in Year: %v\n",
		targetRemainingDaysInYear)

	baseRemainingDaysInYear, err = baseDateTimeDto.GetRemainingDaysInYear(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	fmt.Printf("Base Days Remaining in Year: %v\n",
		baseRemainingDaysInYear)

	var mainCycleStartDateTime DateTimeTransferDto

	// Primary Decision Tree
	if baseTargetComparisonResult == 0 {
		// base and target have equivalent date/times
		// Julian Day Number is Zero and time fraction is Zero.

		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err

	} else if baseTargetComparisonResult == 1 {
		// base is greater than target date/time
		// target date/time must be negative
		// Julian Day Number must be negative

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForNegativeJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = -1

		// RemainderYears =
		// BaseYear - TargetYear -1
		remainderYears =
			big.NewInt(
				mainCycleStartDateTime.GetYear() -
					targetDateTimeDto.GetYear() - 1 )

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForNegativeJDNNo()

	} else {
		// target is greater than base date/time
		// base is less than target date/time
		// target date/time could be positive or negative
		// Use Main Cycle Start Date/Time which is less Than target

		mainCycleStartDateTime =
			calCycleCfg.GetMainCycleStartDateForPositiveJDNNo()

		mainCycleStartDateTime.SetTag("Main Cycle Start Date Time")

		fmt.Printf("%s\n",
			mainCycleStartDateTime.String())

		julianDayNumberSign = 1

		// RemainderYears or Whole Years Interval =
		// TargetYear - mainCycleStartDateTime -1
		remainderYears =
			big.NewInt(
				targetDateTimeDto.GetYear() -
					mainCycleStartDateTime.GetYear() - 1)

		mainCycleAdjustmentDays =
			calCycleCfg.GetMainCycleAdjustmentDaysForPositiveJDNNo()

	}

	fmt.Printf("Initial Remainder Years: %v\n",
		remainderYears.Text(10))

	ordinalDays =
		big.NewInt(
			int64(
				baseRemainingDaysInYear +
					targetYearOrdinalNumber))

	cycles := calCycleCfg.GetCalendarCycleConfigurations()
	daysDuration := big.NewInt(0)

	if remainderYears.Cmp(big.NewInt(0)) == 1 {

		var cycleCount *big.Int
		totalYears := big.NewInt(0)

		for i:=0; i < len(cycles); i++ {

			cycleCount =
				big.NewInt(0).
					Quo(remainderYears,
						cycles[i].GetYearsInCycle())

			remainderYears =
				big.NewInt(0).
					Sub(remainderYears,
						big.NewInt(0).
							Mul(cycleCount, cycles[i].GetYearsInCycle()))

			fmt.Printf("cycles[%v].cycleCount=%v\n",
				i, cycleCount.Text(10))

			fmt.Printf("cycles[%v].yearsInCycle=%v\n",
				i, cycles[i].GetYearsInCycle().Text(10))

			fmt.Printf("cycles[%v].remainderYears=%v\n",
				i, remainderYears.Text(10))


			cycles[i].SetCycleCount(cycleCount)
			cycles[i].SetRemainderYears(remainderYears)

			totalYears =
				big.NewInt(0).
					Add(totalYears, cycles[i].GetCycleCountTotalYears())

			fmt.Printf("cycles[%v].CycleCountTotalDays=%v\n",
				i, cycles[i].GetCycleCountTotalDays().Text(10))

			daysDuration =
				big.NewInt(0).
					Add(daysDuration, cycles[i].GetCycleCountTotalDays())
		}

		fmt.Printf("daysDuration: %v\n",
			daysDuration.Text(10))

		daysDuration =
			big.NewInt(0).
				Add(daysDuration, mainCycleAdjustmentDays)

		fmt.Printf("Adjusted daysDuration: %v\n",
			daysDuration.Text(10))

	}

	fmt.Printf("ordinalDays: %v\n",
		ordinalDays.Text(10))

	daysDuration =
		big.NewInt(0).
			Add(daysDuration, ordinalDays)

	fmt.Printf("Final daysDuration: %v\n",
		daysDuration.Text(10))

	if !daysDuration.IsInt64() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: 'daysDuration' is too large and cannot convert to int64!\n" +
			"daysDuration='%v'\n",
			daysDuration.Text(10))
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}

	julianDayNumber = daysDuration.Int64()

	// Compute Time Components

	noonNanoseconds := int64(time.Hour * 12)

	var targetDayTotalNanoseconds int64

	targetDayTotalNanoseconds, err =
		targetDateTimeDto.GetTotalTimeInNanoseconds(ePrefix)

	if err != nil {
		return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
	}


	twentyFourHourNanosecondsFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(int64(time.Hour) * 24)

	julianDayNoTimeAdjustment := int64(0)

	if targetDayTotalNanoseconds == noonNanoseconds {

		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetFloat64(0.0)

	} else if targetDayTotalNanoseconds < noonNanoseconds {
		targetDayTotalNanoseconds += noonNanoseconds

		julianDayNoTimeAdjustment--

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)


	} else {
		// targetDayTotalNanoseconds > noonNanoseconds


		targetDayTotalNanoseconds -= noonNanoseconds

		targetNanosecondFloat :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				SetInt64(targetDayTotalNanoseconds)


		julianDayNumberTimeFraction =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Quo(targetNanosecondFloat, twentyFourHourNanosecondsFloat)
	}

	julianDayNumber += julianDayNoTimeAdjustment

	fmt.Printf("julianDayNoTimeAdjustment: %v\n",
		julianDayNoTimeAdjustment)

	julianDayNoFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64(julianDayNumber)

	julianDayNumberTime =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Add(julianDayNoFloat, julianDayNumberTimeFraction)


	if julianDayNumberSign == -1 {

		julianDayNumber *= -1

		julianDayNumberTime =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Neg(julianDayNumberTime)
	}

	separator := strings.Repeat("-", 75)
	fmt.Println()
	fmt.Println(separator)

	fmt.Printf("Final Julian Day Number: %v\n",
		julianDayNumber)

	fmt.Println(separator)
	fmt.Println(baseDateTimeDto.String())
	fmt.Println(targetDateTimeDto.String())
	fmt.Println(separator)

	return julianDayNumber, julianDayNumberTime, julianDayNumberTimeFraction, err
}
