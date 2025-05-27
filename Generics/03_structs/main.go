package main

import "fmt"

// Courtesy of:
// https://www.youtube.com/watch?v=BxQZmIEJ1Oo
// Golang Data Structures For Beginners: Generic List

type GenericList[T comparable] struct {
  data []T
}

func (l *GenericList[T]) Insert(value T) {
  l.data = append(l.data, value)
}

func (l *GenericList[T]) Get(i int) T {
  if i > len(l.data)-1 {
    panic("Index out of range")
  }

  for it := 0; it < len(l.data); it++ {
    if i == it {
      return l.data[it]
    }
  }
  panic("Value not found")
}

func (l *GenericList[T]) RemoveByValue(value T) {
  for i := 0; i < len(l.data); i++ {
    if l.data[i] == value {
      l.data = append(l.data[:i], l.data[i+1:]...)
    }
  }
}

func (l *GenericList[T]) Remove(i int) {
  if i > len(l.data)-1 {
    panic("Index out of range")
  }

  for it := 0; it < len(l.data); it++ {
    if it == i {
      l.data = append(l.data[:it], l.data[it+1:]...)
    }
  }
}

func New[T comparable]() *GenericList[T] {
  return &GenericList[T]{
    data: []T{},
  }
}

func main() {

  glist := New[string]()

  glist.Insert("bob")
  glist.Insert("foo")
  glist.Insert("bar")
  glist.Insert("alice")

  fmt.Printf("%+v\n\n", glist)

  glist.Remove(1)
  glist.RemoveByValue("alice")

  fmt.Printf("%+v\n\n", glist)

}
