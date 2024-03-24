package implementations

import (
	"auth-service/internal/config"
	"auth-service/internal/domain"
	"auth-service/internal/domain/entity"
	"auth-service/internal/service"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type TokenService struct {
	jwtSecret             string
	expirationTimeInHours int
}

func NewTokenService(config *config.Config) domain.TokenService {
	return &TokenService{
		jwtSecret:             config.JwtSecret,
		expirationTimeInHours: config.TokenTimeToLiveInHours,
	}
}

func (j *TokenService) CreateToken(claims []entity.Claim) (string, error) {
	payload := jwt.MapClaims{}
	payload["exp"] = time.Now().Add(time.Hour * time.Duration(j.expirationTimeInHours)).Unix()
	for _, claim := range claims {
		payload[claim.Title] = claim.Value
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(j.jwtSecret))

	if err != nil {
		return "", err
	}

	return token, err
}

func (j *TokenService) ParseClaims(jwtToken string) ([]entity.Claim, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		return []byte{}, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, service.ErrorInvalidToken
	}
	entitiesClaims := make([]entity.Claim, len(claims))
	for title, value := range claims {
		entitiesClaims = append(entitiesClaims, entity.Claim{
			Title: title,
			Value: value.(string),
		})
	}

	return entitiesClaims, nil
}
