package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "Server: ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("Initializing server ...")
	srv := server.New(logger)

	logger.Println("Server running on port :8080")
	if err := srv.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Server error: %v", err)
	}
}
