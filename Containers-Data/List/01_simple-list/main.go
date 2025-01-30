package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Create new list and stuff some numbers into it
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate over list
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

}

/*	Output
	$ go run main.go
	1
	2
	3
	4
*/
