package main

import (
	"log/slog"
	"net/http"
	"sushiroll/internal/handlers"
)

func main() {
	addr := ":8080"

	mux := handlers.SetupRoutes()

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	slog.Info("Starting web server", "addr", addr)
	err := server.ListenAndServe()
	panic(err)
}
