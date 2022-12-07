package handlers

import (
	"net/http"

	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	hashedSecretCode, err := hashSecretCode(req.SecretCode)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	newUser := models.User{FirstName: req.FirstName, LastName: req.LastName, PersonalNumber: req.PersonalNumber, SecretCode: hashedSecretCode, PhoneNumber: req.PhoneNumber}

	if err := h.userStore.CreateUser(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userStore.GetUsers()

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	var uri models.GetUserByIDUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	user, err := h.userStore.GetUserByID(uri.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var uri models.UpdateUserUri
	var req models.UpdateUserRequest

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	updatedUser, err := h.userStore.UpdateUser(uri.ID, req)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var uri models.DeleteUserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	deletedUser, err := h.userStore.DeleteUser(uri.ID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return
		}

		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, deletedUser)
}

func hashSecretCode(secretCode string) (hashSecretCode string, err error) {
	return util.HashSecretCode(secretCode)
}
