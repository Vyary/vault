package server

import (
	"net/http"

	"vault/internal/utils"
)

func (s *Server) GetUniques2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetUniques2()
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		utils.Encode(w, r, u, true)
	})
}
