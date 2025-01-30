package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now().In(time.Local)

	tStr := t.Format("2006-01-02 15:04:05 -0700 MST")

	lenLeadStr := len("2006-01-02 15:04:05 ")

	offsetStr := tStr[lenLeadStr: lenLeadStr + 5]

	fmt.Println("  Time String: ", tStr)
	fmt.Println("       Offset: ", offsetStr)
	fmt.Println("Length Offset: ", len(offsetStr))

}
