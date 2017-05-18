package handler

import (
	"encoding/json"
	"net/http"
)

func RespondWithCustomErrors(w http.ResponseWriter, errorBody interface{}, statusCode int) error {
	w.WriteHeader(statusCode)
	body, err := json.Marshal(errorBody)

	if err != nil {
		return err
	}
	w.Write(body)
	return nil
}
