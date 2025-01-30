package main

import (
	"fmt"
)

func main()  {
	nStr := "1234567890"
	numInts := 6
	//numFracs := 4
	absAllNumRunes := []rune(nStr)


	fmt.Println("absAllNumRunes: ", string(absAllNumRunes))
	intRunes := absAllNumRunes[0:numInts]
	fmt.Println("      intRunes: ", string(intRunes))
	fracRunes := absAllNumRunes[numInts:]
	fmt.Println("     fracRunes: ", string(fracRunes))


}

/* Output
absAllNumRunes:  1234567890
      intRunes:  123456
     fracRunes:  7890
 */

