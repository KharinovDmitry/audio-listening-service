package service

import (
	"context"
	"errors"
)

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUnknownRole        = errors.New("unknown role")
)

type AuthService interface {
	SignUp(ctx context.Context, login string, password string, role string) (err error)
	SignIn(ctx context.Context, login string, password string) (token string, err error)
}
