package core

// UsersService defines an interface to create, read, update and destroy user related
// data.
type UsersService interface {
	All() ([]*User, error)
}
