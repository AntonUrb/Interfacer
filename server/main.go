package main

import (
	"log"
	server "servermodule/srv"
)

func main() {
	// Start the server with the provided configuration
	log.Fatal(server.Start())
}
