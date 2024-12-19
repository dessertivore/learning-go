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

func postRequest() error {
	start := time.Now()
	// HTTP endpoint
	postURL := "http://127.0.0.1:8888/addition"

	// JSON body
	body := []byte(`{
		"numsToAdd": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	}`)

	// Create a HTTP post request
	r, err := http.NewRequest("POST", postURL, bytes.NewBuffer(body))
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
	// Call API 10 times, asynchronously, and time total time taken
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("Request number:", i)
		go func() {
			defer wg.Done()
			if err := postRequest(); err != nil {
				fmt.Println("Error:", err)
			}
			}()
		}
	wg.Wait()
	fmt.Println("Total time taken:", time.Since(start))

}