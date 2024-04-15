package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	router "servermodule/pkg"
	models "servermodule/servermodels"
)

// The server struct represents the HTTP server instance.
// It holds a reference to a router, which handles incoming HTTP requests.
type server struct {
	router *router.Router
}

// newServer creates a new server instance with a configured router.
func NewServer() *server {
	s := &server{
		router: router.New(),
	}
	s.configureRouter()

	return s
}

// The configureRouter() method configures the router with the necessary route handlers.
// Currently, it only sets up a handler for the /network endpoint using the GET method.
func (s *server) configureRouter() {
	s.router.GET("/network", s.requestHandler())
}

// The requestHandler() method is the handler function for the /network endpoint.
// It retrieves interface details based on query parameters.
// This function is returned as an http.HandlerFunc.
func (s *server) requestHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParams := r.URL.Query()
		interfaceParam := queryParams.Get("interface")

		// Check if only "interface" query parameter is provided
		if len(queryParams) == 1 && interfaceParam != "" {
			// Retrieve details of the specified interface
			interfaceDetails, err := models.GetInterfaceByName(interfaceParam)
			if err != nil {
				s.error(w, http.StatusNotFound, err)
				return
			}

			var test models.NetworkInterfaces
			test.Interfaces = append(test.Interfaces, *interfaceDetails)

			s.respond(w, http.StatusOK, test)
			return
		}

		// Check if no query parameter is provided
		if len(queryParams) == 0 {
			// Retrieve details of all network interfaces
			interfaces, err := models.GetInterfaces()
			if err != nil {
				s.error(w, http.StatusInternalServerError, err)
				return
			}
			var test models.NetworkInterfaces
			test.Interfaces = interfaces

			s.respond(w, http.StatusOK, test)
			return
		}

		// If any other query parameter is provided, return an error
		s.error(w, http.StatusBadRequest, errors.New("only ?interface={interface_name} input format is allowed"))
	}
}

// ServeHTTP handles incoming HTTP requests by delegating them to the router.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// The error() method responds to HTTP errors with a JSON error message.
// It takes the HTTP status code and an error object as input parameters and returns a JSON response with the error message.
func (s *server) error(w http.ResponseWriter, code int, err error) {
	log.Printf("HTTP error %d: %s", code, err.Error()) // Log the error
	s.respond(w, code, models.Error{Error: err.Error()})
}

// The respond() method sets the content type to JSON and writes the provided data to the response writer.
// It takes the HTTP status code and any data to be sent in the response body as input parameters.
func (s *server) respond(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
