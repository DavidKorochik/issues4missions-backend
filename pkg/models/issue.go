package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Issue struct {
	IssueID         uuid.UUID `json:"issue_id" gorm:"primaryKey"`
	Title           string    `json:"title" gorm:"not null"`
	Description     string    `json:"description" gorm:"not null"`
	Status          string    `json:"status" gorm:"not null;default:open"`
	IsCompleted     bool      `json:"is_completed" gorm:"not null;default:false"`
	User            User
	UserID          uuid.UUID
	DepartmentRefer string         `json:"department_name" gorm:"column:department_name"`
	Department      Department     `json:"-" gorm:"foreignKey:DepartmentRefer"`
	CreatedAt       time.Time      `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"not null;default:now()"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at" gorm:"not null"`
}

func (i *Issue) BeforeCreate(tx *gorm.DB) (err error) {
	// We want to use NewRandom, because it returns an error and we want better error handling.
	i.IssueID, err = uuid.NewRandom()
	return
}
