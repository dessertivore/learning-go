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

func postAllFruit() {
	var wg sync.WaitGroup
	for _, i := range []string{"banana", "apple", "orange", "strawberries"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := PostToHuma("shopping", ShoppingJSON{Items: []string{i}}, &ShoppingJSON{})
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println(resp)
		}()
	}
	wg.Wait()
}

func main() {
	// testingGoRoutineWithAndWithoutDefer()
	for i := 0; i < 10; i++ {
		postAllFruit()
	}
}
