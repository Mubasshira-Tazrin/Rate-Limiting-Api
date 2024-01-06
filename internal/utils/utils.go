package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/constants"
)

func GenerateAuthToken() string {
	randomBytes := make([]byte, constants.AuthTokenLen)
	_, err := rand.Read(randomBytes)

	if err != nil {
		fmt.Println("Error generating random bytes:", err)
		return ""
	}

	authToken := base64.URLEncoding.EncodeToString(randomBytes)

	return authToken
}

func ConstructRedisKeys(authToken string) (usageKey, limitKey string) {
	return fmt.Sprintf("global:%s:usage", authToken), fmt.Sprintf("global:%s:limit", authToken)
}
