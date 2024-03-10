package api

import (
	"encoding/json"
	"errors"
	"io"
)

func toJson(data []byte, obj any) error {
	if obj == nil {
		return errors.New("Invalid struct pointer.")
	}

	err := json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	return nil
}

func readBody(body io.ReadCloser) ([]byte, error) {
	defer body.Close()
	return io.ReadAll(body)
}
