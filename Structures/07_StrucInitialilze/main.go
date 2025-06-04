package main

import "fmt"

type NumSepsDto struct {
	DecimalSeparator   rune
	ThousandsSeparator rune
	CurrencySymbol     rune
}

type NumSepsProfile struct {
	NumSepSourceName  string
	UseDefaultNumSeps bool
	ValidateNumSeps   bool
	NumSeps           NumSepsDto
}

func main() {
	test01()
}

func test01() {

	nsProfile := NumSepsProfile{
		NumSepSourceName:  "intAry",
		UseDefaultNumSeps: false,
		ValidateNumSeps:   true,
		NumSeps:           NumSepsDto{},
	}

	myNumSeps := NumSepsDto{
		DecimalSeparator:   '.',
		ThousandsSeparator: ',',
		CurrencySymbol:     '$',
	}

	nsProfile.NumSeps = myNumSeps

	fmt.Printf("------------------\n")
	fmt.Printf("      test01      \n")
	fmt.Printf("------------------\n\n")

	fmt.Printf("myNumSeps    \n")
	fmt.Printf("%c\n", myNumSeps)
	fmt.Printf("\n\n")
	fmt.Printf("nsProfile    \n")
	fmt.Printf("%v\n", nsProfile)
	fmt.Printf("\n\nnsProfile.NumSeps\n%c\n", nsProfile.NumSeps)

	nsProfile.NumSeps.DecimalSeparator = 'X'
	nsProfile.NumSeps.ThousandsSeparator = 'Y'
	nsProfile.NumSeps.CurrencySymbol = 'Z'
	fmt.Printf("\n\n")
	fmt.Printf("After altering nsProfile.NumSeps    \n")
	fmt.Printf("nsProfile.NumSeps\n%c\n", nsProfile.NumSeps)
	fmt.Printf("Original myNumSeps\n%c\n", myNumSeps)

	nsProfile.NumSeps.DecimalSeparator = '.'
	nsProfile.NumSeps.ThousandsSeparator = ','
	nsProfile.NumSeps.CurrencySymbol = '$'

	myNumSeps.DecimalSeparator = 'A'
	myNumSeps.ThousandsSeparator = 'B'
	myNumSeps.CurrencySymbol = 'C'

	fmt.Printf("\n\n")
	fmt.Printf("After altering myNumSeps    \n")
	fmt.Printf("nsProfile.NumSeps\n%c\n", nsProfile.NumSeps)
	fmt.Printf("Original myNumSeps\n%c\n", myNumSeps)
}
