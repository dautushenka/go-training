package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteResponseError(w http.ResponseWriter, message string, statusCode int) {
	WriteJSONResponse(w, ErrorResponse{Code: statusCode, Message: message}, statusCode)
}

func WriteJSONResponse(w http.ResponseWriter, obj interface{}, statusCode int) {
	body, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := fmt.Fprintln(w, string(body)); err != nil {
		panic(err)
	}

}
