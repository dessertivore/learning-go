.PHONY: requests huma echo euler

huma: 
	go run api-huma/test_api.go
echo: 
	go run api-echo/test_api.go
requests:  
	go run requests/*.go
euler:
	go run try-euler/*.go