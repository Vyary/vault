package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Time(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			if rec := recover(); rec != nil {
				slog.Error("Panic occurred", "error", rec)
				w.WriteHeader(http.StatusInternalServerError)
			}
			duration := time.Since(start)
			slog.Info(fmt.Sprintf("%s %s - %s", r.Method, r.URL.Path, duration.String()), "method", r.Method, "path", r.URL.Path, "took", duration.String())
		}()

		next.ServeHTTP(w, r)
	})
}
