package errors

import(
	"net/http"
)

type ErrorResponse struct {
	Message string
	Status int
	Error string
}

func NewInternalServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "Internal Server Error",
	}
}

func NewBadRequestError(message string) *ErrorResponse {
	return &ErrorResponse{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "Bad Request",
	}
}