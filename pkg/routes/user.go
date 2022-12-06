package routes

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/config"
	"github.com/DavidKorochik/issues4missions-backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
	requestHandler config.RequestHandler
}

func NewUserRoutes(userController controllers.UserController, requestHandler config.RequestHandler) *UserRoutes {
	return &UserRoutes{
		userController: userController,
		requestHandler: requestHandler,
	}
}

func (ur *UserRoutes) SetupUserRoutes(r *gin.Engine) {

}
