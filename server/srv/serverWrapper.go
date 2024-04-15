package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Start starts the HTTP server with the provided configuration.
// It initializes a new server instance and listens on the specified port.
func Start() error {
	// Create a new server instance
	srv := NewServer()

	// Print a message indicating that the server is running
	log.Printf("Server listening on port %s\n", os.Getenv("PORT"))

	// Start the HTTP server and listen on the specified port
	if err := http.ListenAndServe(os.Getenv("PORT"), srv); err != nil {
		return fmt.Errorf("failed to start server: %v", err)
	}

	return nil
}
