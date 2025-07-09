# GoAPI-Tutorial

## Run the API
1. `cd goapi-tutorial`
2. `go run cmd/api/main.go`

## Test the API
Use Postman or Curl
```BASH
curl --location 'http://localhost:8080/account/coins/?username=alex' \
--header 'Authorization: 123ABC'
```
Expect response
```JSON
{
    "Code": 200,
    "Balance": 100
}
```