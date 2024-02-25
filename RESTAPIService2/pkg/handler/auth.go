package handler

import (
	todo "RESTAPIService2"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input todo.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]interface{}{
		"id": id,
	}
	if err = json.NewEncoder(w).Encode(jsonResponse); err != nil {
		// Если возникает ошибка при кодировании JSON, можно просто логировать ее
		logrus.Println("Error encoding JSON response:", err)
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		newErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]interface{}{
		"token": token,
	}
	if err = json.NewEncoder(w).Encode(jsonResponse); err != nil {
		// Если возникает ошибка при кодировании JSON, можно просто логировать ее
		logrus.Println("Error encoding JSON response:", err)
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}
