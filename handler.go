package main

import (
	"fmt"
	"net/http"
)

func createAuthTokenHandler(w http.ResponseWriter, r *http.Request) {
	authToken := generateAuthToken()
	storeAuthTokenInRedis(authToken)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "auth_token: %s", authToken)
}

func performPostCallHandler(w http.ResponseWriter, r *http.Request) {

	authToken := r.Header.Get("Authorization")

	if authToken == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Authorization key is missing")
		return
	}

	usageKey, limitKey := constructRedisKeys(authToken)

	limit, usage, err := getRedisValues(limitKey, usageKey)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error retrieving limit/usage from Redis")
		return
	}

	if limit == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Limit key not found in Redis")
		return
	}

	if limit > usage {
		go incrementUsage(usageKey)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "API call successful")
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "API limit exceeded")
	}
}
