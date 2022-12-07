package server

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	r  *gin.Engine
	db *gorm.DB
}

func NewServer(db *db.Database) *Server {
	r := gin.Default()

	return &Server{
		r: r,
	}
}

func (s *Server) StartServer(address string) error {
	return s.r.Run(address)
}
