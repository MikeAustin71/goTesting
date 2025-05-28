package main

import (
  "fmt"
  "sync"
)

// Courtesy of
// https://www.youtube.com/@anthonygg_
// https://www.youtube.com/watch?v=DBIPG-t9dfc

type SafeMap[K comparable, V any] struct {
  mu   sync.RWMutex // Automatically instantiated with zero value
  data map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {

  // sync.RWMutex Automatically instantiated with zero value
  return &SafeMap[K, V]{
    data: make(map[K]V),
  }
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
  m.mu.Lock()
  defer m.mu.Unlock()

  m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
  m.mu.RLock()
  defer m.mu.RUnlock()

  value, ok := m.data[key]

  if !ok {
    return value, fmt.Errorf("key %v Not Found", key)
  }

  return value, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
  m.mu.Lock()
  defer m.mu.Unlock()

  _, ok := m.data[key]

  if !ok {
    return fmt.Errorf("key %v Not Found", key)
  }

  m.data[key] = value

  return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
  m.mu.Lock()
  defer m.mu.Unlock()
  _, ok := m.data[key]
  if !ok {
    return fmt.Errorf("key %v Not Found", key)
  }

  delete(m.data, key)

  return nil
}

func (m *SafeMap[K, V]) Has(key K) bool {
  m.mu.RLock()
  defer m.mu.RUnlock()

  _, ok := m.data[key]

  return ok
}

func main() {

  myMap := New[int, int]()

  for i := 0; i < 10; i++ {
    go func(i int) {
      myMap.Insert(i, i*2)

      value, err := myMap.Get(i)

      if err != nil {
        fmt.Printf("Error in Get #%d - %v", i, err.Error())
        return
      }

      if value != i*2 {
        fmt.Printf("Error in Get #%d\n"+
          "Expected Value '%v'\n"+
          "Received Value '%v'\n\n", i, i*2, value)
        return
      }

    }(i)
  }
  fmt.Printf("Successful Completion!")
}
