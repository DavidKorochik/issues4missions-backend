package controllers

import (
	"net/http"

	"github.com/DavidKorochik/issues4missions-backend/pkg/services"
	"github.com/DavidKorochik/issues4missions-backend/pkg/util"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserController{
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
}
