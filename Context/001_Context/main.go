package main

import (
  "context"
  "fmt"
  "log"
  "time"
)

// Thanks to Anthony CG
//
// How To Use The Context Package In Golang?
// https://www.youtube.com/watch?v=kaZOXRqFPCw

func main() {
  start := time.Now()

  //ctx := context.Background()
  ctx := context.WithValue(context.Background(), "foo", "bar")

  userID := 10
  val, err := fetchUserData(ctx, userID)
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println("result: ", val)
  fmt.Println("took: ", time.Since(start))
}

type Response struct {
  value int
  err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {

  val := ctx.Value("foo")

  fmt.Println(val.(string))

  ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)

  defer cancel()

  respchan := make(chan Response)

  go func() {

    val, err := fetchThirdPartyStuffWhichCanBeSlow()

    respchan <- Response{
      value: val,
      err:   err,
    }

  }()

  for {

    select {
    case <-ctx.Done():
      return 0, fmt.Errorf("fetching data from 3rd party took too long")

    case resp := <-respchan:
      return resp.value, resp.err
    }
  }
}

func fetchThirdPartyStuffWhichCanBeSlow() (int, error) {
  time.Sleep(500 * time.Millisecond)

  return 666, nil
}
