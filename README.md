# Chinook API

A RESTful API built with Go and [Gin](https://github.com/gin-gonic/gin) for accessing and managing data from the [Chinook](https://github.com/lerocha/chinook-database) SQLite sample database.

## Features

- JWT authentication (login, signup, refresh token)
- List, create, update, and delete artists and albums
- Get artist/album by ID
- Secure endpoints with Bearer token
- Swagger/OpenAPI documentation
- Structured logging with Zerolog
- Custom error handling (404/500)
- Health check endpoint

## Project Structure

```
.
├── main.go
├── chinook.db
├── internal/
│   ├── config/         # Configuration and DB setup
│   ├── handlers/       # HTTP handlers
│   ├── logging/        # Logging setup (Zerolog)
│   ├── models/         # Data models
│   ├── repositories/   # Data access logic
│   ├── routes/         # API route definitions
│   └── utils/          # Utility functions
├── docs/               # Swagger docs
├── .air.toml           # Live reload config (Air)
├── .env                # Environment variables
├── .env.example        # Example env file
├── app.log             # Application log file
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.18+
- [chinook.db](https://github.com/lerocha/chinook-database) (already included)

### Install dependencies

```sh
go mod tidy
```

### Environment Setup

Copy `.env.example` to `.env` and fill in required values:

```sh
cp .env.example .env
```

Set at least:

```
JWT_SECRET=your_jwt_secret
API_VERSION=v1
FRONTEND_WEB_URL=http://localhost:3000
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

## Authentication

### Signup

```sh
curl -X POST http://localhost:8080/api/v1/auth/signup \
  -H "Content-Type: application/json" \
  -d '{"username":"youruser","email":"your@email.com","password":"yourpassword"}'
```

### Login

```sh
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"youruser","password":"yourpassword"}'
```

Response:

```json
{
  "token": "<jwt_token>",
  "refresh_token": "<refresh_token>"
}
```

### Refresh Token

```sh
curl -X POST http://localhost:8080/api/v1/auth/refresh \
  -H "Content-Type: application/json" \
  -d '{"refresh_token":"<refresh_token>"}'
```

### Authenticated Requests

Add the JWT token to the `Authorization` header:

```sh
curl http://localhost:8080/api/v1/artists \
  -H "Authorization: Bearer <jwt_token>"
```

## API Endpoints

| Method | Endpoint                      | Description                | Auth Required |
| ------ | ----------------------------- | -------------------------- | ------------- |
| GET    | `/api/v1/health`              | Health check               | No            |
| POST   | `/api/v1/auth/signup`         | Register new user          | No            |
| POST   | `/api/v1/auth/login`          | Login and get tokens       | No            |
| POST   | `/api/v1/auth/refresh`        | Refresh JWT token          | No            |
| GET    | `/api/v1/auth/me`             | Get current user info      | Yes           |
| GET    | `/api/v1/artists`             | List all artists           | Yes           |
| GET    | `/api/v1/artists/:id`         | Get artist by ID           | Yes           |
| POST   | `/api/v1/artists`             | Create artist              | Yes           |
| PUT    | `/api/v1/artists/:id`         | Update artist              | Yes           |
| DELETE | `/api/v1/artists/:id`         | Delete artist              | Yes           |
| GET    | `/api/v1/albums`              | List all albums            | Yes           |
| GET    | `/api/v1/albums/:id`          | Get album by ID            | Yes           |
| POST   | `/api/v1/albums`              | Create album               | Yes           |
| PUT    | `/api/v1/albums/:id`          | Update album               | Yes           |
| DELETE | `/api/v1/albums/:id`          | Delete album               | Yes           |

## Swagger Documentation

Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) for interactive API docs.

## Logging

- All requests and errors are logged in structured JSON format to `app.log` using Zerolog.

## Example Usage

Get all artists (authenticated):

```sh
curl http://localhost:8080/api/v1/artists \
  -H "Authorization: Bearer <jwt_token>"
```

Get artist by ID:

```sh
curl http://localhost:8080/api/v1/artists/1 \
  -H "Authorization: Bearer <jwt_token>"
```

## License

MIT