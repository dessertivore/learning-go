### A basic API using Huma, which randomly picks a restaurant for dinner

To start the server:
```sh
go run .
```

From another terminal:
```sh
curl http://localhost:8888/
```

# Endpoints:
`/restaurant` - `GET`

`/addition` - `POST` - e.g. `{"numsToAdd": "1,2,3,4,5"}` returns 
```
{
    "$schema": "http://127.0.0.1:8888/schemas/APIOutputBody.json",
    "message": "15"
}
```