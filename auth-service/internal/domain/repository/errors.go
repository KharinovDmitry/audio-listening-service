package repository

import "errors"

var (
	ErrAlreadyExists = errors.New("user already exists")
	ErrNotFound      = errors.New("not found")
)
