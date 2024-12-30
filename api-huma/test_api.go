package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing_go_apis/data"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"

	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// Define a struct for the response
type RestaurantAPIOutput struct {
	Body struct {
		MainOutput string `json:"message" doc:"Output of API call"`
	}
}

// Define a struct for the addition input
type AdditionInput struct {
	Body struct {
		Input []int `json:"numsToAdd" maxLength:"100" doc:"Numbers to add together, in a list."`
	}
}
type AdditionOutput struct {
	Body struct {
		Output int `json:"sum" maxLength:"100" doc:"Sum of inputted numbers."`
	}
}

type ShoppingList struct {
	Body struct {
		Items []string `json:"items" xml:"items" form:"items" query:"items"`
	}
}

func addToShoppingList(listSoFar ShoppingList, additionalElement []string) ShoppingList {
	// Add the items to the shopping list
	listSoFar.Body.Items = append(listSoFar.Body.Items, additionalElement...)
	return listSoFar
}

func main() {
	// Create a new router & API
	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))

	// Register GET /restaurant handler.
	huma.Register(api, huma.Operation{
		OperationID: "get-restaurant",
		Method:      http.MethodGet,
		Path:        "/restaurant",
		Summary:     "Pick a random restaurant for dinner.",
		Description: "Pick a random restaurant.",
		Tags:        []string{"Restaurants"},
	}, func(ctx context.Context, input *struct{}) (*RestaurantAPIOutput, error) {
		resp := &RestaurantAPIOutput{}
		resp.Body.MainOutput = data.Restaurants[rand.Intn(5)+1]
		return resp, nil
	})
	// Register POST /addition.
	huma.Register(api, huma.Operation{
		OperationID: "post-sleep-and-add",
		Method:      http.MethodPost,
		Path:        "/addition",
		Summary:     "Sleep for a bit then add comma-separated numbers.",
		Description: "A slow endpoint.",
		Tags:        []string{"Addition", "Sleep"},
	}, func(ctx context.Context, input *AdditionInput) (*AdditionOutput, error) {
		resp := &AdditionOutput{}

		log.Printf("Nums to sum: %v", input.Body.Input)

		var sum int
		for _, num := range input.Body.Input {
			// Catch errors if invalid int provided
			sum += num
		}

		// Sleep for 5 seconds to simulate a big query
		time.Sleep(5 * time.Second)
		resp.Body.Output = sum
		return resp, nil
	})

	// Register POST /shopping.
	listSoFar := ShoppingList{}
	huma.Register(api, huma.Operation{
		OperationID: "post-shopping",
		Method:      http.MethodPost,
		Path:        "/shopping",
		Summary:     "Add items to the shopping list.",
		Description: "Add items to the shopping list.",
		Tags:        []string{"Shopping"},
	}, func(c context.Context, input *ShoppingList) (*ShoppingList, error) {
		time.Sleep(1 * time.Second)
		fmt.Println("Adding to shopping list:", input.Body.Items)
		listSoFar = addToShoppingList(listSoFar, input.Body.Items)
		fmt.Println("New shopping list:", listSoFar.Body.Items)
		return &listSoFar, nil
	})
	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
