package main

import (
	"fmt"
	"math/big"
	"strings"
	"time"
)

func main() {
	
		
	
	


	ePrefix := "BigFloat002.main() "

	separator := strings.Repeat("-", 75)

	var timeFraction, twentyFourHourNanosecondsFloat,
	convertedNanoseconds *big.Float

	var noonNanoseconds int64
	noonNanoseconds = int64(time.Hour * 12)


	var remainingNanoseconds int64
	
	var hour, minute, second, nanosecond,
		dayAdjustment int
	
	twentyFourHourNanosecondsFloat =
		big.NewFloat(0.0).
			SetMode(big.ToNearestAway).
			SetPrec(1024).
			SetInt64( int64(time.Hour * 24))


	for i:=0; i < 7; i++ {
		dayAdjustment = 0

		switch i {
		case 0:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.5)

		case 1:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.25)

		case 2:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.75)


		case 3:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.125)

		case 4:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.0)

		case 5:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.9)

		case 6:
			timeFraction =
				big.NewFloat(0.0).
					SetMode(big.ToNearestAway).
					SetPrec(1024).
					SetFloat64(0.6)

		default:
			fmt.Printf(ePrefix + "\n" +
				"Error: Iterator i INVALID!\n" +
				"i='%v'\n", i)
			return
		}

		convertedNanoseconds =
			big.NewFloat(0.0).
				SetMode(big.ToNearestAway).
				SetPrec(1024).
				Mul(timeFraction, twentyFourHourNanosecondsFloat)

		remainingNanoseconds, _ =
			convertedNanoseconds.Int64()

		if remainingNanoseconds < noonNanoseconds {
			remainingNanoseconds += noonNanoseconds
			dayAdjustment = 1
			} else {
			// remainingNanoseconds >= noonNanoseconds
			remainingNanoseconds -= noonNanoseconds
			dayAdjustment = 0
		}

		hour = int(remainingNanoseconds / int64(time.Hour))
		
		remainingNanoseconds -= int64(hour) * int64(time.Hour)
		
		minute = int(remainingNanoseconds / int64(time.Minute))
		
		remainingNanoseconds -= int64(minute) * int64(time.Minute)
		
		second = int(remainingNanoseconds / int64(time.Second))
		
		remainingNanoseconds -= int64(second) * int64(time.Second)
		
		nanosecond = int(remainingNanoseconds)
		
		
		fmt.Println(ePrefix)
		fmt.Println(separator)
		fmt.Printf("JDN Time Fraction: %17.8f\n",
			timeFraction)

		fmt.Printf("   Day Adjustment: %v\n",
			dayAdjustment)

		fmt.Printf("             Hour: %v\n",
			hour)

		fmt.Printf("           Minute: %v\n",
			minute)

		fmt.Printf("           Second: %v\n",
			second)

		fmt.Printf("       Nanosecond: %v\n",
			nanosecond)

		fmt.Println(separator)

	}


}
