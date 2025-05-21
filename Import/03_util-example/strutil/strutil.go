package strutil

import "fmt"

/*
	This example ranges over a string returning both
	the index and the value.  The index is the index
	of each rune in the string. The value is the rune
	itself.
*/

type StrUtil struct{}

func (sUtil *StrUtil) RangeOverChars(str string) (result string) {

	if str == "" {
		result = "Hello World!"
	} else {
		result = str
	}

	fmt.Printf("Input String= %s\n\n", result)

	for i, v := range result {
		fmt.Printf("index=%d unicode value=%U character value=%q\n", i, v, v)
	}

	return result
}

/* Output
$ go run main.go
index=0 unicode value=U+0048 character value='H'
index=1 unicode value=U+0065 character value='e'
index=2 unicode value=U+006C character value='l'
index=3 unicode value=U+006C character value='l'
index=4 unicode value=U+006F character value='o'
index=5 unicode value=U+0020 character value=' '
index=6 unicode value=U+0057 character value='W'
index=7 unicode value=U+006F character value='o'
index=8 unicode value=U+0072 character value='r'
index=9 unicode value=U+006C character value='l'
index=10 unicode value=U+0064 character value='d'
index=11 unicode value=U+0021 character value='!'
*/
