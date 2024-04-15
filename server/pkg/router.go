package router

import (
	"net/http"
)

// Router is a simple HTTP router that wraps around http.ServeMux.
type Router struct {
	mux *http.ServeMux
}

// New creates a new instance of Router.
func New() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

// Handle registers a handler for the given pattern.
func (r *Router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

// GET registers a handler for the HTTP GET method and the given pattern.
func (r *Router) GET(pattern string, fn http.HandlerFunc) {
	r.Handle(http.MethodGet+" "+pattern, fn)
}

// ServeHTTP dispatches the request to the handler registered to handle it.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
