package server

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"vault/internal/utils"
)

func (s *Server) Uniques2Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetUniques2()
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(u); err != nil {
			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}
	})
}

func (s *Server) Exch2Handler(tableName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetExch(tableName)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Encoding", "gzip")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(u); err != nil {
			utils.Error(w, http.StatusInternalServerError, "Failed to encode response", err, nil)
			return
		}
	})
}
