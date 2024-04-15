package router

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestRouter_Handle is a test function for the Handle method of the Router struct.
func TestRouter_Handle(t *testing.T) {
	tests := []struct {
		name          string
		method        string
		pattern       string
		handler       http.Handler
		requestMethod string
		requestPath   string
		expectedCode  int
	}{
		{
			name:          "GET request for registered pattern",
			method:        "GET",
			pattern:       "/test",
			handler:       http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}),
			requestMethod: "GET",
			requestPath:   "/test",
			expectedCode:  http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new router instance
			router := New()

			// Register the handler with the router
			router.Handle(tt.pattern, tt.handler)

			// Define a test HTTP request
			req, err := http.NewRequest(tt.requestMethod, tt.requestPath, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a recorder to capture the response
			rr := httptest.NewRecorder()

			// Use the router to handle the request
			router.ServeHTTP(rr, req)

			// Check if the response status code is as expected
			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedCode)
			}
		})
	}
}
