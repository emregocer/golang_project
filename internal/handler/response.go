package handler

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Data interface{} `json:"data"`
}

func NewAPIResponse(data interface{}) *APIResponse {
	return &APIResponse{Data: data}
}

func Message(status int, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func JsonSuccessResponse(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}

func JsonErrorResponse(w http.ResponseWriter, code int, msg string) error {
	response, err := json.Marshal(map[string]string{"message": msg})
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)

	return nil
}
