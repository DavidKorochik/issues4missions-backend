package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/token"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationHeaderType = "Bearer"
	authorizationPayloadKey = "x-auth-token"
)

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(authorizationHeaderKey)
		config := config.LoadEnvVariables("../../.")
		maker := token.NewJWTMaker(config.JWTSecret)

		if len(authHeader) == 0 {
			err := errors.New("Authorization header is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		fields := strings.Fields(authHeader)

		if len(fields) < 2 {
			err := errors.New("Invalid authorization header format")
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		if fields[0] != authorizationHeaderType {
			err := errors.New("Authorization header type is not provided")
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		payload, err := maker.VerifyJWTToken(fields[1])

		if err != nil {
			err := errors.New("Authorization token is not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, util.ErrorResponse(err))
			return
		}

		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
