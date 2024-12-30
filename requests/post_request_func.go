package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AdditionAPIOutput struct {
	MainOutput int `json:"sum" doc:"Sum of numbers inputted"`
}

type AdditionInput struct {
	NumsToAdd []int `json:"numsToAdd"`
}

type ShoppingJSON struct {
	Items []string `json:"items"`
}

func PostToHuma(endpoint string, input interface{}, output interface{}) (interface{}, error) {
	start := time.Now()
	// Construct the full URL
	postURL := fmt.Sprintf("%s%s", "http://127.0.0.1:8888/", endpoint)
	// JSON body
	// input := AdditionInput{NumsToAdd: nums}
	jsonData, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("error marshaling input: %v", err)
	}
	fmt.Println("Sending:", string(jsonData))

	// Create a HTTP post request
	r, err := http.NewRequest("POST", postURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	derr := json.NewDecoder(res.Body).Decode(output)
	if derr != nil {
		fmt.Println("Error:", derr)
		return nil, derr
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status:", res.StatusCode)
		return nil, fmt.Errorf("status: %s", res.Status)
	}
	full := time.Since(start)
	fmt.Println("Output:", output)
	fmt.Println("Time taken:", full)
	return output, nil
}
