# Shortlink Backend API

<p align="center">
  <a href="https://go.dev/"><img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" /></a>
  <a href="https://github.com/gin-gonic/gin"><img src="https://img.shields.io/badge/Gin-008B8B?style=for-the-badge&logo=gin&logoColor=white" alt="Gin" /></a>
  <a href="https://www.postgresql.org/"><img src="https://img.shields.io/badge/PostgreSQL-4169E1?style=for-the-badge&logo=postgresql&logoColor=white" alt="PostgreSQL" /></a>
  <a href="https://redis.io/"><img src="https://img.shields.io/badge/Redis-DC382D?style=for-the-badge&logo=redis&logoColor=white" alt="Redis" /></a>
  <a href="https://jwt.io/"><img src="https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white" alt="JWT" /></a>
  <a href="https://www.docker.com/"><img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" /></a>
  <a href="https://swagger.io/"><img src="https://img.shields.io/badge/Swagger-85EA2D?style=for-the-badge&logo=swagger&logoColor=black" alt="Swagger" /></a>
</p>

A RESTful backend for a shortlink application, built with Go and the Gin framework. It handles user authentication, custom URL shortening, high-performance link redirection, click analytics and tracking, and user link management. Built as a submission test during the Koda Academy bootcamp, Gin Exercise Koda Batch 9.

The frontend lives in a separate repository: [shortlink_frontend](https://github.com/alfannurfaiz9/shortlink_frontend).

## Features

**Authentication**
- Register and login with JWT (HS256). The token is also tracked in Redis, so a logout invalidates it immediately.
- Passwords are hashed before they are stored.
- Logout removes the session from Redis.

**User**
- Create shortener link
- Delete created link
- Analytics

## Tech Stack

| Technology | Purpose |
| --- | --- |
| [Go](https://go.dev/) | Main programming language |
| [Gin](https://gin-gonic.com/) | HTTP router and web framework |
| [PostgreSQL](https://www.postgresql.org/) | Primary relational database |
| [pgx](https://github.com/jackc/pgx) | PostgreSQL driver and connection pool |
| [Redis](https://redis.io/) | Login sessions, OTP, rate limiting, PIN attempt counters |
| [golang-jwt](https://github.com/golang-jwt/jwt) | JSON Web Token creation and validation |
| [godotenv](https://github.com/joho/godotenv) | Loads `.env` configuration |
| [swaggo](https://github.com/swaggo/swag) | Swagger/OpenAPI documentation |
| [golang-migrate](https://github.com/golang-migrate/migrate) | Database migrations |
| Docker and GitHub Actions | Containerization and image publishing |

## Architecture

The code follows a layered structure, so each layer has one job:

```
Request -> Router -> Middleware -> Controller -> Service -> Repository -> PostgreSQL
                                                   |
                                                   +-> Redis
```

- **Router** registers routes and attaches middleware.
- **Middleware** handles CORS, JWT and Redis session checks.
- **Controller** parses and validates requests and builds responses.
- **Service** contains the business logic.
- **Repository** runs the SQL queries.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (the version in `go.mod`)
- [Docker](https://www.docker.com/)
- [golang-migrate](https://github.com/golang-migrate/migrate) CLI (`migrate`)
- [swag](https://github.com/swaggo/swag) CLI (only if you regenerate the docs)

### 1. Clone the repository

```bash
git clone https://github.com/alfannurfaiz9/shortlink_backend.git
cd shortlink_backend
```

### 2. Configure the environment

Create a `.env` file in the project root and point it to your backend:

```env
HOST={HOST}
PORT={PORT}

DBUSER={DBUSER}
DBPASS={DBPASS}
DBHOST={HOST}
DBPORT={DBPORT}
DBNAME={YOUR_DATABASE_NAME}

REDIS_USER={YOUR_REDIS_USER}
REDIS_PASS={YOUR_REDIS_PASSWORD}
REDIS_HOST={YOUR_REDIS_HOST}
REDIS_PORT={YOUR_REDIST_PORT}

JWT_KEY={YOUR_JWT_KEY}
JWT_ISSUER={YOUR_ISSER}
```

### 3. Pull PostgreSQL and Redis Image

```bash
docker pull postgres:16
docker pull redis:8.4.7
```

### 4. Start PostgreSQL and Redis
Create `.env` file for PostgreSQL
```env
POSTGRES_PASSWORD={YOUR_PASSWORD}
POSTGRES_USER={YOUR_USERNAME}
POSTGRES_DB={YOUR_DATABASE_NAME}
```


PostgreSQL (it reads the `POSTGRES_*` values from `.env`):

```bash
docker run --name {CONTAINER_NAME} -d --env-file {POSTGRES_ENV_PATH} -p {HOST_PORT}:5432 -v {DOCKER_VOLUME}:/var/lib/postgresql postgres:16
```

Redis:

```bash
docker run --name {CONTAINER_NAME} -d -p {HOST_PORT}:6379 -v ./cache/redis.conf:/usr/local/etc/redis/redis.conf redis:8.4.7 redis-server /usr/local/etc/redis/redis.conf
```


### 5. Migrate the database

```bash
make migrate-up
```

### 6. Run the application

```bash
go mod download
go run main.go
```

### 7. Run the react application

[![frontend](https://img.shields.io/badge/FRONTEND-REACT-blue?logo=react&logoColor=WHITE)](https://github.com/alfannurfaiz9/shortlink_frontend.git)

## API Reference

Interactive documentation is served by Swagger UI at **`/docs/index.html`** 

Protected routes require the header `Authorization: Bearer <token>`.

### Auth

| Method | Endpoint | Auth | Description |
| --- | --- | --- | --- |
| `POST` | `/auth/register` | No | Create a new account |
| `POST` | `/auth/login` | No | Log in and receive a JWT |
| `POST` | `/auth/logout` | Yes | Log out and invalidate the session |

### User

| Method | Endpoint | Auth | Description |
| --- | --- | --- | --- |
| `POST` | `/links` | Yes | Create a short link |
| `GET` | `/links` | Yes | Get user's link |
| `DELETE` | `/links/:id` | Yes | Delete a link |
| `GET` | `/:slug` | No | Redirect to original URL |

## Security Notes

- **Sessions:** a JWT is valid only while its session exists in Redis (60 minutes). Logging out removes it right away.
- **CORS:** allowed origins are defined in `internal/middleware/cors.middleware.go`. Update them for your own frontend.

## Makefile Commands

| Command | Description |
| --- | --- |
| `make migrate-create NAME=<name>` | Create a new migration pair |
| `make migrate-up` | Apply all migrations |
| `make migrate-down N=<count>` | Roll back N migrations |
| `make db-reset` | Roll back, migrate, and seed again |

## Changelog

| Version | Desctiption |
| ------- | ----------- |
| 1.0.0   | Initial app |

## License

This project is licensed under the [MIT License](https://opensource.org/license/mit).

## Related Project

[![frontend](https://img.shields.io/badge/FRONTEND-REACT-blue?logo=react&logoColor=WHITE)](https://github.com/alfannurfaiz9/shortlink_backend.git)
