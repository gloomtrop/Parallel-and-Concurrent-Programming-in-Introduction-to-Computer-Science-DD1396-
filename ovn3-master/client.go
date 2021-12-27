package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		before := time.Now()
		// res := Get(server[0], client)
		res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(100 * time.Millisecond)
	}
}

type Response struct {
	Body       string
	StatusCode int
}

type Atomic struct {
	mu  sync.Mutex
	res *Response
}

func (r *Response) String() string {
	return fmt.Sprintf("%q (%d)", r.Body, r.StatusCode)
}

func (r *Atomic) Add(res *Response) {
	r.mu.Lock() // Wait for the lock to be free and then take it.
	r.res = res
	r.mu.Unlock() // Release the lock.
}

// Get makes an HTTP Get request and returns an abbreviated response.
// The response is empty if the request fails.
func Get(url string, client *http.Client) *Response {
	res, err := client.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), res.StatusCode}
}

// MultiGet makes an HTTP Get request to each url and returns
// the response from the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is 503
// – Service unavailable.
func MultiGet(urls []string, client *http.Client) *Response {
	var response Atomic
	// before := time.Now()
	wg := new(sync.WaitGroup)
	wg.Add(1)

	for i := 0; i < len(urls); i++ {

		var res *Response
		go func(url string) {
			// fmt.Println(url)

			res = Get(url, client)
			// fmt.Println(res)
			if res.StatusCode == 200 && response == (Atomic{}) {
				// fmt.Println(res.StatusCode)
				response.Add(res)
				wg.Done()
			}
		}(urls[i])
	}

	if waitTimeout(wg, time.Second) {
		return &Response{"Service Unavailable", 503}
	} else {
		return response.res
	}

}

//https://stackoverflow.com/questions/32840687/timeout-for-waitgroup-wait
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
