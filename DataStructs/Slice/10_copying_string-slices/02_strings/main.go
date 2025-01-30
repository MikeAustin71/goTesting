package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {

	s := "How now brown cow. Where is my cow"

	targetSubStr := "cow"
	replacementStr := "cat"
	testStrReplace(s, targetSubStr, replacementStr, 1)
}

func test1(tStr string) {

	fmt.Printf("Original String: %v",
		tStr)
	fmt.Println()

	newRunes := make([]rune, 0, 50)

	for _, r := range tStr {

		if r == 'x' {
			continue
		}
		newRunes = append(newRunes, r)
	}

	fmt.Printf("Final String: %v",
		string(newRunes))
	fmt.Println()
}

func testStringRemove(
	universeStr string,
	targetStr string,
	maxRemovalCount int) {

	if maxRemovalCount < 1 {
		maxRemovalCount = math.MaxInt32
	}

	fmt.Printf(" Original String: %v",
		universeStr)
	fmt.Println()

	tStr2 := universeStr

	tStrResult := strings.Replace(tStr2, targetStr, "", maxRemovalCount)

	fmt.Printf(" strings.Replace: \"%v\"",
		tStrResult)
	fmt.Println()

	newRunes := make([]rune, 0, 100)

	targetRunes := []rune(targetStr)
	lenTargetRunes := len(targetRunes)
	universeRunes := []rune(universeStr)
	lenUniverseRunes := len(universeStr)

	isCorrectCharCnt := 0
	removalCnt := 0

	for i := 0; i < lenUniverseRunes; i++ {

		if universeRunes[i] == targetRunes[0] &&
			i+lenTargetRunes <= lenUniverseRunes &&
			removalCnt < maxRemovalCount {

			isCorrectCharCnt = 1

			for j := 1; j < lenTargetRunes; j++ {

				if universeRunes[i+j] == targetRunes[j] {
					isCorrectCharCnt++
				} else {
					break
				}
			}

			if isCorrectCharCnt == lenTargetRunes {
				i += lenTargetRunes - 1
				removalCnt++
			} else {
				newRunes = append(newRunes, universeRunes[i])
			}

			continue

		} else {
			newRunes = append(newRunes, universeRunes[i])
		}
	}

	fmt.Printf("Processed String: \"%v\"",
		string(newRunes))
	fmt.Println()

}

func testStrReplace(
	universeStr string,
	targetSubStr string,
	replacementStr string,
	maxReplacementCount int) {

	if maxReplacementCount < 1 {
		maxReplacementCount = math.MaxInt32
	}

	fmt.Printf(" Original String: %v",
		universeStr)
	fmt.Println()

	tStr2 := universeStr

	timeStart := time.Now()

	tStrResult := strings.Replace(tStr2, targetSubStr, replacementStr, maxReplacementCount)

	timeEnd := time.Now()

	timeDuration := timeEnd.Sub(timeStart)
	stringsReplaceTime := GetElapsedTime(timeDuration)

	fmt.Printf(" strings.Replace: %v",
		tStrResult)
	fmt.Println()

	timeStart = time.Now()

	newRunes := make([]rune, 0, 100)

	targetRunes := []rune(targetSubStr)
	replacementRunes := []rune(replacementStr)
	lenTargetRunes := len(targetRunes)
	universeRunes := []rune(universeStr)
	lenUniverseRunes := len(universeStr)

	isCorrectCharCnt := 0
	replacementCnt := 0

	for i := 0; i < lenUniverseRunes; i++ {

		if universeRunes[i] == targetRunes[0] &&
			i+lenTargetRunes <= lenUniverseRunes &&
			replacementCnt < maxReplacementCount {

			isCorrectCharCnt = 1

			for j := 1; j < lenTargetRunes; j++ {

				if universeRunes[i+j] == targetRunes[j] {
					isCorrectCharCnt++
				} else {
					break
				}
			}

			if isCorrectCharCnt == lenTargetRunes {
				i += lenTargetRunes - 1
				replacementCnt++
				// copy(newRunes[i:], replacementRunes)
				newRunes = append(newRunes, replacementRunes...)

			} else {
				newRunes = append(newRunes, universeRunes[i])
			}

			continue

		} else {
			newRunes = append(newRunes, universeRunes[i])
		}
	}

	timeEnd = time.Now()

	customTime := GetElapsedTime(timeEnd.Sub(timeStart))

	fmt.Printf("Processed String: %v",
		string(newRunes))
	fmt.Println()

	fmt.Printf("Strings Replace Elapsed Time: %v\n",
		stringsReplaceTime)

	fmt.Printf("  Mike's Custom Elapsed Time: %v\n",
		customTime)

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
