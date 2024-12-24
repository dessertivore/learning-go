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
go run requests/async_post.go
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

### Project Euler code
```sh
go run try-euhler/problem_1.go "1000"
```
To find all multiples of 3 and 5 under 1000. Can edit sys arg as desired to change
upper limit.