package main

import (
	"fmt"
	"regexp"
)

/*
Tebeka Shorts: Define a Regular Expression
Function

https://www.youtube.com/watch?v=NvpPcfBGUWw
https://pkg.go.dev/regexp#MustCompile
https://pkg.go.dev/regexp#Compile
https://pkg.go.dev/regexp#Regexp.ReplaceAllStringFunc
*/

func main() {

	urlTemplate := `https://$HOST:$PORT`

	conf := map[string]string{
		"HOST": "www.ardanlabs.com",
		"PORT": "443",
	}

	re := regexp.MustCompile(`\$[A-Z_]+`)

	sub := func(match string) string {
		key := match[1:] // Remove $ prefix
		return conf[key]
	}

	url := re.ReplaceAllStringFunc(urlTemplate, sub)

	fmt.Println(url)
	// https://www.ardanlabs.com:443
}
