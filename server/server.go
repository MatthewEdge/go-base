package server

import "net/http"

// Server creates a contextual HTTP server with shared dependencies
type Server struct {
	router *http.ServeMux
}

// New returns a new Server with routes initialized
func New() *Server {

	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	s := &Server{
		router: r,
	}

	return s
}

// ServeHTTP to allow Server to be a HTTP Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
