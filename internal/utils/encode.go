package utils

import (
	"compress/gzip"
	"encoding/json"
	"net/http"
	"strings"
)

func Encode(w http.ResponseWriter, r *http.Request, data any, compress bool) {
	w.Header().Set("Content-Type", "application/json")

	if compress && strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()

		if err := json.NewEncoder(gz).Encode(data); err != nil {
			Error(w, http.StatusInternalServerError, "Failed to gzip response",err, nil)
			return
		}
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		Error(w, http.StatusInternalServerError, "Failed to encode response",err, nil)
		return
	}
}
