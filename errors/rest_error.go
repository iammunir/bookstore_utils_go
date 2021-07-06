package errors

import (
	"encoding/json"
	"errors"
	"net/http"
)

type RestError struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Causes  interface{} `json:"causes"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}

func NewInternalServerError(message string, err error) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Internal Server Error",
		Causes:  err.Error(),
	}
}

func NewRestErrorFromBytes(bytes []byte) (*RestError, error) {
	var apiErr RestError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return &apiErr, nil
}
