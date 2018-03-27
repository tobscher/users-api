package core

import uuid "github.com/satori/go.uuid"

// UsersService defines an interface to create, read, update and destroy user related
// data.
type UsersService interface {
	All() ([]*User, error)
	Show(id uuid.UUID) (*User, error)
}
