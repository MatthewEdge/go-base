package server

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// Server creates a contextual HTTP server with shared dependencies
type Server struct {
	router *chi.Mux
	db     *sqlx.DB

	// Deployed version
	version string
}

// New returns a new Server with routes initialized
func New(db *sqlx.DB, version string) *Server {

	s := &Server{
		db:      db,
		version: version,
	}

	r := chi.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	// Health and Readiness
	r.Get("/health", s.healthCheck())
	r.Get("/ready", s.readyCheck())
	r.Get("/version", s.versionCheck())

	s.router = r
	return s
}

// WriteJSON is a helper to respond with a JSON message body.
// If marshalling fails, it will respond with a HTTP 500
func (s *Server) WriteJSON(w http.ResponseWriter, statusCode int, msg interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(msg)
	if err != nil {
		s.WriteError(w, 500, "Response marshalling failed")
	}
}

// Standard error JSON
type error struct {
	Error string `json:"error"`
}

// WriteError responds with a standardized error body
func (s *Server) WriteError(w http.ResponseWriter, statusCode int, msg string) {
	s.WriteJSON(w, statusCode, error{
		Error: msg,
	})
}

// ServeHTTP to allow Server to be a HTTP Handler
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
