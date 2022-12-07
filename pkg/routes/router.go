package routes

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/handlers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	router  *gin.Engine
	handler *handlers.Handler
}

func NewRouter(handler *handlers.Handler) *Routes {
	return &Routes{
		router:  gin.Default(),
		handler: handler,
	}
}
