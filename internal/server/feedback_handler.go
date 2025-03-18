package server

import (
	"log/slog"
	"net/http"
	"vault/internal/models"
	"vault/internal/utils"
)

func (s *Server) FeedbackHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), "POST feedback")
		defer span.End()

		var feedback models.Feedback
		if err := utils.Decode(r, &feedback); err != nil {
			utils.Error(w, http.StatusBadRequest, "Invalid request body", err, nil)
			return
		}

		problems := feedback.Validate()
		if len(problems) > 0 {
			utils.Error(w, http.StatusUnprocessableEntity, "Feedback validation failed.", nil, problems)
			return
		}

		err := s.db.SaveFeedback(ctx, feedback)
		if err != nil {
			utils.Error(w, http.StatusInternalServerError, "An internal server error occurred. Please try again later.", nil, nil)
			slog.Error("creating strategy", "error", err, "message", feedback)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
