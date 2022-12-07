package store

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserStore struct {
	db       *gorm.DB
	userNil  *models.User
	usersNil *[]models.User
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db:       db,
		userNil:  &models.User{},
		usersNil: &[]models.User{},
	}
}

func (us *UserStore) CreateUser(user *models.User) (err error) {
	return us.db.Create(&user).Error
}

func (us *UserStore) GetUsers() (users *[]models.User, err error) {
	if err := us.db.Preload("Issues").Preload("DepartmentRefer").Preload("RoleRefer").Find(&users).Error; err != nil {
		return us.usersNil, err
	}

	return
}

func (us *UserStore) GetByID(id uuid.UUID) (user *models.User, err error) {
	if err := us.db.Preload("Issues").Preload("DepartmentRefer").Preload("RoleRefer").First(&user, "user_id = ?", id).Error; err != nil {
		return us.userNil, err
	}

	return
}

func (us *UserStore) UpdateUser(id uuid.UUID, userUpdates models.UpdateUserRequest) (updatedUser *models.User, err error) {
	if err := us.db.Model(&models.User{}).Where("user_id = ?", id).Updates(userUpdates).Error; err != nil {
		return us.userNil, err
	}

	return us.GetByID(id)
}

func (us *UserStore) DeleteUser(id uuid.UUID) (deletedUser *models.User, err error) {
	deletedUser, err = us.GetByID(id)

	if err != nil {
		return us.userNil, err
	}

	if err := us.db.Delete(&models.User{}, "user_id = ?", id).Error; err != nil {
		return us.userNil, err
	}

	return
}
