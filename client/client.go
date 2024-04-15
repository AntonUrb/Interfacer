package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	models "clientmodule/clientmodels"
)

// Client represents the HTTP client.
type Client struct {
	endpoint string        // The endpoint to call.
	interval time.Duration // The interval between calls.
}

// NewClient creates a new instance of Client.
func NewClient(endpoint string, interval time.Duration) *Client {
	return &Client{
		endpoint: endpoint,
		interval: interval,
	}
}

// Start starts the HTTP client.
func (c *Client) Start() {
	// Call the endpoint immediately
	c.CallEndpoint()

	// Create a ticker for periodic calls to the endpoint
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	// Loop to continuously call the endpoint at the specified interval
	for range ticker.C {
		c.CallEndpoint()
	}
}

func (c *Client) CallEndpoint() {
	// Make a GET request to the server's endpoint
	resp, err := http.Get(c.endpoint)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into a slice of NetworkInterface
	var body models.NetworkInterfaces
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	// Check for errors returned from the server
	if body.Error != "" {
		fmt.Println("Error:", body.Error)
		return
	}

	// Print the network interfaces
	fmt.Println("Network Interfaces:")
	for _, iface := range body.Interfaces {
		fmt.Printf("Name: %s\n", iface.Name)
		fmt.Println("IP Addresses:")
		if len(iface.IPAddresses) == 0 {
			fmt.Println("\tnull")
		} else {
			for _, ip := range iface.IPAddresses {
				fmt.Printf("\t%s\n", ip)
			}
		}
		if iface.MACAddress == "" {
			fmt.Println("MAC Address: not found")
		} else {
			fmt.Printf("MAC Address: %s\n", iface.MACAddress)
		}
		fmt.Printf("MTU: %d\n", iface.MTU)
		fmt.Printf("Speed: %s\n", iface.Speed)
		fmt.Printf("Duplex: %s\n", iface.Duplex)
		fmt.Printf("Admin Status: %s\n", iface.AdminStatus)
		fmt.Printf("Operational Status: %s\n", iface.OperationalStatus)
		fmt.Println()
	}
}

func main() {
	// Construct the server endpoint using environment variables
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	apiEndpoint := os.Getenv("API_ENDPOINT")
	interfaceName := os.Getenv("INTERFACE")

	endPoint := host + port + apiEndpoint
	if interfaceName != "" {
		endPoint += "?interface=" + interfaceName
	}
	fmt.Println("Server endpoint:", endPoint)

	// Parse interval duration from environment variable or default to 5 seconds
	intervalDuration, err := time.ParseDuration(os.Getenv("INTERVAL"))
	if err != nil {
		intervalDuration = 5 * time.Second
	}

	// Create a new HTTP client with the specified endpoint and interval
	client := NewClient(endPoint, intervalDuration)

	// Start the HTTP client
	client.Start()
}
