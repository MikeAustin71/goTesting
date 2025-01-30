package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start Time", start)
	time.Sleep(6 * time.Second)
	elapsedTime := time.Since(start)

	fmt.Println("Elapsed Time", elapsedTime)
}

