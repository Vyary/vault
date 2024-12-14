package middleware

import (
	"net/http"
	"time"
)

func CacheControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const cacheDuration = 10 * time.Minute
		const cacheControlHeader = "public, max-age=600"

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
