package models

import (
	"time"

	"github.com/google/uuid"
)

type Session struct {
	SessionID    uuid.UUID `json:"session_id" gorm:"primaryKey"`
	UserRefer    uuid.UUID `json:"user_id" gorm:"not null;column:user_id"`
	UserID       uuid.UUID `gorm:"foreignKey:UserRefer;references:UserID"`
	RefreshToken string    `json:"refresh_token" gorm:"not null"`
	UserAgent    string    `json:"user_agent" gorm:"not null"`
	ClientIP     string    `json:"client_ip" gorm:"not null"`
	IsBlocked    bool      `json:"is_blocked" gorm:"not null;default:false"`
	ExpiresAt    time.Time `json:"expires_at" gorm:"not null"`
	CreatedAt    time.Time `json:"created_at" gorm:"not null;default:now()"`
}
