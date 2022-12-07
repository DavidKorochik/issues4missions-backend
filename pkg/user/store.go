package user

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
)

type Store interface {
	CreateUser(user *models.User) (err error)
	GetUsers() (users *[]models.User, err error)
	GetByID(id uuid.UUID) (user *models.User, err error)
	UpdateUser(id uuid.UUID, userUpdates models.UpdateUserRequest) (updatedUser *models.User, err error)
	DeleteUser(id uuid.UUID) (deletedUser *models.User, err error)
}
