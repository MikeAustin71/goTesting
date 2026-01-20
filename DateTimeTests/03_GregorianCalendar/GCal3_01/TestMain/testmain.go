package TestMain

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"

	"github.com/mikeaustin71/DateTimeTests/03_GregorianCalendar/datetime"
)

type TestMain struct {
	input  string
	output string
}

func (testMain *TestMain) TestTypeAssertion(
	ePrefix string) (err error) {

	ePrefix += "TestMain.TestTypeAssertion() "

	err = nil

	var dayOfWeekNo datetime.IDayOfWeekNumber

	dayOfWeekNo = datetime.UsWeekDayNo.Saturday()

	var ok bool

	_, ok = dayOfWeekNo.(datetime.ISO8601DayOfWeekNo)

	if ok == true {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Expected dayOfWeekNo.(datetime.ISO8601DayOfWeekNo)=false !\n"+
			"Instead, dayOfWeekNo.(datetime.ISO8601DayOfWeekNo)=true.\n"+
			"dayOfWeekNo='%v'\n",
			dayOfWeekNo.XDayOfWeekNumberingSystemType().String())
		return err
	}

	_, ok = dayOfWeekNo.(datetime.UsDayOfWeekNo)

	if ok == false {
		err = fmt.Errorf(ePrefix+"\n"+
			"Error: Expected dayOfWeekNo.(datetime.UsDayOfWeekNo)=true !\n"+
			"Instead, dayOfWeekNo.(datetime.UsDayOfWeekNo)=false.\n"+
			"dayOfWeekNo='%v'\n",
			dayOfWeekNo.XDayOfWeekNumberingSystemType().String())
		return err
	}

	var usDayOfWeek datetime.UsDayOfWeekNo

	usDayOfWeek = dayOfWeekNo.(datetime.UsDayOfWeekNo)

	lineLength := 77
	txtStr := "Testing Type Assertion"
	testUtil := testUtilities{}

	separator1 := strings.Repeat("=", lineLength)
	fmt.Println(separator1)
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, ePrefix))
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, txtStr))
	fmt.Println(separator1)
	spacer := strings.Repeat(" ", 5)
	txtStr = "usDayOfWeek.String()= " + usDayOfWeek.String()
	fmt.Println(spacer + txtStr)
	txtStr = "usDayOfWeek.XDayOfWeekName()= " + usDayOfWeek.XDayOfWeekName()
	fmt.Println(spacer + txtStr)
	txtStr = "reflect.TypeOf(usDayOfWeek).String())= " +
		reflect.TypeOf(usDayOfWeek).String()

	fmt.Println(spacer + txtStr)
	fmt.Println(separator1)
	txtStr = "*** SUCCESS! *** "
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, txtStr))
	fmt.Println(separator1)
	fmt.Println()
	fmt.Println()

	return err
}

func (testMain *TestMain) TestEnumSwitch(
	ePrefix string) error {

	LineLength := 77
	ePrefix += "TestMain.testEnumSwitch()"
	txtStr := "Testing Enumerator in Switch Statement"

	spacer := strings.Repeat(" ", (LineLength-len(ePrefix))/2)

	yearNumType := datetime.CalendarYearNumType(61).Astronomical()

	separator1 := strings.Repeat("=", LineLength)

	fmt.Println(separator1)
	spacer = strings.Repeat(" ", (LineLength-len(txtStr))/2)
	fmt.Println(spacer + ePrefix)
	fmt.Println(spacer + txtStr)
	fmt.Println(separator1)

	switch yearNumType {

	case datetime.CalYearType.CE():
		fmt.Println("Found CalYearType.CE()")
	case datetime.CalYearType.BCE():
		fmt.Println("Found datetime.CalYearType.BCE()")
	case datetime.CalYearType.Astronomical():
		fmt.Println("Found CalYearType.Astronomical()")
	case datetime.CalYearType.None():
		fmt.Println("Found datetime.CalYearType.None()")
	default:
		fmt.Println("default: Switch Failed and Found Nothing!")
	}
	fmt.Println()
	fmt.Println(separator1)

	return nil
}

func (testMain *TestMain) TestYearMonthDayFromOrdinalDayNo(
	ePrefix string) (
	err error) {

	ePrefix += "TestYearMonthDayFromOrdinalDayNo() "

	err = nil
	lineLength := 77
	separator1 := strings.Repeat("=", lineLength)
	separator2 := strings.Repeat("-", lineLength)
	separator3 := strings.Repeat("*", lineLength)
	separator4 := strings.Repeat("%", lineLength)

	year := int64(4715)
	yearNumType := datetime.CalYearType.BCE()

	ordinalDate := 197
	expectedYear := int64(-4716)
	expectedMonth := 7
	expectedDay := 15

	var astronomicalYear int64
	var month, day int

	gregCalBData := datetime.CalendarGregorianBaseData{}

	astronomicalYear,
		month,
		day,
		err = gregCalBData.GetYearMonthDayFromOrdinalDayNo(
		ordinalDate,
		year,
		yearNumType,
		ePrefix)

	if err != nil {
		return err
	}

	expectedDate := fmt.Sprintf("%04d-%02d-%02d",
		expectedYear, expectedMonth, expectedDay)

	actualDate := fmt.Sprintf("%04d-%02d-%02d",
		astronomicalYear, month, day)

	testUtil := testUtilities{}

	fmt.Println(separator1)
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, ePrefix))
	fmt.Println(separator2)
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, "** Input Data **"))
	fmt.Println(separator2)
	fmt.Printf("     Ordinal Date: %v\n", ordinalDate)
	fmt.Printf("             year: %v\n", year)
	fmt.Printf("      yearNumType: %v\n", yearNumType.String())
	fmt.Println(separator3)
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, "** Results **"))
	fmt.Println(separator3)
	fmt.Printf("Astronomical Year: %v\n", astronomicalYear)
	fmt.Printf("            month: %v\n", month)
	fmt.Printf("              day: %v\n", day)
	fmt.Println(separator4)
	fmt.Println(testUtil.centerStrLeftSpace(lineLength, "** Results Analysis **"))
	fmt.Printf("         Expected: %v\n", expectedDate)
	fmt.Printf("           Actual: %v\n", actualDate)
	fmt.Println(separator4)
	if expectedDate == actualDate {
		fmt.Println(testUtil.centerStrLeftSpace(lineLength, "'Expected' MATCHES 'Actual'"))
		fmt.Println(testUtil.centerStrLeftSpace(lineLength, "!!!! SUCCESS !!!!"))
	} else {
		fmt.Println(testUtil.centerStrLeftSpace(lineLength, "'Expected' DOES NOT MATCH 'Actual'."))
		fmt.Println(testUtil.centerStrLeftSpace(lineLength, "&&&&&&&& FAILURE  &&&&&&&&"))
	}
	fmt.Println(separator1)

	return err
}

/*func (testMain *TestMain) TestBigFloatRounding(
	ePrefix string) (err error) {

	ePrefix += "TestYearMonthDayFromOrdinalDayNo() "

	precision := uint(500)

	floatFmtStr := "%46.30f"


}
*/

type testUtilities struct {
	lock *sync.Mutex
}

func (testUtils *testUtilities) centerStrLeftSpace(
	lineLength int,
	textString string) (finalString string) {

	lenTxtStr := len(textString)

	if lenTxtStr >= lineLength {
		return textString
	}

	var dividend = float64(lineLength - lenTxtStr)

	var quotient float64

	quotient = dividend / 2.0

	quotient += 0.5

	spacerLen := int(math.Floor(quotient))

	spacer := strings.Repeat(" ", spacerLen)

	finalString = spacer + textString

	return finalString
}
