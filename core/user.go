package core

import (
	"github.com/satori/go.uuid"
)

// User defines a user model.
type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
