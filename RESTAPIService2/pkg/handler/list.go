package handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) createList(w http.ResponseWriter, r *http.Request) {
	// ОШИБКИ ЧЕРЕЗ ЛОГГЕР
	id := r.Context().Value(userCtx)

	idStr, ok := id.(string)
	if !ok {
		http.Error(w, "User id is not a string", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse := map[string]interface{}{
		"id": idStr,
	}
	if err := json.NewEncoder(w).Encode(jsonResponse); err != nil {
		logrus.Println("Error encoding JSON response:", err)
		newErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *Handler) getAllLists(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) getListById(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) updateList(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) deleteList(w http.ResponseWriter, r *http.Request) {

}
