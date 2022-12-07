package controllers

import (
	"net/http"

	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/services"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserStore {
	return &UserController{
		service: service,
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.service.GetUsers()

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	newUser := models.User{FirstName: req.FirstName, LastName: req.LastName, PersonalNumber: req.PersonalNumber, SecretCode: req.SecretCode, PhoneNumber: req.PhoneNumber}
	createdUser, err := uc.service.CreateUser(newUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var req models.UpdateUserRequest
	var uri models.UpdateUserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	updatedUser, err := uc.service.UpdateUser(uri.ID, req)

	if err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	var uri models.DeleteUserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if err := uc.service.DeleteUser(uri.ID); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted!"})
}
