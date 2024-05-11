package service

import (
	"auth-service/internal/domain/model"
	"errors"
)

var (
	ErrInvalidToken = errors.New("invalid token")
)

type TokenService interface {
	CreateToken(claims []model.Claim) (string, error)
	ParseClaims(jwtToken string) ([]model.Claim, error)
}
