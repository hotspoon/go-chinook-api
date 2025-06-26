# Chinook API

A simple RESTful API built with Go and [Gin](https://github.com/gin-gonic/gin) for accessing artist data from the [Chinook](https://github.com/lerocha/chinook-database) SQLite sample database.

## Features

- List all artists
- Get artist by ID
- Built with idiomatic Go and Gin
- Uses raw SQL with the standard library

## Project Structure

```
.
├── main.go
├── chinook.db
├── internal/
│   ├── config/         # Database setup
│   ├── handlers/       # HTTP handlers
│   ├── models/         # Data models
│   ├── repositories/   # Data access logic
│   └── routes/         # API route definitions
└── .air.toml           # Live reload config (Air)
```

## Getting Started

### Prerequisites

- Go 1.18+
- [chinook.db](https://github.com/lerocha/chinook-database) (already included)

### Install dependencies

```sh
go mod tidy
```

### Run the API

```sh
go run main.go
```

Or use [Air](https://github.com/cosmtrek/air) for live reload:

```sh
air
```

The API will be available at [http://localhost:8080](http://localhost:8080).

## API Endpoints

| Method | Endpoint           | Description      |
| ------ | ------------------ | ---------------- |
| GET    | `/api/artists`     | List all artists |
| GET    | `/api/artists/:id` | Get artist by ID |

## Example

Get all artists:

```sh
curl http://localhost:8080/api/artists
```

Get artist by ID:

```sh
curl http://localhost:8080/api/artists/1
```
