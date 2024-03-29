package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateAuthToken() string {
	tokenLength := 32
	randomBytes := make([]byte, tokenLength)
	_, err := rand.Read(randomBytes)
	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}

	authToken := base64.URLEncoding.EncodeToString(randomBytes)
	return authToken
}

func constructRedisKeys(authToken string) (usageKey, limitKey string) {
	return fmt.Sprintf("global:%s:usage", authToken), fmt.Sprintf("global:%s:limit", authToken)
}
