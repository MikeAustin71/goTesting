package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

type PrintOutDto struct {
	FactorLabels             []string
	FactorValues             []string
	FloatLabels              []string
	FloatOutputDecimalPlaces []int
	FloatValues              []*big.Float
	ExecutionTimeStr         string
}

func (prtOutDto PrintOutDto) New() PrintOutDto {

	newDto := PrintOutDto{}

	newDto.FactorLabels = make([]string, 0)
	newDto.FactorValues = make([]string, 0)
	newDto.FloatLabels = make([]string, 0)
	newDto.FloatOutputDecimalPlaces = make([]int, 0)
	newDto.FloatValues = make([]*big.Float, 0)
	newDto.ExecutionTimeStr = ""

	return newDto
}

func main() {
	TestTruncateBigFloat()
}

func TestRoundBigFloat() {

	var err error

	var newFloat *big.Float

	floatStr := "7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	ePrefix := "Test RoundBigFloat() "

	roundToDecPlaces := 20

	startTime := time.Now()

	result := RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime := time.Now()

	output := PrintOutDto{}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Rounded Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 0

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#3")

	floatStr = "-7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 0

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#4")

	floatStr = "7.44444444444444444444444444444444444444444444"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 3

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#5")

	floatStr = "-7.44444444444444444444444444444444444444444444"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 3

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#6")

	floatStr = "25.244"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 2

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#7")

	floatStr = "-25.244"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	roundToDecPlaces = 2

	startTime = time.Now()

	result = RoundBigFloat(newFloat, 1024, uint(roundToDecPlaces))

	endTime = time.Now()

	output.FactorLabels = []string{"Round To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", roundToDecPlaces)}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#8")

}

func TestCeilingBigFloat() {

	var err error

	var newFloat *big.Float

	floatStr := "7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	precision := uint(1024)

	ePrefix := "Test Ceiling() "

	startTime := time.Now()

	result := CeilingBigFloat(newFloat, precision)

	endTime := time.Now()

	output := PrintOutDto{}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Ceiling Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "7.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#3")

	floatStr = "-7.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#4")

	floatStr = "-2.70000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#5")

	floatStr = "2.30000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#6")

	floatStr = "-2.30000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = CeilingBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#7")

}

func TestFloorBigFloat() {

	var err error

	var newFloat *big.Float

	floatStr := "7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	ePrefix := "Test Floor() "

	precision := uint(1024)

	startTime := time.Now()

	result := FloorBigFloat(newFloat, precision)

	endTime := time.Now()

	output := PrintOutDto{}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Floor Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "7.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#3")

	floatStr = "-7.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#4")

	floatStr = "-2.70000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#5")

	floatStr = "2.30000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#6")

	floatStr = "-2.30000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = FloorBigFloat(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatValues = []*big.Float{newFloat, result}

	PrintBigFloatResults(output, ePrefix+"#7")

}

func TestBigFloatIntegerValue() {

	var err error

	var newFloat *big.Float

	floatStr := "79832.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	precision := uint(1024)

	ePrefix := "Test FloatIntegerValue() "

	startTime := time.Now()

	result := BigFloatIntegerValue(newFloat, precision)

	endTime := time.Now()

	output := PrintOutDto{}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Integer Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-79832.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = BigFloatIntegerValue(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Integer Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "79832.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = BigFloatIntegerValue(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Integer Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#3")

	floatStr = "0.00000000000000000000000000000000000000000000"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = BigFloatIntegerValue(newFloat, precision)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{}

	output.FloatLabels = []string{"Initial Value",
		"Integer Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#4")

}

func TestGetIntegerLength() {

	var err error

	var newFloat *big.Float

	floatStr := "79832.55555555555555555555555555555555555555555555"
	precision := uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	ePrefix := "Test GetIntegerLength() "

	startTime := time.Now()

	result, leadingSign := GetIntegerLengthFromBigFloat(newFloat)

	endTime := time.Now()

	output := PrintOutDto{}

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Expected Integer Length",
		"Actual Integer Length",
		"Leading Sign"}

	output.FactorValues = []string{"5",
		fmt.Sprintf("%v", result),
		leadingSign}

	output.FloatLabels = []string{"Initial Value"}

	output.FloatValues = []*big.Float{newFloat}

	output.FloatOutputDecimalPlaces = []int{50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "0.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result, leadingSign = GetIntegerLengthFromBigFloat(newFloat)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatLabels = []string{"Initial Value"}

	output.FactorValues = []string{"6",
		fmt.Sprintf("%v", result),
		leadingSign}

	output.FloatValues = []*big.Float{newFloat}

	output.FloatOutputDecimalPlaces = []int{50}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "-79832.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result, leadingSign = GetIntegerLengthFromBigFloat(newFloat)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatLabels = []string{"Initial Value"}

	output.FactorValues = []string{"6",
		fmt.Sprintf("%v", result),
		leadingSign}

	output.FloatValues = []*big.Float{newFloat}

	output.FloatOutputDecimalPlaces = []int{50}

	PrintBigFloatResults(output, ePrefix+"#3")

	floatStr = "0.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result, leadingSign = GetIntegerLengthFromBigFloat(newFloat)

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatLabels = []string{"Initial Value"}

	output.FactorValues = []string{"1",
		fmt.Sprintf("%v", result),
		leadingSign}

	output.FloatValues = []*big.Float{newFloat}

	output.FloatOutputDecimalPlaces = []int{50}

	PrintBigFloatResults(output, ePrefix+"#4")
}

func TestBigFloatFractionalValue() {

	ePrefix := "Test BigFloatFractionalValue() "

	var err error

	var newFloat *big.Float

	floatStr := "79832.55555555555555555555555555555555555555555555"
	precision := uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime := time.Now()

	result, numericSign := BigFloatFractionalValue(newFloat, precision)

	endTime := time.Now()

	output := PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Numeric Sign"}

	output.FactorValues = []string{
		fmt.Sprintf("%v", numericSign),
	}

	output.FloatLabels = []string{
		"Initial Value",
		"Fractional Value"}

	output.FloatValues = []*big.Float{
		newFloat,
		result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-79832.55555555555555555555555555555555555555555555"
	precision = uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result, numericSign = BigFloatFractionalValue(newFloat, precision)

	endTime = time.Now()

	output = PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Numeric Sign"}

	output.FactorValues = []string{
		fmt.Sprintf("%v", numericSign),
	}

	output.FloatLabels = []string{
		"Initial Value",
		"Fractional Value"}

	output.FloatValues = []*big.Float{
		newFloat,
		result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#2")

	floatStr = "0.00000000000000000000000000000000000000000000"
	precision = uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result, numericSign = BigFloatFractionalValue(newFloat, precision)

	endTime = time.Now()

	output = PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Numeric Sign"}

	output.FactorValues = []string{
		fmt.Sprintf("%v", numericSign),
	}

	output.FloatLabels = []string{
		"Initial Value",
		"Fractional Value"}

	output.FloatValues = []*big.Float{
		newFloat,
		result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#3")

}

func TestBigFloatIntFracValue() {

	ePrefix := "Test BigFloatIntFracValue() "

	var err error

	var newFloat *big.Float
	var floatStr string
	var precision uint

	floatStr = "79832.55555555555555555555555555555555555555555555"
	precision = uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime := time.Now()

	resultInt, resultFrac := BigFloatIntFracValue(newFloat, precision)

	endTime := time.Now()

	output := PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatLabels = []string{
		"Initial Value",
		"Integer Value",
		"Fractional Value"}

	output.FloatValues = []*big.Float{
		newFloat,
		resultInt,
		resultFrac}

	output.FloatOutputDecimalPlaces = []int{50, 50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	floatStr = "-79832.55555555555555555555555555555555555555555555"
	precision = uint(1024)

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	resultInt, resultFrac = BigFloatIntFracValue(newFloat, precision)

	endTime = time.Now()

	output = PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FloatLabels = []string{
		"Initial Value",
		"Integer Value",
		"Fractional Value"}

	output.FloatValues = []*big.Float{
		newFloat,
		resultInt,
		resultFrac}

	output.FloatOutputDecimalPlaces = []int{50, 50, 50}

	PrintBigFloatResults(output, ePrefix+"#2")

}

func TestTruncateBigFloat() {

	var err error

	var newFloat *big.Float

	floatStr := "7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	ePrefix := "Test TruncateBigFloat() "

	truncateToDecPlaces := 20

	startTime := time.Now()

	result := TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime := time.Now()

	output := PrintOutDto{}.New()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#1")

	truncateToDecPlaces = 20

	floatStr = "-7.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#2")

	truncateToDecPlaces = 1

	floatStr = "78.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#3")

	truncateToDecPlaces = 1

	floatStr = "-78.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#4")

	truncateToDecPlaces = 0

	floatStr = "78.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#5")

	truncateToDecPlaces = 0

	floatStr = "-78.55555555555555555555555555555555555555555555"

	newFloat,
		_,
		err =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			Parse(floatStr, 10)

	if err != nil {
		fmt.Printf("Error returned by big.Float Parse().\n"+
			"Error='%v'\n", err.Error())
		return
	}

	startTime = time.Now()

	result = TruncateBigFloat(newFloat, 1024, uint(truncateToDecPlaces))

	endTime = time.Now()

	output.ExecutionTimeStr = GetElapsedTime(endTime.Sub(startTime))

	output.FactorLabels = []string{"Truncate To Decimal Places"}

	output.FactorValues = []string{fmt.Sprintf("%v", truncateToDecPlaces)}

	output.FloatLabels = []string{"Initial Value",
		"Truncated Value"}

	output.FloatValues = []*big.Float{newFloat, result}

	output.FloatOutputDecimalPlaces = []int{50, 50}

	PrintBigFloatResults(output, ePrefix+"#6")

}

// BigFloatFractionalValue - Returns the absolute value of
// the fraction contained in a type *big.Float, floating point
// number. Again, the returned fractional number is always
// positive. In addition, an inter is returned signaling the
// numeric sign of the original floating point number,
// 'bigFloatNum'.
//
// If the input parameter, 'bigFloatNum' is equal to zero,
// the return parameter, 'numSign' is set to zero. If 'bigFloatNum'
// is less than zero, 'numSign' is set to -1. And, if 'bigFloatNum'
// is positive, the returned 'numSign' is set to +1.
func BigFloatFractionalValue(
	bigFloatNum *big.Float,
	precision uint) (floatFractionalValue *big.Float, numSign int) {

	numSign = 0

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatFractionalValue, numSign
	}

	var newBigFloatNum *big.Float

	if bigFloatNum.Sign() == -1 {
		numSign = -1
		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatNum)

	} else {
		// newBigFloatNum.Sign() == +1

		numSign = 1

		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(bigFloatNum)

	}

	bigIntVal, _ := newBigFloatNum.Int(nil)

	bigIntFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntVal)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Sub(newBigFloatNum, bigIntFloat)

	floatFractionalValue = floatFractionalValue.SetMode(floatFractionalValue.Mode())

	return floatFractionalValue, numSign
}

// BigFloatIntFracValue - Receives a floating point number
// of type *big.Float ('bigFloatNum') and returns the integer
// and fractional components as type *big.Flot floating point
// numbers.
//
// The numeric sign (plus + or minus -) of the original 'bigFloatNum'
// floating point number is always preserved in the returned integer
// component ('floatIntegerValue').  The fractional component
// ('floatFractionalValue') is always returned as a positive value.
func BigFloatIntFracValue(
	bigFloatNum *big.Float,
	precision uint) (
	floatIntegerValue,
	floatFractionalValue *big.Float) {

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatIntegerValue, floatFractionalValue
	}

	var numSign int

	var newBigFloatNum *big.Float

	if bigFloatNum.Sign() == -1 {
		numSign = -1
		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(bigFloatNum)

	} else {
		// newBigFloatNum.Sign() == +1

		numSign = 1

		newBigFloatNum =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(bigFloatNum)

	}

	bigIntVal, _ := newBigFloatNum.Int(nil)

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(bigIntVal)

	floatFractionalValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Sub(newBigFloatNum, floatIntegerValue)

	if numSign == -1 {

		floatIntegerValue =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Neg(floatIntegerValue)

	}

	floatIntegerValue = floatIntegerValue.SetMode(floatIntegerValue.Mode())

	floatFractionalValue = floatFractionalValue.SetMode(floatFractionalValue.Mode())

	return floatIntegerValue, floatFractionalValue
}

// BigFloatIntegerValue
// Returns the integer portion of a type *big.Float
// as a *big.Float.
func BigFloatIntegerValue(
	bigFloatNum *big.Float,
	precision uint) (floatIntegerValue *big.Float) {

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floatIntegerValue
	}

	var newBigFloat *big.Float

	newBigFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	if bigFloatNum.IsInt() {
		floatIntegerValue =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		return floatIntegerValue
	}

	var intValOfFloat *big.Int

	intValOfFloat, _ = newBigFloat.Int(nil)

	floatIntegerValue =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(intValOfFloat)

	floatIntegerValue = floatIntegerValue.SetMode(floatIntegerValue.Mode())

	return floatIntegerValue
}

func CeilingBigFloat(
	bigFloatNum *big.Float,
	precision uint) (ceiling *big.Float) {

	ceiling =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return ceiling
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	if newBigFloat.IsInt() {
		ceiling =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		ceiling = ceiling.SetMode(ceiling.Mode())

		return ceiling
	}

	var newInterimFloat *big.Float

	if newBigFloat.Sign() == -1 {
		// Input big float value is negative
		newInterimFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

	} else {
		// newBigFloat must be positive

		bigFloat1 :=
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(big.NewInt(1))

		newInterimFloat =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Add(newBigFloat, bigFloat1)
	}

	var intCeiling *big.Int

	intCeiling, _ = newInterimFloat.Int(nil)

	ceiling =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(intCeiling)

	ceiling = ceiling.SetMode(ceiling.Mode())

	return ceiling
}

func FloorBigFloat(
	bigFloatNum *big.Float,
	precision uint) (floor *big.Float) {

	floor =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return floor
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	if newBigFloat.IsInt() {

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				Set(newBigFloat)

		floor = floor.SetMode(floor.Mode())

		return floor
	}

	bigIntFloor, _ := newBigFloat.Int(nil)

	if newBigFloat.Sign() == 1 {
		// Numeric sign of  bigFloatNum is positive (+).

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(bigIntFloor)

	} else {
		// newBigFloat.Sign() == -1
		// bigFloatNum is LESS THAN zero

		bigIntFloor =
			big.NewInt(0).
				Add(bigIntFloor,
					big.NewInt(-1))

		floor =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(precision).
				SetInt(bigIntFloor)
	}

	floor = floor.SetMode(floor.Mode())

	return floor
}

func GetIntegerLengthFromBigFloat(
	bigFloatNum *big.Float) (intLength int, leadingSign string) {

	intLength = 0
	leadingSign = ""

	if bigFloatNum == nil {
		return intLength, leadingSign
	}

	var intValOfFloat *big.Int

	intValOfFloat, _ = bigFloatNum.Int(nil)

	intValStr := fmt.Sprintf("%v",
		intValOfFloat.Text(10))

	if intValStr[0] == '+' ||
		intValStr[0] == '-' {
		leadingSign = intValStr[0:1]
	}

	return len(intValStr), leadingSign
}

func GetIntValueOfBigFloat(
	bigFloatNum *big.Float) (intValOfFloat *big.Int) {

	intValOfFloat =
		big.NewInt(0)

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return intValOfFloat
	}

	intValOfFloat, _ = bigFloatNum.Int(nil)

	return intValOfFloat
}

func RoundBigFloat(
	bigFloatNum *big.Float,
	precision uint,
	roundToDecPlaces uint) (roundedFloat *big.Float) {

	if roundToDecPlaces > precision {
		precision = roundToDecPlaces + 100
	}

	roundedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	roundedFloat = roundedFloat.SetMode(roundedFloat.Mode())

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return roundedFloat
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	var bigInt5 *big.Int

	roundValue :=
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces+1)), nil)

	if newBigFloat.Sign() == -1 {

		bigInt5 = big.NewInt(-5)

	} else {
		bigInt5 = big.NewInt(5)
	}

	ratRound := big.NewRat(1, 1).
		SetFrac(bigInt5, roundValue)

	fracRound :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratRound)

	newNumFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Add(newBigFloat, fracRound)

	roundValue =
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(roundToDecPlaces)), nil)

	roundFrac :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(roundValue)

	newIntNumFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Mul(newNumFloat, roundFrac)

	newIntNum, _ :=
		newIntNumFloat.Int(nil)

	ratResult :=
		big.NewRat(1, 1).
			SetFrac(newIntNum, roundValue)

	roundedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratResult)

	roundedFloat = roundedFloat.SetMode(roundedFloat.Mode())

	return roundedFloat
}

func TruncateBigFloat(
	bigFloatNum *big.Float,
	precision uint,
	truncateToDecPlaces uint) (truncatedFloat *big.Float) {

	if truncateToDecPlaces > precision {
		precision = truncateToDecPlaces + 100
	}

	truncatedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetFloat64(0.0)

	truncatedFloat = truncatedFloat.SetMode(truncatedFloat.Mode())

	if bigFloatNum == nil ||
		bigFloatNum.Sign() == 0 {
		return truncatedFloat
	}

	newBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Set(bigFloatNum)

	truncateScale :=
		big.NewInt(0).
			Exp(big.NewInt(10),
				big.NewInt(int64(truncateToDecPlaces)), nil)

	truncateScaleFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetInt(truncateScale)

	newIntBigFloat :=
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			Mul(newBigFloat, truncateScaleFloat)

	newInt, _ := newIntBigFloat.Int(nil)

	ratResult :=
		big.NewRat(1, 1).
			SetFrac(newInt, truncateScale)

	truncatedFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(precision).
			SetRat(ratResult)

	truncatedFloat = truncatedFloat.SetMode(truncatedFloat.Mode())

	return truncatedFloat
}

func GetElapsedTime(elapsedTime time.Duration) string {

	nanosecondsElapsed := int64(elapsedTime)

	if nanosecondsElapsed == 0 {

		return fmt.Sprintf("Nanoseconds: %d\n",
			nanosecondsElapsed)

	} else if nanosecondsElapsed < 0 {
		return "ERROR!"
	}

	resultStr := ""

	temp := int64(time.Minute)
	tempResult := int64(0)

	if nanosecondsElapsed >= temp {
		tempResult = nanosecondsElapsed / temp

		resultStr += fmt.Sprintf("Minutes: %d   ",
			tempResult)

		nanosecondsElapsed -= tempResult * temp
	}

	temp = int64(time.Second)

	if nanosecondsElapsed >= temp {

		tempResult = nanosecondsElapsed / temp

		resultStr += fmt.Sprintf("Seconds: %d   ",
			tempResult)

		nanosecondsElapsed -= tempResult * temp

	}

	temp = int64(time.Millisecond)

	if nanosecondsElapsed >= temp {

		tempResult = nanosecondsElapsed / temp

		resultStr += fmt.Sprintf("Milliseconds: %d   ",
			tempResult)

		nanosecondsElapsed -= tempResult * temp

	} else {

		resultStr += "Milliseconds: 0   "

	}

	temp = int64(time.Microsecond)

	if nanosecondsElapsed >= temp {

		tempResult = nanosecondsElapsed / temp

		resultStr += fmt.Sprintf("Microseconds: %d   ",
			tempResult)

		nanosecondsElapsed -= tempResult * temp

	} else {

		resultStr += "Microseconds: 0   "

	}

	resultStr += fmt.Sprintf("Nanoseconds: %d\n",
		nanosecondsElapsed)

	return resultStr
}

func PrintBigFloatResults(
	output PrintOutDto,
	funcName string) {

	lineSepLength := 60

	lenFuncName := len(funcName)

	if lenFuncName > lineSepLength {
		lineSepLength = lenFuncName + 8
	}

	lineSeparator := strings.Repeat("-", lineSepLength)

	fmt.Println(lineSeparator)
	leftMargin := (lineSepLength - lenFuncName) / 2
	fmt.Println(strings.Repeat(" ", leftMargin) +
		funcName)
	fmt.Println(lineSeparator)

	if len(output.FactorLabels) > 0 {

		for i := 0; i < len(output.FactorLabels); i++ {
			fmt.Printf("%v: %v\n",
				output.FactorLabels[i],
				output.FactorValues[i])
		}

	}

	fmt.Println(lineSeparator)

	if len(output.FloatLabels) > 0 {

		labelFieldLen := 20

		outputFloatIntFieldLength := 0

		for h := 0; h < len(output.FloatLabels); h++ {

			if len(output.FloatLabels[h]) > labelFieldLen {
				labelFieldLen = len(output.FloatLabels[h]) + 2
			}

			valIntLen, _ := GetIntegerLengthFromBigFloat(output.FloatValues[h])
			if valIntLen > outputFloatIntFieldLength {
				outputFloatIntFieldLength = valIntLen
			}

		}

		spacer := strings.Repeat(" ", labelFieldLen+outputFloatIntFieldLength+3)
		fmt.Printf(spacer + "         1         2         3         4         5         6         7         8\n")
		fmt.Printf(spacer + "12345678901234567890123456789012345678901234567890123456789012345678901234567890\n")

		for i := 0; i < len(output.FloatLabels); i++ {

			label := output.FloatLabels[i]

			if len(label) < labelFieldLen {
				spacer := labelFieldLen - len(label)

				spacerStr := strings.Repeat(" ", spacer)

				// %45.40f
				fmtField := fmt.Sprintf("%d.%d",
					outputFloatIntFieldLength+1+output.FloatOutputDecimalPlaces[i],
					output.FloatOutputDecimalPlaces[i])

				label = spacerStr + label
				fmtStr := "%" + fmtField + "f"

				fmt.Printf(label+": "+fmtStr+"\n",
					output.FloatValues[i])
			}
		}
	}

	if len(output.ExecutionTimeStr) > 0 {
		fmt.Println(lineSeparator)
		fmt.Printf("Execution Time: %v\n",
			output.ExecutionTimeStr)
		fmt.Println(lineSeparator)
		fmt.Println()

	}

}
