package repository

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
	UserNill  models.User
	UsersNill []models.User
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB:        db,
		UserNill:  models.User{},
		UsersNill: []models.User{},
	}
}
