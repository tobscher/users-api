package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tobscher/users-api/core"
)

// Users handles all users request
type Users struct {
	service core.UsersService
}

func (h *Users) index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	users, err := h.service.All()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getErrorResponse(err error) *core.ErrorResponse {
	return &core.ErrorResponse{
		Message: err.Error(),
	}
}
