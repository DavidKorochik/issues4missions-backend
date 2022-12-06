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

func (us *UserService) GetUsers() ([]models.User, error) {
	var users []models.User

	if err := us.repository.Find(&users).Error; err != nil {
		return us.repository.UsersNill, err
	}

	return users, nil
}

func (us *UserService) CreateUser(newUser models.User) (models.User, error) {
	if err := us.repository.Create(&newUser).Error; err != nil {
		return us.repository.UserNill, err
	}

	return newUser, nil
}

func (us *UserService) UpdateUser(id uuid.UUID, userUpdates models.User) (models.User, error) {
	if err := us.repository.Model(&models.User{}).Where("user_id = ?", id).Updates(userUpdates).Error; err != nil {
		return us.repository.UserNill, err
	}

	udpatedUser, err := us.FindUserByID(id)

	if err != nil {
		return us.repository.UserNill, err
	}

	return udpatedUser, nil
}

func (us *UserService) DeletedUser(id uuid.UUID) error {
	if err := us.repository.Delete(&models.User{}, "user_id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (us *UserService) FindUserByID(id uuid.UUID) (models.User, error) {
	var user models.User

	if err := us.repository.First(&user, "user_id = ?", id).Error; err != nil {
		return us.repository.UserNill, err
	}

	return user, nil
}
