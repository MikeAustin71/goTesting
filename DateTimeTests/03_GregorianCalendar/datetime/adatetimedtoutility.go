package datetime

import (
	"errors"
	"strings"
	"sync"
)

type aDateTimeDtoUtility struct {
	lock *sync.Mutex
}

// copyIn - Receives a pointer to an incoming ADateTimeDto instance
// (dTimeTransDto2) and copies all of the internal data fields to
// ADateTimeDto instance dTimeTransDto1.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  aDateTimeDto1       *ADateTimeDto
//    - Data field values from the 'aDateTimeDto2' instance
//      will be copied into the data fields of ADateTimeDto
//      this instance, 'aDateTimeDto1' . If the method completes
//      successfully, data fields in both instances will have
//      identical values.
//
//  aDateTimeDto2       *ADateTimeDto
//    - Data field values from this ADateTimeDto instance
//      will be copied into the ADateTimeDto instance,
//      'aDateTimeDto1'. If the method completes successfully,
//      data fields in both instances will have identical values.
//
//  ePrefix             string
//    - This is an error prefix which is included in all returned
//      error messages. Usually, it contains the names of the calling
//      method or methods.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//
//  error
//     - If successful the returned error Type is set equal to 'nil'.
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. Note this error
//       message will incorporate the method chain and text passed
//       by input parameter, 'ePrefix'.
//
func (aDateTimeDtoUtil *aDateTimeDtoUtility) copyIn(
	aDateTimeDto1 *ADateTimeDto,
	aDateTimeDto2 *ADateTimeDto,
	ePrefix string) error {

	if aDateTimeDtoUtil.lock == nil {
		aDateTimeDtoUtil.lock = new(sync.Mutex)
	}

	aDateTimeDtoUtil.lock.Lock()

	defer aDateTimeDtoUtil.lock.Unlock()

	ePrefix += "aDateTimeDtoUtility.copyIn() "

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
		ePrefix + " Checking Validity: aDateTimeDto1 ")
	
	if err != nil {
		return err
	}

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto2,
		ePrefix + " Checking Validity: aDateTimeDto2 ")

	if err != nil {
		return err
	}

	err =
	aDateTimeDto1.date.CopyIn(
		&aDateTimeDto2.date,
		ePrefix)

	if err != nil {
		return err
	}

	err =
	aDateTimeDto1.time.CopyIn(
		&aDateTimeDto2.time,
		ePrefix)

	if err != nil {
		return err
	}

	aDateTimeDto1.tag = aDateTimeDto2.tag

	aDateTimeDto1.julianDayNumber, err =
		aDateTimeDto2.julianDayNumber.CopyOut(ePrefix)

	if err != nil {
		return err
	}

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto1,
		ePrefix + " Checking Output Validity: aDateTimeDto1 ")

	return err
}

// CopyOut - Returns a deep copy of the input parameter,
// 'dTimeTransDto'; an instance of ADateTimeDto.
//
func (aDateTimeDtoUtil *aDateTimeDtoUtility) copyOut(
	aDateTimeDto *ADateTimeDto,
	ePrefix string) (ADateTimeDto, error) {

	if aDateTimeDtoUtil.lock == nil {
		aDateTimeDtoUtil.lock = new(sync.Mutex)
	}

	aDateTimeDtoUtil.lock.Lock()

	defer aDateTimeDtoUtil.lock.Unlock()

	ePrefix += "aDateTimeDtoUtility.copyOut() "

	newDTimeTransDto := ADateTimeDto{}
	
	if aDateTimeDto == nil {
		return newDTimeTransDto, errors.New(ePrefix + "\n" +
			"Input parameter 'aDateTimeDto' is INVALID!\n" +
			"aDateTimeDto == nil pointer!\n")
	}

	if aDateTimeDto.lock == nil {
		aDateTimeDto.lock = new(sync.Mutex)
	}

	aDateTimeDtoNanobot := aDateTimeDtoNanobot{}

	_, err := aDateTimeDtoNanobot.testDateTransferDtoValidity(
		aDateTimeDto,
		ePrefix + " Checking Input Parameter Validity: 'aDateTimeDto' ")

	if err != nil {
		return newDTimeTransDto, err
	}

	newDTimeTransDto.date, err = aDateTimeDto.date.CopyOut(ePrefix)

	if err != nil {
		return ADateTimeDto{}, err
	}

	newDTimeTransDto.time, err = aDateTimeDto.time.CopyOut(ePrefix)

	if err != nil {
		return ADateTimeDto{}, err
	}

	newDTimeTransDto.julianDayNumber, err =
		aDateTimeDto.julianDayNumber.CopyOut(ePrefix)

	if err != nil {
		return ADateTimeDto{}, err
	}

	newDTimeTransDto.tag = aDateTimeDto.tag

	newDTimeTransDto.lock = new(sync.Mutex)

	_, err = aDateTimeDtoNanobot.testDateTransferDtoValidity(
		&newDTimeTransDto,
		ePrefix + " Checking Output Validity: newDTimeTransDto ")

	return newDTimeTransDto, err
}

// generateDateTimeStr - Converts input years, months, days, hours,
// minutes, seconds and subMicrosecondNanoseconds to a formatted date time string
// the golang format string passed in input parameter 'dateFormatStr'.
//
func (aDateTimeDtoUtil *aDateTimeDtoUtility) generateDateTimeStr(
	year int64,
	month,
	days int,
	usDayOfWeekNumber UsDayOfWeekNo,
	hours,
	minutes,
	seconds,
	nanoseconds int,
	dateFormatStr,
	ePrefix string) (string, error) {

// TODO - What about time zone?

	if aDateTimeDtoUtil.lock == nil {
		aDateTimeDtoUtil.lock = new(sync.Mutex)
	}

	aDateTimeDtoUtil.lock.Lock()

	defer aDateTimeDtoUtil.lock.Unlock()

	ePrefix += "aDateTimeDtoUtility.generateDateTimeStr() "

	var err error

	/*	replacementTokens := map[string]string{
			"!YearFourDigit!":"",
			"!YearTwoDigit!":"",
			"!YearOneDigit!":"",
			"!DayOfWeek!":"",
			"!DayOfWeekAbbrv!":"",
			"!MonthName!":"",
			"!MonthNameAbbrv!":"",
			"!MonthNumberTwoDigit!":"",
			"!MonthNumberOneDigit!":"",
			"!DateDayTwoDigit!":"",
			"!DateDayLeadUnderScore!":"",
			"!DateDayOneDigit!":"",
			"!HourTwentyFourTwoDigit!":"",
			"!HourTwelveTwoDigit!":"",
			"!HourTwelveOneDigit!":"",
			"!AMPMUpperCase!",
			"!AMPMLowerCase!",
			"!MinutesTwoDigit!":"",
			"!MinutesOneDigit!":"",
			"!SecondsTwoDigit!":"",
			"!SecondsOneDigit!":"",
			"!NanosecondsTrailingZeros!":"",
			"!NanosecondsNoTrailingZeros!":"",
			"!MillisecondsTrailingZeros!":"",
			"!MillisecondsNoTrailingZeros!":"",
			"!MicrosecondsTrailingZeros!":"",
			"!MicrosecondsNoTrailingZeros!":"",
			"!OffsetUTC!":"",
			"!TimeZone!":"",
		}
	*/

	replacementTokens := map[string]string{}

	dtMech := DTimeNanobot{}

	resultStr := dtMech.PreProcessDateFormatStr(dateFormatStr)

	aDateTimeDtoMech := aDateTimeDtoMechanics{}

	resultStr, err = aDateTimeDtoMech.processDayOfWeek(
		resultStr,
		usDayOfWeekNumber,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = aDateTimeDtoMech.processYears(
		resultStr,
		year,
		replacementTokens)

	resultStr, err = aDateTimeDtoMech.processMonths(
		resultStr,
		month,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processDateDay(
		resultStr,
		days,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processHours(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processMinutes(
		resultStr,
		minutes,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processSeconds(
		resultStr,
		seconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processNanoseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processMicroseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processMilliseconds(
		resultStr,
		nanoseconds,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr, err = aDateTimeDtoMech.processAmPm(
		resultStr,
		hours,
		replacementTokens,
		ePrefix)

	if err != nil {
		return resultStr, err
	}

	resultStr = aDateTimeDtoMech.processOffset(
		resultStr,
		-999,
		-999,
		replacementTokens)

	resultStr = aDateTimeDtoMech.processTimeZone(
		resultStr,
		"UTC",
		replacementTokens)

	for key, value := range replacementTokens {

		resultStr = strings.Replace(resultStr,key,value,1)

	}

	return resultStr, err
}

