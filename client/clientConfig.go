package main

import (
	"os"
)

// Config represents the configuration for the server.
type Config struct {
	Port        string // Port to listen on for incoming HTTP requests.
	Host        string // Host address for the server.
	APIEndpoint string // API endpoint for the server.
}

// NewConfig creates a new instance of Config and reads configuration from environment variables.
func NewConfig() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080" // Default port if not provided
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "http://http-server" // Default host if not provided
	}

	apiEndpoint := os.Getenv("API_ENDPOINT")
	if apiEndpoint == "" {
		apiEndpoint = "/network" // Default API endpoint if not provided
	}

	return &Config{
		Port:        port,
		Host:        host,
		APIEndpoint: apiEndpoint,
	}
}
