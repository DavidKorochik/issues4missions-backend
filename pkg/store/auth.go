package store

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"gorm.io/gorm"
)

type AuthStore struct {
	db       *gorm.DB
	userNil  *models.User
	usersNil *[]models.User
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{
		db:       db,
		userNil:  &models.User{},
		usersNil: &[]models.User{},
	}
}

func (as *AuthStore) AuthUser(personalNumber string) (user *models.User, err error) {
	if err := as.db.First(&user, "personal_number = ?", personalNumber).Error; err != nil {
		return as.userNil, nil
	}

	return
}
