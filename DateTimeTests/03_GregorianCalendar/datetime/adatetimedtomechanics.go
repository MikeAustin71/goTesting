package datetime

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

type aDateTimeDtoMechanics struct {
	lock *sync.Mutex
}

// Empty - Resets the internal data fields of 'dTimeTransDto'
// to invalid values. Effectively, the 'dTimeTransDto' instance
// is rendered blank, invalid and empty.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) empty(
	dTimeTransDto *ADateTimeDto,
	ePrefix string) error {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	ePrefix += "aDateTimeDtoMechanics.empty() "

	if dTimeTransDto == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'dTimeTransDto' " +
			"has a 'nil' pointer!\n")
	}

	err := dTimeTransDto.date.Empty(ePrefix)

	if err != nil {
		return err
	}

	err = dTimeTransDto.time.Empty(ePrefix)

	if err != nil {
		return err
	}

	err = dTimeTransDto.julianDayNumber.Empty(ePrefix)

	if err != nil {
		return err
	}

	dTimeTransDto.dateTimeFmt = ""

	dTimeTransDto.tag = ""

	return nil
}

// exchangeValues - Performs a data exchange. The data fields
// from ADateTimeDto instance 'aDateTimeDto1' are copied to the
// ADateTimeDto instance, 'aDateTimeDto2'. In turn, the data
// fields from 'aDateTimeDto2' are in turn copied to 'aDateTimeDto1'.
//
// This is a copy procedures employ a deep copy operation.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  aDateTimeDto1       *ADateTimeDto
//     - A pointer to an ADateTimeDto instance. The internal data
//       values from input parameter 'aDateTimeDto2' will be copied
//       to the data fields of this instance, 'aDateTimeDto1'.
//
//  aDateTimeDto2       *ADateTimeDto
//     - A pointer to another ADateTimeDto instance. The internal data
//       values of this second instance will be populated with data fields
//       from the input parameter instance, 'aDateTimeDto1'.
//
//
//  ePrefix             string
//     - Error Prefix. A string consisting of the method chain used
//       to call this method. In case of error, this text string is
//       included in the error message. Note: Be sure to leave a space
//       at the end of 'ePrefix'. If no Error Prefix is desired, simply
//       provide an empty string for this parameter.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  err                error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) exchangeValues(
	aDateTimeDto1 *ADateTimeDto,
	aDateTimeDto2 *ADateTimeDto,
	ePrefix string) error {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	ePrefix += "aDateTimeDtoMechanics.exchangeValues() "

	if aDateTimeDto1 == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'aDateTimeDto1' is INVALID!\n" +
			"aDateTimeDto1 == nil pointer\n")
	}

	if aDateTimeDto2 == nil {
		return errors.New(ePrefix + "\n" +
			"Input parameter 'aDateTimeDto2' is INVALID!\n" +
			"aDateTimeDto2 == nil pointer\n")
	}

	if aDateTimeDto1.lock == nil {
		aDateTimeDto1.lock = new(sync.Mutex)
	}

	if aDateTimeDto2.lock == nil {
		aDateTimeDto2.lock = new(sync.Mutex)
	}

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_, err := aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto1,
		ePrefix + " Checking Input Parameter Validity: aDateTimeDto1 ")

	if err != nil {
		return err
	}

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto2,
		ePrefix + " Checking Input Parameter Validity: aDateTimeDto2 ")

	if err != nil {
		return err
	}

	err = aDateTimeDto1.date.ExchangeValues(
		&aDateTimeDto2.date,
		ePrefix)

	if err != nil {
		return err
	}

	err = aDateTimeDto1.time.ExchangeValues(
		&aDateTimeDto2.time,
		ePrefix)

	if err != nil {
		return err
	}

	temp := aDateTimeDto1.tag

	aDateTimeDto1.tag = aDateTimeDto2.tag

	aDateTimeDto2.tag = temp

	var tempJDN JulianDayNoDto

	tempJDN, err =
		aDateTimeDto1.julianDayNumber.CopyOut(
			ePrefix + "Saving Temp Julian Day Number ")

	if err != nil {
		return err
	}

	err = aDateTimeDto1.julianDayNumber.Empty(
		ePrefix + "Emptying 'aDateTimeDto1.julianDayNumber' ")

	if err != nil {
		return err
	}

	aDateTimeDto1.julianDayNumber,
		err =
		aDateTimeDto2.julianDayNumber.CopyOut(
			ePrefix + "Copying JDN from 'aDateTimeDto2' to 'aDateTimeDto1' \n" )

	err = aDateTimeDto2.julianDayNumber.Empty(
		ePrefix + "Emptying 'aDateTimeDto2.julianDayNumber' ")

	if err != nil {
		return err
	}

	aDateTimeDto2.julianDayNumber,
		err = tempJDN.CopyOut(
		ePrefix + "Copying JDN from 'tempJDN' to 'aDateTimeDto2' \n")

	return nil
}

// processAmPm - processes and returns correct AM/PM format
// for 12-Hour presentations.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processAmPm(
	inputStr string,
	hourNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processAmPm() "

	if hourNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is LESS THAN ONE!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if hourNumber > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is GREATER THAN 23!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "PM") {

		amPmStr := "PM"

		if hourNumber < 1 && hourNumber > 0 {
			amPmStr = "AM"
		}

		tokenMap["!AMPMUpperCase!"] = amPmStr

		resultStr = strings.Replace(resultStr,
			"PM",
			"!AMPMUpperCase!",
			1)

	} else if strings.Contains(resultStr, "pm") {

		amPmStr := "pm"

		if hourNumber < 1 && hourNumber > 0 {
			amPmStr = "am"
		}

		tokenMap["!AMPMLowerCase!"] = amPmStr

		resultStr = strings.Replace(resultStr,
			"pm",
			"!AMPMLowerCase!",
			1)

	}

	return resultStr, err
}

// processDateDay - processes and returns correct date day format
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processDateDay(
	inputStr string,
	dateDayNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()
	aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processDateDay() "

	if dateDayNumber < 1 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dateDayNumber' is LESS THAN ONE!\n" +
			"dateDayNumber='%v'\n", dateDayNumber)

		return resultStr, err
	}

	if dateDayNumber > 31 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dateDayNumber' is GREATER THAN 31!\n" +
			"dateDayNumber='%v'\n", dateDayNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "02") {

		dateDayStr := fmt.Sprintf("%02d",
			dateDayNumber)

		tokenMap["!DateDayTwoDigit!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"02",
			"!DateDayTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "_2") {

		var dateDayStr string

		if dateDayNumber < 10 {
			dateDayStr = fmt.Sprintf("_%d",
				dateDayNumber)
		} else {

			dateDayStr = fmt.Sprintf("%d",
				dateDayNumber)
		}

		tokenMap["!DateDayLeadUnderScore!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"_2",
			"!DateDayLeadUnderScore!",
			1)

	} else if strings.Contains(resultStr, "2") {

		dateDayStr := fmt.Sprintf("%02d",
			dateDayNumber)

		tokenMap["!DateDayOneDigit!"] = dateDayStr

		resultStr = strings.Replace(resultStr,
			"2",
			"!DateDayOneDigit!",
			1)

	}

	return resultStr, err
}

// processDayOfWeek - processes and returns correct day of week format
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processDayOfWeek(
	inputStr string,
	dayOfWeekNumber UsDayOfWeekNo,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	ePrefix += "aDateTimeDtoMechanics.processDayOfWeek() "

	resultStr = inputStr

	if ! dayOfWeekNumber.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'dayOfWeekNumber' is INVALID!\n" +
			"dayOfWeekNumber='%v'\n", dayOfWeekNumber.String())

		return resultStr, err
	}

	var dayOfWeekAbbrv string

	// Process Day Of Week
	if strings.Contains(resultStr, "Monday") {

		tokenMap["!DayOfWeek!"] = dayOfWeekNumber.String()

		resultStr = strings.Replace(resultStr,
			"Monday",
			"!DayOfWeek!",
			1)

	} else if strings.Contains(resultStr, "Mon") {


		dayOfWeekAbbrv, err = dayOfWeekNumber.XDayOfWeekNameAbbreviation(
			3,
			ePrefix + "dayOfWeekNumber.XDayOfWeekNameAbbreviation() ")

		if err != nil {
			return resultStr, err
		}

		tokenMap["!DayOfWeek!"] = dayOfWeekAbbrv

		resultStr = strings.Replace(resultStr,
			"Mon",
			"!DayOfWeek!",
			1)
	}

	return resultStr, err
}

// processHours - processes and returns correct hour format
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processHours(
	inputStr string,
	hourNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processHours() "

	if hourNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is LESS THAN ONE!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if hourNumber > 23 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'hourNumber' is GREATER THAN 23!\n" +
			"hourNumber='%v'\n", hourNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "15") {

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwentyFourTwoDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"15",
			"!HourTwentyFourTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "03") {

		if hourNumber > 12 {
			hourNumber -= 12
		}

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwelveTwoDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"03",
			"!HourTwelveTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "3") {

		if hourNumber > 12 {
			hourNumber -= 12
		}

		hourStr := fmt.Sprintf("%02d", hourNumber)

		tokenMap["!HourTwelveOneDigit!"] = hourStr

		resultStr = strings.Replace(resultStr,
			"3",
			"!HourTwelveOneDigit!",
			1)
	}

	return resultStr, err
}


// processMicroseconds - processes and returns correct microseconds format
// Make certain this method is called after 'processNanoseconds()'.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processMicroseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processMicroseconds() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	microsecondsNumber := nanosecondsNumber / 1000

	if strings.Contains(resultStr, ".000000") {

		microsecondStr := fmt.Sprintf("%06d", microsecondsNumber)

		tokenMap["!MicrosecondsTrailingZeros!"] = microsecondStr

		resultStr = strings.Replace(resultStr,
			".000000",
			".!MicrosecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999999") {

		microsecondStr := fmt.Sprintf("%d", microsecondsNumber)

		tokenMap["!MicrosecondsNoTrailingZeros!"] = microsecondStr

		resultStr = strings.Replace(resultStr,
			".999999",
			".!MicrosecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processMilliseconds - processes and returns correct milliseconds format
// Make certain this method is called after 'processNanoseconds()' and
// 'processMicroseconds()'.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processMilliseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processMilliseconds() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	millisecondsNumber := nanosecondsNumber / 1000000

	if strings.Contains(resultStr, ".000") {

		millisecondStr := fmt.Sprintf("%03d", millisecondsNumber)

		tokenMap["!MillisecondsTrailingZeros!"] = millisecondStr

		resultStr = strings.Replace(resultStr,
			".000",
			".!MillisecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999") {

		millisecondStr := fmt.Sprintf("%d", millisecondsNumber)

		tokenMap["!MillisecondsNoTrailingZeros!"] = millisecondStr

		resultStr = strings.Replace(resultStr,
			".999",
			".!MillisecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processMinutes - processes and returns correct minute format
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processMinutes(
	inputStr string,
	minuteNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processMinutes() "

	if minuteNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minuteNumber' is LESS THAN ONE!\n" +
			"minuteNumber='%v'\n", minuteNumber)

		return resultStr, err
	}

	if minuteNumber > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'minuteNumber' is GREATER THAN 59!\n" +
			"minuteNumber='%v'\n", minuteNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "04") {

		minuteStr := fmt.Sprintf("%02d", minuteNumber)

		tokenMap["!MinutesTwoDigit!"] = minuteStr

		resultStr = strings.Replace(resultStr,
			"04",
			"!MinutesTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "4") {

		minuteStr := fmt.Sprintf("%d", minuteNumber)

		tokenMap["!MinutesOneDigit!"] = minuteStr

		resultStr = strings.Replace(resultStr,
			"4",
			"!MinutesOneDigit!",
			1)

	}

	return resultStr, err
}

// processMonths - processes and returns correct month
// formatting.
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processMonths(
	inputStr string,
	monthNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processAmPm() "

	monthsOfYear := map[int] string {
		1: "January",
		2: "February",
		3: "March",
		4: "April",
		5: "May",
		6: "June",
		7: "July",
		8: "August",
		9: "September",
		10: "October",
		11: "November",
		12: "December",
	}

	monthStr := monthsOfYear[monthNumber]

	if monthStr == "" {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input Parameter 'monthNumber' is INVALID!\n" +
			"monthNumber='%v'\n", monthNumber)
		return resultStr, err
	}

	// Process Months
	if strings.Contains(inputStr, "January") {

		tokenMap["!MonthName!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"January",
			"!MonthName!",
			1)

	} else if strings.Contains(inputStr, "Jan") {

		monthStr = monthStr[0:3]

		tokenMap["!MonthNameAbbrv!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"January",
			"!MonthNameAbbrv!",
			1)

	} else if strings.Contains(inputStr, "01") {

		monthStr = fmt.Sprintf("%02d", monthNumber)

		tokenMap["!MonthNumberTwoDigit!"] = monthStr

		resultStr = strings.Replace(resultStr,
			"01",
			"!MonthNumberTwoDigit!",
			1)

	} else if strings.Contains(inputStr, "1") {

		lenInputStr := len(inputStr)
		lastStrIdx := lenInputStr - 1

		for i:=0; i < lenInputStr; i++ {

			if inputStr[i] == '1' {

				if i == lastStrIdx {

					monthStr = fmt.Sprintf("%d", monthNumber)
					tokenMap["!MonthNumberOneDigit!"] = monthStr

					resultStr = resultStr[0:i] + "!MonthNumberOneDigit!"

					break
				} else {

					if inputStr[i+1] == '5' {
						continue

					} else {

						monthStr = fmt.Sprintf("%d", monthNumber)

						tokenMap["!MonthNumberOneDigit!"] = monthStr

						resultStr = resultStr[0:i+1] +
							"!MonthNumberOneDigit!" +
							resultStr[i+2:]
						break
					}
				}
			}
		}
	}

	return resultStr, err
}

// processNanoseconds - processes and returns correct subMicrosecondNanoseconds format
// Make certain to call this method before calling 'processMicroseconds()'
// and 'processMilliseconds()'.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processNanoseconds(
	inputStr string,
	nanosecondsNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	ePrefix += "aDateTimeDtoMechanics.processAmPm() "

	if nanosecondsNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is LESS THAN ZERO!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if nanosecondsNumber > 999999999 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'nanosecondsNumber' is GREATER THAN 999,999,999!\n" +
			"nanosecondsNumber='%v'\n", nanosecondsNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, ".000000000") {

		nanosecondStr := fmt.Sprintf("%09d", nanosecondsNumber)

		tokenMap["!NanosecondsTrailingZeros!"] = nanosecondStr

		resultStr = strings.Replace(resultStr,
			".000000000",
			".!NanosecondsTrailingZeros!",
			1)

	} else if strings.Contains(resultStr, ".999999999") {

		nanosecondStr := fmt.Sprintf("%d", nanosecondsNumber)

		tokenMap["!NanosecondsNoTrailingZeros!"] = nanosecondStr

		resultStr = strings.Replace(resultStr,
			".999999999",
			".!NanosecondsNoTrailingZeros!",
			1)

	}

	return resultStr, err
}

// processOffset - processes and returns offset hours and
// offset minutes format. If the 'offSetHours' value or the
// 'offSetMinutes' value is invalid, the offset will be deleted
// from the formatted date/time string.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processOffset(
	inputStr string,
	offSetHours int,
	offSetMinutes int,
	tokenMap map[string]string) (resultStr string) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	isValid := true

	if offSetHours < -23 || offSetHours > 23 {
		isValid = false
	}

	if offSetMinutes < -59 || offSetMinutes > 59 {
		isValid = false
	}

	var offsetFmtStr string
	var keys []string


	if !isValid {

		offsetFmtStr = ""

		keys = []string {
			" Z0700",
			"Z0700",
			" Z07:00",
			"Z07:00",
			" -0700",
			"-0700",
			" -07:00",
			"-07:00",
			" -07",
			"-07",
		}

		lenKeys := len(keys)

		for i:=0; i < lenKeys; i++ {

			if strings.Contains(resultStr, keys[i]){

				tokenMap["!OffsetUTC!"] = offsetFmtStr

				resultStr = strings.Replace(resultStr,
					keys[i],
					"!OffsetUTC!",
					1)
				break
			}

		}

		return resultStr
	}

	keys = []string {
		"Z0700",
		"Z07:00",
		"-0700",
		"-07:00",
		"-07",
	}

	numberSign := 1

	if offSetHours < 0 {
		numberSign = -1
		offSetHours *= -1
	}

	if offSetMinutes < 0 {
		offSetMinutes *= -1
	}

	zPrefix := "Z"

	leadPrefix := ""

	if numberSign == -1 {
		zPrefix = "Z-"
		leadPrefix = "-"
	}

	values := []string {
		fmt.Sprintf(zPrefix + "%02d%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(zPrefix + "%02d:%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d:%02d",
			offSetHours,
			offSetMinutes),
		fmt.Sprintf(leadPrefix + "%02d",
			offSetHours),
	}

	lenKeys := len(keys)

	for i:=0; i < lenKeys; i++ {

		if strings.Contains(resultStr, keys[i]){

			tokenMap["!OffsetUTC!"] = values[i]

			resultStr = strings.Replace(resultStr,
				keys[i],
				"!OffsetUTC!",
				1)
			break
		}
	}

	return resultStr
}

// processSeconds - processes and returns correct second format
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processSeconds(
	inputStr string,
	secondNumber int,
	tokenMap map[string]string,
	ePrefix string) (resultStr string, err error) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr
	err = nil

	ePrefix += "aDateTimeDtoMechanics.processSeconds() "

	if secondNumber < 0 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'secondNumber' is LESS THAN ONE!\n" +
			"secondNumber='%v'\n", secondNumber)

		return resultStr, err
	}

	if secondNumber > 59 {
		err = fmt.Errorf(ePrefix + "\n" +
			"Error: Input parameter 'secondNumber' is GREATER THAN 59!\n" +
			"secondNumber='%v'\n", secondNumber)

		return resultStr, err
	}

	if strings.Contains(resultStr, "05") {

		secondStr := fmt.Sprintf("%02d", secondNumber)

		tokenMap["!SecondsTwoDigit!"] = secondStr

		resultStr = strings.Replace(resultStr,
			"05",
			"!SecondsTwoDigit!",
			1)

	} else if strings.Contains(resultStr, "5") {

		secondStr := fmt.Sprintf("%d", secondNumber)

		tokenMap["!SecondsOneDigit!"] = secondStr

		resultStr = strings.Replace(resultStr,
			"5",
			"!SecondsOneDigit!",
			1)

	}

	return resultStr, err
}

// processYears - Returns the formatted year in target date time string
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processYears(
	inputStr string,
	year int64,
	tokenMap map[string]string) (resultStr string) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	// Process Years
	if strings.Contains(inputStr, "2006") {

		tokenMap["!YearFourDigit!"] =
			fmt.Sprintf("%d", year)

		resultStr = strings.Replace(resultStr,
			"2006",
			"!YearFourDigit!",
			1)

	} else if strings.Contains(inputStr, "06") {

		yearStr := fmt.Sprintf("%02d",
			year % 100)

		tokenMap["!YearTwoDigit!"] =
			yearStr

		resultStr = strings.Replace(resultStr,
			"06",
			"!YearTwoDigit!",
			1)

	} else if strings.Contains(inputStr, "6") {

		yearStr := fmt.Sprintf("%d",
			year % 100)

		tokenMap["!YearOneDigit!"] =
			yearStr

		resultStr = strings.Replace(resultStr,
			"6",
			"!YearOneDigit!",
			1)
	}

	return resultStr
}

// processTimeZone - Returns the formatted time zone in
// the target date time string.
//
func (aDateTimeDtoMech *aDateTimeDtoMechanics) processTimeZone(
	inputStr string,
	timeZoneAbbrv string,
	tokenMap map[string]string) (resultStr string) {

	if aDateTimeDtoMech.lock == nil {
		aDateTimeDtoMech.lock = new(sync.Mutex)
	}

	aDateTimeDtoMech.lock.Lock()

	defer aDateTimeDtoMech.lock.Unlock()

	resultStr = inputStr

	// Process Time Zones
	if strings.Contains(inputStr, "MST") {

		tokenMap["!TimeZone!"] = timeZoneAbbrv

		resultStr = strings.Replace(resultStr,
			"MST",
			"!TimeZone!",
			1)

	}

	return resultStr
}
