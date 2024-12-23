package main

import (
	"net/http"
	"testing_go_apis/data"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/rand"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/restaurant", func(c echo.Context) error {
		return c.String(http.StatusOK, "You should go to "+data.Restaurants[rand.Intn(5)+1])
	})
	e.Logger.Fatal(e.Start(":8080"))
}
