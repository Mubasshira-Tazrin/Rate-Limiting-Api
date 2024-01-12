package main

import "net/http"

var redisAddress = "127.0.0.1:6379"

func main() {
	http.HandleFunc("/v1/create-auth-token", createAuthTokenHandler)
	http.HandleFunc("/v1/postcall", performPostCallHandler)

	http.ListenAndServe(":8080", nil)
}
