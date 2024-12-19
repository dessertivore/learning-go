package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

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
type APIOutput struct {
	Body struct {
		MainOutput string `json:"message" doc:"Output of API call"`
}}

// Define a struct for the addition input
type AdditionInput struct {
	Body struct {
    Input string `json:"numsToAdd" maxLength:"100" doc:"Numbers to add together, comma-separated"`}
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
	}, func(ctx context.Context, input *struct{}) (*APIOutput, error) {
		resp := &APIOutput{}
		resp.Body.MainOutput = Restaurants[rand.Intn(5)+1]
		return resp, nil
	})
	// Register POST /addition.
	huma.Register(api, huma.Operation{
		OperationID: "post-sleep-and-add",
		Method:      http.MethodPost,
		Path:        "/addition",
		Summary:     "Sleep for a bit then add comma-separated numbers.",
		Description: "A slow endpoint.",
		Tags:        []string{"Addition","Sleep"},
	}, func(ctx context.Context, input *AdditionInput) (*APIOutput, error) {
		resp := &APIOutput{}
		parts:=strings.Split(input.Body.Input, ",")
		
		log.Printf("Parts: %v", parts)

		var sum int
		for i:= 0; i<len(parts);i++ {
			// Catch errors if invalid int provided
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				return nil, err
			}
			sum += num}
		
		// Sleep for 5 seconds to simulate a big query
		time.Sleep(5 * time.Second)
		resp.Body.MainOutput = strconv.Itoa(sum)
		return resp, nil
	})

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}