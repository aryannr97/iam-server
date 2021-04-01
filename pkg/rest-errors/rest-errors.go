package resterrors

import (
	"net/http"

	"github.com/aryannr97/iam-server/models"
)

// NewBadRequest return common error object
func NewBadRequest(msg string) models.RestError {
	return models.RestError{
		StatusCode: http.StatusBadRequest,
		Status:     "BAD_REQUEST",
		ErrMessage: msg,
	}
}

// NewInternalServerError return common error object
func NewInternalServerError(msg string) models.RestError {
	return models.RestError{
		StatusCode: http.StatusInternalServerError,
		Status:     "INTERNAL_SERVER_ERROR",
		ErrMessage: msg,
	}
}
