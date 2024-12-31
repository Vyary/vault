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

func (s *Server) GetRunes2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetExch("runes2")
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		utils.Encode(w, r, u, true)
	})
}

func (s *Server) GetCores2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetExch("cores2")
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		utils.Encode(w, r, u, true)
	})
}

func (s *Server) GetFragments2() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, err := s.db.GetExch("fragments2")
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", err, nil)
			return
		}

		utils.Encode(w, r, u, true)
	})
}
