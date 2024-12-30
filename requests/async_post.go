// Call /addition endpoint 3 times with different values
// and print the result
package main

import (
	"fmt"
	"sync"
	"time"
)

func testingGoRoutineWithAndWithoutDefer() {
	// Call API twice using a goroutine alongisde a synchronous func call
	// The main goroutine does not wait for the new goroutine to complete, it only waits
	// for the synchronous (blocking) function to complete before continuing to next
	// iteration of loop
	start := time.Now()
	for i := 0; i < 2; i++ {
		go PostAdditionRequest([]int{1, 2, 3})
		PostAdditionRequest([]int{1, 2, 3})
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
			if err := PostAdditionRequest([]int{5, 6, 7}); err != nil {
				fmt.Println("Error:", err)
			}
		}()
	}
	// This blocks the main goroutine until the wg counter is at 0, i.e. all goroutines
	// have completed
	wg.Wait()
	fmt.Println("Time taken for second part requests, i.e. 10 batches, all asynchronously run:", time.Since(start))

}

func main() {
	testingGoRoutineWithAndWithoutDefer()
}
