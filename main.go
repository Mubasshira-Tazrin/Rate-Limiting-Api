package main

import "net/http"

var redisAddress = "localhost:6379"

func main() {
	http.HandleFunc("/v1/create-auth-token", createAuthTokenHandler)
	http.HandleFunc("/v1/postcall", performPostCallHandler)

	http.ListenAndServe(":8080", nil)
}
