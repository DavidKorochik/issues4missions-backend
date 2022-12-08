package handlers

import (
	"net/http"
	"time"

	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/token"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) AuthUser(c *gin.Context) {
	config := config.LoadEnvVariables("../../.")

	var req models.AuthUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user, err := h.authStore.AuthUser(req.PersonalNumber)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	err = util.CheckSecretCode(req.SecretCode, user.SecretCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	secretKey := config.JWTSecret
	maker := token.NewJWTMaker(secretKey)

	// TODO: Get in return the payload struct
	token, _, err := maker.GenerateJWT(user.UserID, time.Hour)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	refreshToken, _, err := maker.GenerateJWT(user.UserID, time.Hour*24)

	c.JSON(http.StatusOK, gin.H{"token": token, "refresh_token": refreshToken})
}
