package app

import (
	"net/http"
	"os"
	"time"

	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/constants"
)

func StartAPIRateLimiter() {
	// init logger(slog)
	l := initSlogger()

	// init handlers
	ah := AuthHandlers{Logger: l}

	http.HandleFunc("/v1/create-auth-token", ah.createAuthTokenHandler)
	http.HandleFunc("/v1/postcall", ah.performPostCallHandler)

	server := &http.Server{
		Addr:         constants.ServerAddr,
		Handler:      nil, // Using the default handler (Mux: nil means to use http.DefaultServeMux)
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		l.Error("unable to start server", "err", err)
		os.Exit(1) // equivalent to, log.Fatalf
	}
}
