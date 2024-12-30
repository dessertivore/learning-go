// Call /addition endpoint 3 times with different values
// and print the result
package main

import (
	"fmt"
	"sync"
)

// func testingGoRoutineWithAndWithoutDefer() {
// 	// Call API twice using a goroutine alongisde a synchronous func call
// 	// The main goroutine does not wait for the new goroutine to complete, it only waits
// 	// for the synchronous (blocking) function to complete before continuing to next
// 	// iteration of loop
// 	start := time.Now()
// 	for i := 0; i < 2; i++ {
// 		go PostToHuma("addition", AdditionInput{NumsToAdd: []int{1, 2, 3}}, &AdditionAPIOutput{})
// 		PostToHuma("addition", AdditionInput{NumsToAdd: []int{1, 2, 3}}, &AdditionAPIOutput{})
// 	}
// 	fmt.Println("Time taken for 2 batches of concurrent requests:", time.Since(start))

// 	// Call API 10 times, asynchronously, and time total time taken
// 	// A new goroutine is created for each loop iteration- these all run concurrently
// 	// with the main goroutine
// 	var wg sync.WaitGroup
// 	start = time.Now()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		fmt.Println("Request number:", i)
// 		go func() {
// 			defer wg.Done()
// 			if err := PostToHuma("addition", AdditionInput{NumsToAdd: []int{5, 6, 7}}, &AdditionAPIOutput{}); err != nil {
// 				fmt.Println("Error:", err)
// 			}
// 		}()
// 	}
// 	// This blocks the main goroutine until the wg counter is at 0, i.e. all goroutines
// 	// have completed
// 	wg.Wait()
// 	fmt.Println("Time taken for second part requests, i.e. 10 batches, all asynchronously run:", time.Since(start))

// }

func postAllFruit(fruitList *ShoppingJSON) *ShoppingJSON {
	var wg sync.WaitGroup
	// Create a channel to receive responses from goroutines
	messages := make(chan *ShoppingJSON, 4)
	// If fruitList is not nil and has items, send it to the API
	if fruitList != nil && len(fruitList.Items) > 0 {
		resp, err := PostToHuma("shopping", fruitList, &ShoppingJSON{})
		if err != nil {
			fmt.Println("Error:", err)
		}
		respJSON, ok := resp.(*ShoppingJSON)
		if !ok {
			fmt.Println("Response is not a *ShoppingJSON")
		}
		// Send the response to the messages channel
		messages <- respJSON
		// If fruitList is nil or empty, send all fruits to the API, 1 by 1
	} else {
		for _, i := range []string{"banana", "apple", "orange", "strawberries"} {
			wg.Add(1)
			go func() {
				defer wg.Done()
				resp, err := PostToHuma("shopping", ShoppingJSON{Items: []string{i}}, &ShoppingJSON{})
				if err != nil {
					fmt.Println("Error:", err)
				}
				fmt.Println(resp)
				respJSON, ok := resp.(*ShoppingJSON)
				if !ok {
					fmt.Println("Response is not a *ShoppingJSON")
				}
				// Send the response to the messages channel
				messages <- respJSON

			}()
		}
	}
	// Close the messages channel once all goroutines are done
	wg.Wait()
	close(messages)
	outerResp := <-messages
	return outerResp
}

func main() {
	// testingGoRoutineWithAndWithoutDefer()
	resp := postAllFruit(*new(*ShoppingJSON))
	resp = postAllFruit(resp)
	fmt.Println("Final test output:", resp)
}
