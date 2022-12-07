package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID         uuid.UUID      `json:"user_id" gorm:"primaryKey"`
	FirstName      string         `json:"first_name" gorm:"not null"`
	LastName       string         `json:"last_name" gorm:"not null"`
	PersonalNumber string         `json:"personal_number" gorm:"not null;unique"`
	SecretCode     string         `json:"secret_code" gorm:"not null;unique"`
	PhoneNumber    string         `json:"phone_number" gorm:"not null;unique"`
	RoleRefer      string         `json:"role_name" gorm:"column:role_name;uniqueIndex"`
	Role           Role           `json:"-" gorm:"foreignKey:RoleRefer"`
	Issues         []Issue        `gorm:"foreignKey:UserID"`
	CreatedAt      time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"not null;default:now()"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// We want to use NewRandom, because it returns an error and we want better error handling.
	u.UserID, err = uuid.NewRandom()
	return
}
