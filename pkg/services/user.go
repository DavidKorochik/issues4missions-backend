package services

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/DavidKorochik/issues4missions-backend/pkg/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (us *UserService) GetUsers() (users []models.User, err error) {
	return users, us.repository.Find(&users).Error
}

func (us *UserService) CreateUser(userToCreate models.User) (err error) {
	return us.repository.Create(&userToCreate).Error
}

func (us *UserService) UpdateUser(id uuid.UUID, userUpdates models.User) (err error) {
	return us.repository.Model(&userUpdates).Where("user_id = ?", id).Updates(&userUpdates).Error
}

func (us *UserService) DeletedUser(id uuid.UUID) (err error) {
	return us.repository.Delete(&models.User{}, "user_id = ?", id).Error
}
