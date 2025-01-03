package server

import (
	"encoding/json"
	"net/http"
)

type HealthStatus struct {
	Status  string `json:"status"`
	Details string `json:"details,omitempty"`
}

func (s *Server) HealthHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := s.db.Health()
		if err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			json.NewEncoder(w).Encode(HealthStatus{
				Status:  "unhealthy",
				Details: "Database connection failed",
			})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HealthStatus{Status: "healthy"})
	})
}
