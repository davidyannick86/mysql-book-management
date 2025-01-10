package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseBody(r *http.Request, model interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(body), model); err != nil {
		return err
	}
	return nil
}
