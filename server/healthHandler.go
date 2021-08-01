package server

import "net/http"

// Health Handler to ensure Server is responsive
func (s *Server) healthCheck() http.HandlerFunc {
	type response struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, 200, response{
			Status: "ok",
		})
	}
}

// Readiness Handler to check status of dependencies
func (s *Server) readyCheck() http.HandlerFunc {
	type response struct {
		Status string `json:"status"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, 200, response{
			Status: "ok",
		})
	}
}

// Version Handler to return which version is deployed
func (s *Server) versionCheck() http.HandlerFunc {
	type response struct {
		Version string `json:"version"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		s.WriteJSON(w, 200, response{
			Version: s.version,
		})
	}
}
