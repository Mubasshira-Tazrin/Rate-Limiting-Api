# API Authentication and Rate Limiting Server

This Go-based project implements a simple API authentication and rate limiting server using Redis as the backend. It consists of two main components - the main application and an authentication server.

## Components

1. Main Application (`main.go`):
   - Implements a basic HTTP server with two endpoints: `/v1/create-auth-token` and `/v1/postcall`.
   - Generates authentication tokens and stores them in Redis with associated limits.
   - Checks API limits before processing requests and increments usage asynchronously.

2.  Authentication Server (`authserver.go`):
   - Handles protected endpoints by checking authentication headers and verifying API limits.
   - Communicates with a Redis server to retrieve limits and increment usage.

## Setup and Dependencies

1. Go Version:
   - This project is built using Go. Ensure you have Go installed on your machine.

2. Redis:
   - Make sure you have a Redis server running. Update the `redisAddress` variable in `redis.go` if needed.

3. Dependencies:
   - Install project dependencies using `go get`:
     ```bash
     go get github.com/gomodule/redigo/redis
     ```

## Running the Project

1. Main Application:
   - Run the main application on port `:8080`:
     ```bash
     go run main.go
     ```

2. Authentication Server:
   - Run the authentication server on port `:8081`:
     ```bash
     go run authserver.go
     ```

## Endpoints

1. Create Auth Token:
   - Endpoint: `/v1/create-auth-token`
   - Method: `POST`
   - Generates a new authentication token and stores it in Redis.

2. Protected Endpoint:
   - Endpoint: `/v1/postcall`
   - Method: `POST`
   - Requires a valid Authorization header with the authentication token.
   - Checks API limits and increments usage accordingly.

## Redis Keys

- Redis keys are constructed using the format: `global:<auth_token>:limit` and `global:<auth_token>:usage`.

## Configuration

- Update the `redisAddress` variable in `redis.go` if your Redis server is running on a different address or port.

