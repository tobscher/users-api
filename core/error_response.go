package core

// ErrorResponse is used to return an error from an HTTP response.
type ErrorResponse struct {
	Message string `json:"message"`
}
