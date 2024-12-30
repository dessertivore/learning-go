package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type APIOutput struct {
	MainOutput int `json:"sum" doc:"Sum of numbers inputted"`
}

type AdditionInput struct {
	NumsToAdd []int `json:"numsToAdd"`
}

func PostAdditionRequest(nums []int) error {
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
