package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	testCases := []struct {
		name       string
		statusCode int
		message    string
		expected   string
	}{
		{"Success", http.StatusOK, "API call successful", "API call successful"},
		{"Unauthorized", http.StatusUnauthorized, "authorization key is missing", "authorization key is missing"},
		{"Forbidden", http.StatusForbidden, "error retrieving limit/usage from Redis", "error retrieving limit/usage from Redis"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new HTTP response recorder
			recorder := httptest.NewRecorder()

			// Call the writeResponse function with the test case parameters
			writeResponse(recorder, tc.statusCode, tc.message)

			// Check if the recorded response matches the expected output
			if got := recorder.Body.String(); got != tc.expected {
				t.Errorf("Expected response: %s, but got: %s", tc.expected, got)
			}

			// Check if the status code matches the expected status code
			if got := recorder.Code; got != tc.statusCode {
				t.Errorf("Expected status code: %d, but got: %d", tc.statusCode, got)
			}
		})
	}
}
