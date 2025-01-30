package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack("France")
	l.PushBack("Russia")
	l.PushBack("Great Britain")
	l.PushBack("Italy")
	l.PushBack("Germany")
	l.PushBack("Poland")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

/*	Output
	$ go run main.go
	France
	Russia
	Great Britain
	Italy
	Germany
	Poland
*/
