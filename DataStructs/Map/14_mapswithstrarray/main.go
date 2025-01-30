package main

import "fmt"

type TestMapStrAry struct {
	MapStrArray    map[string][]string
}

func main() {
	tzAbbrv := "GMT+0000"
	tzCanonicalValue := "Africa/Abidjan"
	testMap := TestMapStrAry{}

	testMap.MapStrArray = make(map[string][]string)

	testMap.MapStrArray[tzAbbrv] = []string{tzCanonicalValue}

	tzAbbrv = "CST-0500"
	tzCanonicalValue = "America/Chicago"

	testMap.MapStrArray[tzAbbrv] = []string{tzCanonicalValue}

	for idx, canonicalValue := range testMap.MapStrArray {

		fmt.Printf("tzAbbrv='%v' -  tzCanonicalValue='%v'\n",
			idx, canonicalValue)
	}

	fmt.Println()
	fmt.Println("mapswtihstrarray - main()")
	fmt.Println("Successful Completion!")

}