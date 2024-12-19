# A basic API using Huma, with 2 endpoints

To start the server:
```sh
go run api/test_api.go
```

From another terminal:
```sh
curl http://localhost:8888/
```

### Endpoints:
`/restaurant` - `GET`

`/addition` - `POST` - e.g. `{"numsToAdd": "1,2,3,4,5"}` returns 
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