package implementations

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/domain/service"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Token struct {
	jwtSecret string
	TokenTTL  time.Duration
}

func NewToken(jwtSecret string, ttl time.Duration) *Token {
	return &Token{
		jwtSecret: jwtSecret,
		TokenTTL:  ttl,
	}
}

func (t *Token) CreateToken(claims []model.Claim) (string, error) {
	payload := jwt.MapClaims{}
	payload[model.ExpiredClaimTitle] = time.Now().
		Add(t.TokenTTL).
		Unix()

	for _, claim := range claims {
		payload[claim.Title] = claim.Value
	}

	token, err := jwt.
		NewWithClaims(jwt.SigningMethodHS256, payload).
		SignedString([]byte(t.jwtSecret))

	return token, err
}

func (t *Token) ParseClaims(jwtToken string) ([]model.Claim, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, service.ErrInvalidToken
		}
		return []byte(t.jwtSecret), nil
	})

	if err != nil {
		return nil, service.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, service.ErrInvalidToken
	}
	entitiesClaims := make([]model.Claim, 0, len(claims))
	for title, value := range claims {
		if title == "" || value == "" {
			continue
		}
		entitiesClaims = append(entitiesClaims, model.Claim{
			Title: title,
			Value: value,
		})
	}

	return entitiesClaims, nil
}
