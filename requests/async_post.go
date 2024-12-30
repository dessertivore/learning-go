// Call /addition endpoint 3 times with different values
// and print the result
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type APIOutput struct {
		MainOutput int `json:"sum" doc:"Sum of numbers inputted"`
	}

type AdditionInput struct {
		NumsToAdd []int `json:"numsToAdd"`
	}

func postRequest(nums []int) error {
	start := time.Now()
	// HTTP endpoint
	postURL := "http://127.0.0.1:8888/addition"

	// JSON body
	input := AdditionInput{NumsToAdd: nums}
	jsonData, err := json.Marshal(input)
    if err != nil {
        return fmt.Errorf("error marshaling input: %v", err)
    }
	fmt.Println("Sending:", string(jsonData))

	// Create a HTTP post request
	r, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	post := &APIOutput{}
	derr := json.NewDecoder(res.Body).Decode(post)
	if derr != nil {
		return derr
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status:", res.StatusCode)
		return fmt.Errorf("status: %s", res.Status)
	}
	full := time.Since(start)
	fmt.Println("Sum:", post.MainOutput)
	fmt.Println("Time taken:", full)
	return nil
}

func main () {
	// Call API twice using a goroutine alongisde a synchronous func call
	// The main goroutine does not wait for the new goroutine to complete, it only waits
	// for the synchronous (blocking) function to complete before continuing to next
	// iteration of loop
	start := time.Now()
	for i := 0; i < 2; i++ {
		go postRequest([]int{1, 2, 3}) 
		postRequest([]int{1, 2, 3})
	}
	fmt.Println("Time taken for 2 batches of concurrent requests:", time.Since(start))

	// Call API 10 times, asynchronously, and time total time taken
	// A new goroutine is created for each loop iteration- these all run concurrently
	// with the main goroutine
	var wg sync.WaitGroup
	start = time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("Request number:", i)
		go func() {
			defer wg.Done()
			if err := postRequest([]int{5,6,7}); err != nil {
				fmt.Println("Error:", err)
			}
			}()
		}
	// This blocks the main goroutine until the wg counter is at 0, i.e. all goroutines
	// have completed
	wg.Wait()
	fmt.Println("Time taken for second part requests, i.e. 10 batches, all asynchronously run:", time.Since(start))

}