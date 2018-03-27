package core

// ErrorResponse is used to return an error from an HTTP response.
type ErrorResponse struct {
	Message string `json:"message"`
}

// NewErrorResponse returns a new error response from an error.
func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Message: err.Error(),
	}
}
