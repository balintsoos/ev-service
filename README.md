# ev-service

EV service

## Run

```
docker compose up -d
go run main.go
```

## API

```

curl --request POST \
  --header "Content-Type: application/json" \
  --data '{"make":"Tesla","model":"Model 3 RWD","year":2024}' \
  http://localhost:8080/api/v0/vehicles

curl http://localhost:8080/api/v0/vehicles

curl http://localhost:8080/api/v0/vehicles/1

curl --request PUT \
  --header "Content-Type: application/json" \
  --data '{"make":"Tesla","model":"Model 3 RWD","year":2024,"battery_capacity":57}' \
  http://localhost:8080/api/v0/vehicles/1

curl --request DELETE http://localhost:8080/api/v0/vehicles/1
```

## Development

```
go mod init ev-service
go get github.com/gin-gonic/gin
go get github.com/jackc/pgx/v5
```
