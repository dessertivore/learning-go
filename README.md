# 2 basic APIs, using Huma and Echo, with 2 endpoints, and a Project Euler solution
I've just made this repo to practise all things Go. As the title shows, it's a bit of 
a mish mash of things!

To start the Huma server:
```sh
go run api-huma/test_api.go
```

From another terminal:
```sh
curl http://localhost:8888/restaurant
```

To start the Echo server:
```sh
go run api-echo/test_api.go
```
From another terminal:
```sh
curl http://localhost:8080/restaurant
```

### Endpoints - Huma:
`/restaurant` - `GET`

`/addition` - `POST` - e.g. `curl -X POST http://127.0.0.1:8888/addition -H "Content-Type: application/json" -d '{"numsToAdd": "1,2,3,4,5"}'` returns 
```
{
    "$schema": "http://127.0.0.1:8888/schemas/APIOutputBody.json",
    "message": "15"
}
```

### To call /addition endpoint from this module
```sh
go run requests/*.go
```
This will call the API 10 times in quick succession, in an asynchronous fashion,
time each total request as well as the full func. As it is async, even though each
API call takes 5 seconds, the full func should also only take around 5 seconds.

### Endpoints - Echo:
`/` - `GET`
`/restaurant` - `GET`
`/restaurants` - `GET`

`/addition` - `POST` - e.g. `curl -X POST http://127.0.0.1:8080/addition -H "Content-Type: application/json" -d '{"numsToAdd": "1,2,3,4,5"}'` returns `15`
The addition endpoint here will sleep for 5 seconds, 10 times. As the `sleeping` is done asynchronously, the full function will take 5 seconds. Total time taken will be printed to console.

`/shopping` - `POST` - e.g. `curl -X POST http://127.0.0.1:8080/shopping -H "Content-Type: application/json" -d '{"items": ["banana","apple"]}'` returns `["banana","apple"]` and stores the list for you to add to upon next post request. This is a
basic stateful API endpoint to test concurrency with stateful APIs.

### Project Euler code
```sh
go run try-euler/*.go
```
Runs Project Euler problems 1 and 2, prints answers to console.