/* 	Code taken from Nathan LeClaire's Blog Post
"How to Wait for All Goroutines to Finish Executing Before Continuing".
https://nathanleclaire.com/blog/2014/02/15/how-to-wait-for-all-goroutines-to-finish-executing-before-continuing/

This code "slurps JSON data from 3 different subreddits concurrently. We don’t know how long those HTTP requests are going to take, and we don’t want to cause a race condition by trying to work with data that hasn’t been populated yet in our Go program, so sync.WaitGroup ends up being very handy:"

Note - You may Get Error 429 - To many requests
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://www.reddit.com/r/aww.json",
		"http://www.reddit.com/r/funny.json",
		"http://www.reddit.com/r/programming.json",
	}
	jsonResponses := make(chan string)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					log.Fatal(err)
				} else {
					jsonResponses <- string(body)
				}
			}
		}(url)
	}

	go func() {
		for response := range jsonResponses {
			fmt.Println(response)
		}
	}()

	wg.Wait()
}

/* 	Ouput = A lot of JSON content from the three url's
	Be advised you may Get Error 429 - To many requests
*/
