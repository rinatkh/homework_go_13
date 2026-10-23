package user

import "errors"

var (
	ErrTODO               = errors.New("TODO: implement authentication use case")
	ErrInvalidEmail       = errors.New("email is invalid")
	ErrWeakPassword       = errors.New("password must contain at least 8 characters")
	ErrAlreadyExists      = errors.New("user already exists")
	ErrNotFound           = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
