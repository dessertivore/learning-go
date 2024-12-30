package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"testing_go_apis/data"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

func getRestaurants(c echo.Context) error {
	return c.JSON(http.StatusOK, data.Restaurants)
}

type AdditionBody struct {
	NumsToAdd string `json:"numsToAdd" xml:"numsToAdd" form:"numsToAdd" query:"numsToAdd"`
}

func addition(c echo.Context) error {
	start := time.Now()

	// Get nums to add from the request
	nums := new(AdditionBody)
	if err := c.Bind(nums); err != nil {
		return err
	}
	// Split the string of numbers into an array of ints
	numStrings := strings.Split(nums.NumsToAdd, ",")
	numArray := make([]int, len(numStrings))
	for i := range numArray {
		numArray[i], _ = strconv.Atoi(numStrings[i])
	}

	log.Printf("Nums to add: %v", numArray)

	var sum int
	for _, num := range numArray {
		// Catch errors if invalid int provided
		sum += num
	}

	var wg sync.WaitGroup
	// Repeat sleep 10 times, but asynchronously, so it will only take
	// 5 seconds instead of 50 seconds
	for i := 0; i < 10; i++ {
		wg.Add(1)
		// Sleep for 5 seconds
		go func() {
			defer wg.Done()
			fmt.Println("Sleeping for 5 seconds ", i, " times")
			time.Sleep(5 * time.Second)
		}()
	}
	wg.Wait()
	fmt.Println("Total time taken:", time.Since(start))

	// Return the sum as a string
	return c.String(http.StatusOK, strconv.Itoa(sum))
}

func main() {
	e := echo.New()
	//  Basic Hello World
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// Pick a random restaurant from the list
	e.GET("/restaurant", func(c echo.Context) error {
		return c.String(http.StatusOK, "You should go to "+data.Restaurants[rand.Intn(5)+1])
	})
	//  Return the dict of restaurants
	e.GET("/restaurants", getRestaurants)
	// Add the numbers in the request and sleep 5 seconds * 10 times (asynchronously)
	e.POST("/addition", addition)
	e.Logger.Fatal(e.Start(":8080"))
}
