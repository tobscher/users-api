package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"

	"github.com/tobscher/users-api/core"
)

// users handles all users request
type users struct {
	service core.UsersService
}

func (h *users) index(w http.ResponseWriter, r *http.Request) {
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

func (h *users) show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	id, err := getID(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	user, err := h.service.Show(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(getErrorResponse(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func getErrorResponse(err error) *core.ErrorResponse {
	return &core.ErrorResponse{
		Message: err.Error(),
	}
}

func getID(r *http.Request) (uuid.UUID, error) {
	vars := mux.Vars(r)
	id := vars["id"]

	var userID uuid.UUID
	var err error
	if userID, err = uuid.FromString(id); err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
