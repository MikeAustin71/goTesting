package main

import (
	"fmt"
	"golangmikesamples/DateTimeTests/02_GregorianCalendar/GCal_Libs01"
	"math/big"
	"strings"
)

func main() {

	ePrefix := "GCal_001.main() "

	err := testGregorianDateToJDN(ePrefix)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

}

func testJulianDayToGregorianDateTime(
	ePrefix string) (err error) {

	ePrefix += "testJulianDayToGregorianDateTime() "
	err = nil

	expectedDateTimeStr := "2020-08-03 00:44:38.000000000 UTC  = 2459064.530995370370370370370370370370"
	//  2020-08-03 00:44:38.000000000 UTC  = 2459064.530995370370370370370370370370

	var julianDayNoTime *big.Float
	var b int
	var err2 error

	julianDayNoTime,
		b,
		err2 =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse("2459064.530995370370370370370370370370", 10)

	if err2 != nil {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error returned by big.Float.Parse()\n"+
			"Error='%v'\n", err2.Error())
		return err
	}

	if b != 10 {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Value of base, 'b' returned by big.Float.Parse()\n"+
			"is incorrect! b=%v\n", b)
		return err
	}

	var expectedDateTime GCal_Libs01.DateTimeTransferDto

	expectedDateTime, err = GCal_Libs01.DateTimeTransferDto{}.New(
		false,
		2020,
		8,
		3,
		0,
		0,
		44,
		38,
		ePrefix)

	if err != nil {
		return err
	}

	// -4714-11-24 12:00:00.000000000 UTC = JDN -365.0
	// -4714-11-24 00:00:00.000000000 UTC = JDN -364.5
	// -4713-11-25 12:00:00.000000000 UTC = JDN -1.0
	// -4713-11-25 00:00:00.000000000 UTC = JDN -0.5
	// -4400-11-24 12:00:00.000000000 UTC = JDN 114321.0
	//  0001-01-01 12:00:00.000000000 UTC = JDN 1721426.0
	//  0001-01-01 00:00:00.000000000 UTC = JDN 1721425.5
	//  0001-01-01 11:00:00.000000000 UTC = JDN 1721425.95833
	//  2000-01-01 12:00:00.000000000 UTC  = 2,451,545.0000000000000000
	//  2001-01-01 12:00:00.000000000 UTC  = 2,451,911.0000000000000000
	//  2002-01-01 12:00:00.000000000 UTC  = 2,452,276.0000000000000000
	//  2003-01-01 12:00:00.000000000 UTC  = 2,452,641.0000000000000000
	//  2004-01-01 12:00:00.000000000 UTC  = 2,453,006.0000000000000000
	//  2005-01-01 12:00:00.000000000 UTC  = 2,453,372.0000000000000000
	//  2006-01-01 12:00:00.000000000 UTC  = 2,453,737.0000000000000000
	//  2007-01-01 12:00:00.000000000 UTC  = 2,454,102.0000000000000000
	//  2008-01-01 12:00:00.000000000 UTC  = 2,454,467.0000000000000000
	//  2009-01-01 12:00:00.000000000 UTC  = 2,454,833.0000000000000000
	//  2010-01-01 12:00:00.000000000 UTC  = 2,455,198.0000000000000000
	//  2011-01-01 12:00:00.000000000 UTC  = 2,455,563.0000000000000000
	//  2012-01-01 12:00:00.000000000 UTC  = 2,455,928.0000000000000000
	//  2013-01-01 12:00:00.000000000 UTC  = 2,456,294.0000000000000000
	//  2014-01-01 12:00:00.000000000 UTC  = 2,456,659.0000000000000000
	//  2015-01-01 12:00:00.000000000 UTC  = 2,457,024.0000000000000000
	//  2016-01-01 12:00:00.000000000 UTC  = 2,457,389.0000000000000000
	//  2020-08-03 00:44:38.000000000 UTC  = 2,459,064.5310000000000000
	//                                       2459064.530995370370370370370370370370
	expectedDateTime.SetTag("Expected Date Time")

	calGreg := GCal_Libs01.CalendarGregorianUtility{}

	var actualDateTime GCal_Libs01.DateTimeTransferDto

	actualDateTime,
		err =
		calGreg.JulianDayNumberTimeToDateTime(
			julianDayNoTime,
			ePrefix)

	if err != nil {
		return err
	}

	actualDateTime.SetTag("Actual Date Time")

	separator := strings.Repeat("*", 75)

	fmt.Println(separator)
	fmt.Printf("Expected Date Time: %v\n",
		expectedDateTimeStr)
	fmt.Printf("Julian Day No Time: %42.30f\n",
		julianDayNoTime)

	fmt.Println(separator)
	fmt.Printf("%v", expectedDateTime.String())
	fmt.Printf("%v", actualDateTime.String())
	fmt.Println(separator)

	var compareResult int

	compareResult, err = expectedDateTime.Compare(&actualDateTime, ePrefix)

	if err != nil {
		return err
	}

	if compareResult == 0 {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!! SUCCESS !!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("     Expected Date/Time MATCHES Actual Date Time")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	} else {
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@ FAILURE @@@@@@@@@@@@@@@@@@@@@@@")
		fmt.Println("  Expected Date/Time DOES NOT MATCH Actual Date Time  ")
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")

	}

	return err
}

func testGregorianDateToJDN(
	ePrefix string) (err error) {

	ePrefix += "testGregorianDateToJDN() "
	err = nil

	// 0001-01-01 11:00:00.000000000 UTC = JDN 1721425.95833

	targetYear := int64(0001)
	targetMonth := 1
	targetDay := 1
	targetHour := 11
	targetMinute := 0
	targetSecond := 0
	targetNanosecond := 0

	var expectedTotalDayNoTime *big.Float

	expectedTotalDayNoTime,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse("1721425.958330000000", 10)

	// -4714-11-24 12:00:00.000000000 UTC = JDN -365.0
	// -4714-11-24 00:00:00.000000000 UTC = JDN -364.5
	// -4713-11-25 12:00:00.000000000 UTC = JDN -1.0
	// -4713-11-25 00:00:00.000000000 UTC = JDN -0.5
	// -4400-11-24 12:00:00.000000000 UTC = JDN 114321.0
	//  0001-01-01 12:00:00.000000000 UTC = JDN 1721426.0
	//  0001-01-01 00:00:00.000000000 UTC = JDN 1721425.5
	//  0001-01-01 11:00:00.000000000 UTC = JDN 1721425.95833
	//  1859-08-06 05:30:00.000000000 UTC = JDN 2400262.72917
	//  2000-01-01 12:00:00.000000000 UTC  = 2,451,545.0000000000000000
	//  2001-01-01 12:00:00.000000000 UTC  = 2,451,911.0000000000000000
	//  2002-01-01 12:00:00.000000000 UTC  = 2,452,276.0000000000000000
	//  2003-01-01 12:00:00.000000000 UTC  = 2,452,641.0000000000000000
	//  2004-01-01 12:00:00.000000000 UTC  = 2,453,006.0000000000000000
	//  2005-01-01 12:00:00.000000000 UTC  = 2,453,372.0000000000000000
	//  2006-01-01 12:00:00.000000000 UTC  = 2,453,737.0000000000000000
	//  2007-01-01 12:00:00.000000000 UTC  = 2,454,102.0000000000000000
	//  2008-01-01 12:00:00.000000000 UTC  = 2,454,467.0000000000000000
	//  2009-01-01 12:00:00.000000000 UTC  = 2,454,833.0000000000000000
	//  2010-01-01 12:00:00.000000000 UTC  = 2,455,198.0000000000000000
	//  2011-01-01 12:00:00.000000000 UTC  = 2,455,563.0000000000000000
	//  2012-01-01 12:00:00.000000000 UTC  = 2,455,928.0000000000000000
	//  2013-01-01 12:00:00.000000000 UTC  = 2,456,294.0000000000000000
	//  2014-01-01 12:00:00.000000000 UTC  = 2,456,659.0000000000000000
	//  2015-01-01 12:00:00.000000000 UTC  = 2,457,024.0000000000000000
	//  2016-01-01 12:00:00.000000000 UTC  = 2,457,389.0000000000000000
	//  2020-08-03 00:44:38.000000000 UTC  = 2,459,064.5310000000000000

	calGreg := GCal_Libs01.CalendarGregorianUtility{}

	var julianDayNumber int64
	var julianDayNumberTime, julianDayNumberTimeFraction *big.Float

	julianDayNumber,
		julianDayNumberTime,
		julianDayNumberTimeFraction,
		err = calGreg.GetJDN(
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

	expectedJDNoStr := fmt.Sprintf("%45.30f",
		expectedTotalDayNoTime)

	actualJDNoTimeStr := fmt.Sprintf("%45.30f",
		julianDayNumberTime)

	separator := strings.Repeat("&", 75)
	fmt.Println(separator)
	fmt.Printf("     %v\n", ePrefix)
	fmt.Println(separator)
	fmt.Printf("Expected Julian Day Number: %v\n",
		expectedJDNoStr)
	fmt.Println(separator)
	fmt.Printf("         Julian Day Number: %14d\n",
		julianDayNumber)

	fmt.Printf("     Julian DayNumber Time: %v\n",
		actualJDNoTimeStr)

	fmt.Printf("Julian DayNumber Fraction: %46.30f\n",
		julianDayNumberTimeFraction)
	fmt.Println(separator)
	if expectedJDNoStr == actualJDNoTimeStr {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! SUCCESS !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("Expected Julian Day Number Time MATCHES Actual Julian Day Number Time!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	} else {
		fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%% FAILURE %%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
		fmt.Println("Expected Julian Day Number Time DOES NOT MATCH Actual Julian Day Number Time!")
		fmt.Println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
	}

	fmt.Println(separator)

	return err

}
