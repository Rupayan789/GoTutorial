package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var counter int = 0
var wg *sync.WaitGroup = &sync.WaitGroup{}
var mutex *sync.Mutex = &sync.Mutex{}

func main() {
	var endpoints []string = []string{
		"https://amazon.in",
		"https://yahoo.com",
		"https://github.com",
		"https://go.dev",
	}
	before := time.Now().UnixNano()
	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Printf("Counter :%v\n", counter)
	after := time.Now().UnixNano()
	defer getDuration(before, after)

}

func getDuration(start int64, end int64) {
	gap := float64(end - start)

	fmt.Println(gap)

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	// Without any lock mechanism this is giving undesired result
	mutex.Lock()
	fmt.Printf("Counter %v %v\n", counter, endpoint)
	if counter == 0 {
		time.Sleep(1 * time.Second)
		counter++
	} else {
		counter--
	}
	mutex.Unlock()
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)

}
