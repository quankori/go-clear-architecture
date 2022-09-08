package auth

import "errors"

var (
	ErrUserNotFound       = errors.New("USER_NOT_FOUND")
	ErrWrongPassword      = errors.New("WRONG_PASSWORD")
	ErrUserExisted        = errors.New("USER_EXISTED")
	ErrInvalidAccessToken = errors.New("INVALID_ACCESS_TOKEN")
)
