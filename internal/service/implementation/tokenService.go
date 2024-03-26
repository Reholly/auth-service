package implementation

import (
	"auth-service/config"
	"auth-service/internal/domain/entity"
	"auth-service/internal/domain/helpers"
	service "auth-service/internal/domain/service"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"time"
)

type TokenService struct {
	jwtSecret             string
	expirationTimeInHours int
}

func NewTokenService(config *config.Config) service.TokenService {
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
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, service.ErrorInvalidToken
		}
		return []byte(j.jwtSecret), nil
	})

	if err != nil {
		return nil, errors.Wrap(service.ErrorInvalidToken, err.Error())
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, service.ErrorInvalidToken
	}
	entitiesClaims := make([]entity.Claim, 0, len(claims))
	for title, value := range claims {
		if title == "" || value == "" {
			continue
		}
		entitiesClaims = append(entitiesClaims, entity.Claim{
			Title: title,
			Value: value,
		})
	}

	return entitiesClaims, nil
}

func (j *TokenService) IsModerator(claims []entity.Claim) bool {
	for _, value := range claims {
		if value == helpers.ModeratorRole {
			return true
		}
	}

	return false
}

func (j *TokenService) IsAdmin(claims []entity.Claim) bool {
	for _, value := range claims {
		if value == helpers.AdminRole {
			return true
		}
	}

	return false
}

func (j *TokenService) IsInAdministration(claims []entity.Claim) bool {
	for _, value := range claims {
		if value == helpers.ModeratorRole || value == helpers.AdminRole {
			return true
		}
	}

	return false
}
