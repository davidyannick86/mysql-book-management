package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// * Parse the request body
func ParseBody(r *http.Request, model interface{}) error {
	// * Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// * Unmarshal the request body
	if err := json.Unmarshal([]byte(body), model); err != nil {
		return err
	}
	return nil
}
