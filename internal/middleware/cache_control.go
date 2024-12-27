package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func CacheControl(next http.Handler, minutes int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cacheDuration := time.Duration(minutes) * time.Minute
		cacheControlHeader := fmt.Sprintf("public, max-age=%.0f", cacheDuration.Seconds())

		now := time.Now()
		ifModifiedSince := r.Header.Get("If-Modified-Since")

		if ifModifiedSince != "" {
			t, err := time.Parse(http.TimeFormat, ifModifiedSince)
			if err == nil && !t.Before(now.Add(-cacheDuration)) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}

		w.Header().Set("Last-Modified", now.UTC().Format(http.TimeFormat))
		w.Header().Set("Cache-Control", cacheControlHeader)

		next.ServeHTTP(w, r)
	})
}
