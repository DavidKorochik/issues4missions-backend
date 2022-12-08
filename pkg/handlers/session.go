package handlers

import (
	"net/http"
	"time"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/token"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RenewAccessToken(c *gin.Context) {
	config := config.LoadEnvVariables("../../.")

	var req models.CreateSessionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	secretKey := config.JWTSecret
	maker := token.NewJWTMaker(secretKey)

	payload, err := maker.VerifyJWTToken(req.RefreshToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	session, err := h.sessionStore.GetSessionByID(payload.ID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	if session.IsBlocked {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	if session.UserID != req.UserID {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	if time.Now().After(payload.ExpiredAt) {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	accessToken, _, err := maker.GenerateJWT(payload.UserID, time.Hour*24)

	if err != nil {
		c.JSON(http.StatusUnauthorized, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"access_token": accessToken})
}
