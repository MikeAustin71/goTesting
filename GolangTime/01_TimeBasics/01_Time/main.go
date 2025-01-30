package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Start Time", start)
	time.Sleep(3 * time.Second)
	end := time.Now()
	fmt.Println("End Time", end)
}
