package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
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
	for _,num:= range numArray {
		// Catch errors if invalid int provided
		sum += num}
		
		// Sleep for 5 seconds to simulate a big query
		time.Sleep(5 * time.Second)
	
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
	
	e.POST("/addition", addition)
	e.Logger.Fatal(e.Start(":8080"))
}
