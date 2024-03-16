package infrastructure

import (
	"auth-service/internal/domain"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtManager struct {
	jwtSecret             string
	expirationTimeInHours int
}

func NewJwtManager(jwtSecret string, expirationTimeInHours int) *JwtManager {
	return &JwtManager{
		jwtSecret:             jwtSecret,
		expirationTimeInHours: expirationTimeInHours,
	}
}

func (j *JwtManager) CreateToken(claims []domain.Claim) (string, error) {
	payload := jwt.MapClaims{}
	payload["exp"] = time.Now().Add(time.Hour * time.Duration(j.expirationTimeInHours)).Unix()
	for _, claim := range claims {
		payload[claim.Title] = claim.Value
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString(j.jwtSecret)

	if err != nil {
		return "", err
	}

	return token, err
}

func (j *JwtManager) ParseClaims(jwtToken string) ([]domain.Claim, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrorInvalidToken
	}
	entitiesClaims := make([]domain.Claim, len(claims))
	for title, value := range claims {
		entitiesClaims = append(entitiesClaims, domain.Claim{
			Title: title,
			Value: value.(string),
		})
	}
	return entitiesClaims, nil
}
