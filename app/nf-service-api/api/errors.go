package api

import (
	"errors"
	"net/http"
)

var (
	ErrInvalidArgument  = errors.New("invalid argument")
	ErrNotFound         = errors.New("not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrUserAlreadyExist = errors.New("User with this email already exist")
)

func HTTPStatusCodeByError(err error) int {
	switch err.Error() {
	case ErrNotFound.Error():
		return http.StatusNotFound
	case ErrUnauthorized.Error():
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
