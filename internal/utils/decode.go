package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(r *http.Request, data any) error {
	if r.Body == nil {
		return fmt.Errorf("request body is nil")
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return fmt.Errorf("failed to decode JSON: %w", err)
	}

	return nil
}
