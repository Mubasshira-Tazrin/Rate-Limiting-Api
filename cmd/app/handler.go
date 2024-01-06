package app

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/redis"
	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/utils"
)

type AuthHandlers struct {
	Logger *slog.Logger
}

func (ah AuthHandlers) createAuthTokenHandler(w http.ResponseWriter, _ *http.Request) {
	authToken := utils.GenerateAuthToken()
	utils.ConstructRedisKeys(authToken)

	writeResponse(w, http.StatusOK, fmt.Sprintf("auth_token: %s", authToken))
}

func (ah AuthHandlers) performPostCallHandler(w http.ResponseWriter, r *http.Request) {
	authToken := r.Header.Get("Authorization")

	if authToken == "" {
		writeResponse(w, http.StatusUnauthorized, "authorization key is missing")
		return
	}

	usageKey, limitKey := utils.ConstructRedisKeys(authToken)

	limit, usage, err := redis.GetRedisValues(limitKey, usageKey)
	if err != nil {
		ah.Logger.Error("error retrieving limit/usage from Redis", "err", err.Error())
		writeResponse(w, http.StatusForbidden, "error retrieving limit/usage from Redis")

		return
	}

	switch {
	case limit == 0:
		writeResponse(w, http.StatusForbidden, "limit key not found in Redis")
	case limit > usage:
		go redis.IncrementUsage(usageKey, ah.Logger)
		writeResponse(w, http.StatusOK, "API call successful")
	default:
		writeResponse(w, http.StatusForbidden, "rate limit exceeded")
	}
}

func writeResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	_, _ = fmt.Fprint(w, message)
}
