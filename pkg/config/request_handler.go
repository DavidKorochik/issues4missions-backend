package config

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Gin *gin.Engine
}

func NewRequestHandler() *RequestHandler {
	return &RequestHandler{
		Gin: gin.Default(),
	}
}
