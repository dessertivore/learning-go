package main

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/go-chi/chi/v5"

	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

var Restaurants = map[int]string{
	1:"Purezza",
	2: "Temple of Seitan",
	3: "Mildreds",
	4: "Club Mexicana",
	5: "Unity Diner",
	6: "Dauns Deli",
}

// Define a struct for the response
type PickFood struct {
	Body struct {
		RestaurantSuggestion string `json:"message" doc:"Restaurant suggestion"`
	}
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
	}, func(ctx context.Context, input *struct{}) (*PickFood, error) {
		resp := &PickFood{}
		resp.Body.RestaurantSuggestion = Restaurants[rand.Intn(5)+1]
		return resp, nil
	})

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}