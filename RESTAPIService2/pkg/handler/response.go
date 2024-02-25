package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	logrus.Error(message)

	errRes := ErrorResponse{Message: message}

	jsonErrRes, err := json.Marshal(errRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonErrRes)
}
