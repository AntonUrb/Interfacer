package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	models "servermodule/servermodels"
	server "servermodule/srv"
	"testing"
)

// TestNetworkEndpointParams tests different scenarios related to the /network endpoint parameters.
func TestNetworkEndpointParams(t *testing.T) {
	tests := []struct {
		name          string
		query         string
		expectedCode  int
		expectedError string
	}{
		{
			name:          "ValidParam",
			query:         "interface=lo",
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
		{
			name:          "InvalidInput",
			query:         "interfa=lo",
			expectedCode:  http.StatusBadRequest,
			expectedError: "only ?interface={interface_name} input format is allowed",
		},
		{
			name:          "NoParam",
			query:         "",
			expectedCode:  http.StatusOK,
			expectedError: "",
		},
		{
			name:          "InvalidParam",
			query:         "interface=test",
			expectedCode:  http.StatusNotFound,
			expectedError: "there is no such interface",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/network?"+test.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			// Create a new server instance
			srv := server.NewServer()

			// Serve the request to the server
			srv.ServeHTTP(rr, req)

			// Check if the response status code matches the expected code
			if status := rr.Code; status != test.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, test.expectedCode)
			}

			// Check if the response body contains the expected error message
			var actualError models.Error
			if err := json.NewDecoder(rr.Body).Decode(&actualError); err != nil {
				t.Errorf("failed to decode response body: %v", err)
			}

			if actualError.Error != test.expectedError {
				t.Errorf("error message mismatch: got %q, want %q", actualError.Error, test.expectedError)
			}
		})
	}
}
