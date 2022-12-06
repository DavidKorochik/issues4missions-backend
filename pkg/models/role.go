package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	RoleName  string         `json:"department_name" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:now()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"not null"`
}
