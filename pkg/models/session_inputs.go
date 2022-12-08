package models

import (
	"time"

	"github.com/google/uuid"
)

type CreateSessionRequest struct {
	SessionID    uuid.UUID `json:"session_id" binding:"required,min=1"`
	UserID       uuid.UUID `json:"user_id" binding:"required,min=1"`
	RefreshToken string    `json:"refresh_token" binding:"required"`
	UserAgent    string    `json:"user_agent" binding:"required"`
	ClientIP     string    `json:"client_ip" binding:"required"`
	IsBlocked    bool      `json:"is_blocked" binding:"required,oneof=true false"`
	ExpiresAt    time.Time `json:"expires_at" binding:"required"`
}
