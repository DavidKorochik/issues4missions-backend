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

func NewUserController(service services.UserService) *UserController {
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

}
