package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (maker *JWTMaker) GenerateJWT(userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload := NewPayload(userID, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	accessToken, err := jwtToken.SignedString([]byte(maker.secretKey))

	if err != nil {
		return "", payload, err
	}

	return accessToken, payload, nil
}

func (maker *JWTMaker) VerifyJWTToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("The token is invalid")
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		return nil, err
	}

	return jwtToken.Claims.(*Payload), nil
}
