package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestClient_CallEndpoint(t *testing.T) {
	// Set up a mock HTTP server
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Respond with a mock JSON response
		mockResponse := `{
			"network_interface": [
				{
					"name": "eth0",
					"ip_addresses": ["192.168.1.10"],
					"mac_address": "00:11:22:33:44:55",
					"mtu": 1500,
					"speed": "1 Gbps",
					"duplex": "Full",
					"admin_status": "enabled",
					"operational_status": "UP"
				}
			]
		}`
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	})
	mockServer := httptest.NewServer(mockHandler)
	defer mockServer.Close()

	// Set up environment variables
	os.Setenv("HOST", mockServer.URL)
	os.Setenv("PORT", "")
	os.Setenv("API_ENDPOINT", "/")

	// Create a new client with a short interval for testing
	client := NewClient(mockServer.URL, 100*time.Millisecond)

	// Call the endpoint
	client.CallEndpoint()

	// Sleep for a short duration to allow the client to make another call
	time.Sleep(200 * time.Millisecond)
}
