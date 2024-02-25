package handler

import (
	"context"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "UserId"
)

func (h *Handler) userIdentity(w http.ResponseWriter, r *http.Request) {
	header := r.Header.Get(authorizationHeader)
	if header == "" {
		newErrorResponse(w, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(w, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := context.WithValue(r.Context(), userCtx, userId)
	r = r.WithContext(ctx)
}
