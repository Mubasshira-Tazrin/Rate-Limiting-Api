## Rate Limiting Server

The Rate Limiting API restricts the number of API calls a user can make within a certain time frame. It generates and stores authentication tokens in Redis, associating them with limits on API call usage.

## Components

1. main.go:
   - Entry point for the application, sets up HTTP routes.
1. singleton.go:
   - Implements the singleton pattern for Redis connection pool.
1. redis.go:
   - Manages interactions with Redis, including storing and retrieving token limits and usage.
1. handler.go:
   - Defines HTTP request handlers for creating authentication tokens and making API calls.
1. utils.go:
   - Contains utility functions, including token generation and constructing Redis keys.

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

1. Clone the repository:

     ```bash
     git clone https://github.com/Mubasshira-Tazrin/Rate-Limiting-Api.git
     ```

2. Install dependencies:

     ```bash
      go mod download
     ```
3. Run the application :

   ```bash
      go run main.go
   ```
## Endpoints

1. Create Auth Token:
   - Endpoint: `/v1/create-auth-token`
   - Method: `GET`
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

