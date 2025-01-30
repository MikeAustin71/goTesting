package datetime

import (
	"errors"
	"fmt"
	"sync"
)

type dateTransferDtoNanobot struct {
	lock *sync.Mutex
}


// testDateTransferDtoValidity - Checks the validity of a DateTransferDto
// instance. If the instance is valid this method returns 'true' and sets
// the returned error type to nil.
//
// If the instance is invalid, the method returns 'false' plus an
// error message.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  dateTransDto        *DateTransferDto
//     - A pointer to an instance of DateTransferDto. This method will
//       NOT change the values of internal member variables to achieve
//       the method's objectives. Member variables will be tested for
//       validity.
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Note: Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If the method completes successfully, this boolean value will
//       signal 'false' if 'dateTransDto' member variables contains an
//       invalid value. Absent processing errors, if the validity tests
//       show that 'dateTransDto' member variables values are valid, this
//       'isValid' flag will be set to 'true'.
//
//
//  err                 error
//     - If this method is successful, the returned error Type is set
//       equal to 'nil'. If errors are encountered during processing,
//       the returned error Type will encapsulate an error message.
//       Note that this error message will incorporate the method
//       chain and text passed by input parameter, 'ePrefix'.
//
//       An error will also be returned if the method determines that
//       one or more of the 'dateTransDto' member variables contain
//       invalid values.
//
func (dateTransDtoNanobot *dateTransferDtoNanobot) testDateTransferDtoValidity(
	dateTransDto *DateTransferDto,
	ePrefix string) ( isValid bool, err error) {

	if dateTransDtoNanobot.lock == nil {
		dateTransDtoNanobot.lock = new(sync.Mutex)
	}

	dateTransDtoNanobot.lock.Lock()

	defer dateTransDtoNanobot.lock.Unlock()

	ePrefix += "dateTransferDtoNanobot.testDateTransferDtoValidity() "

	isValid = false
	err = nil

	if dateTransDto == nil {
		err = errors.New(ePrefix +
			"\nInput parameter 'dateTransDto' is INVALID!\n" +
			"dateTransDto=nil!")
		return isValid, err
	}

	if dateTransDto.lock == nil {
		dateTransDto.lock = new(sync.Mutex)
	}

	if !dateTransDto.yearNumberingMode.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"The DateTransferDto instance is corrupted because\n" +
			"'DateTransferDto.yearNumberingMode' is INVALID!\n" +
			"dateTransDto.yearNumberingMode='%v'\n",
			dateTransDto.yearNumberingMode.XValueInt())

		return isValid, err
	}

	if !dateTransDto.yearNumType.XIsValid() {
		err = fmt.Errorf(ePrefix + "\n" +
			"The DateTransferDto instance is corrupted because\n" +
			"'DateTransferDto.yearNumType' is INVALID!\n" +
			"dateTransDto.yearNumType='%v'\n",
			dateTransDto.yearNumType.XValueInt())

		return isValid, err
	}

	if dateTransDto.yearNumberingMode == CalendarYearNumMode(0).Astronomical() &&
		dateTransDto.yearNumType != CalendarYearNumType(0).Astronomical() {
		err = fmt.Errorf(ePrefix + "\n" +
			"The DateTransferDto instance is corrupted because\n" +
			"'DateTransferDto.yearNumberingMode' and 'DateTransferDto.yearNumType'\n" +
			"are NOT in sync.\n" +
			"dateTransDto.yearNumberingMode='%v'\n" +
			"dateTransDto.yearNumType='%v'\n",
			dateTransDto.yearNumberingMode.String(),
			dateTransDto.yearNumType.String())

		return isValid, err
	}

	if dateTransDto.yearNumberingMode == CalendarYearNumMode(0).CommonEra() &&
		dateTransDto.yearNumType != CalendarYearNumType(0).BCE() &&
		dateTransDto.yearNumType != CalendarYearNumType(0).CE() {
		err = fmt.Errorf(ePrefix + "\n" +
			"The DateTransferDto instance is corrupted because\n" +
			"'DateTransferDto.yearNumberingMode' and 'DateTransferDto.yearNumType'\n" +
			"are NOT in sync.\n" +
			"dateTransDto.yearNumberingMode='%v'\n" +
			"dateTransDto.yearNumType='%v'\n",
			dateTransDto.yearNumberingMode.String(),
			dateTransDto.yearNumType.String())

		return isValid, err
	}

	if dateTransDto.calendarBaseData == nil {
		err = errors.New(ePrefix + "\n" +
			"The DateTransferDto instance is corrupted because\n" +
			"'DateTransferDto.calendarBaseData' is INVALID!\n" +
			"dateTransDto.calendarBaseData == 'nil'\n")

		return isValid, err
	}

	isValid,
		err = dateTransDto.calendarBaseData.IsValidDate(
		dateTransDto.astronomicalYear,
		CalendarYearNumType(0).Astronomical(),
		dateTransDto.month,
		dateTransDto.day,
		ePrefix)

	if err != nil {
		isValid = false
		return isValid, err
	}

	return isValid, err
}
