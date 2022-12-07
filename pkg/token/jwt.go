package token

import (
	"fmt"
	"time"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJWT(userID uuid.UUID, duration time.Duration) (string, error) {
	config := config.LoadEnvVariables("../../")
	payload := NewPayload(userID, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.JWTSecret))
}

func VerifyJWTToken(token string) (*Payload, error) {
	config := config.LoadEnvVariables("../../")
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("The token is invalid")
		}

		return []byte(config.JWTSecret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		return nil, err
	}

	return jwtToken.Claims.(*Payload), nil
}
