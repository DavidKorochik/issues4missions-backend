package session

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
)

type Store interface {
	CreateSession(session *models.Session) error
	GetSessionByID(id uuid.UUID) (session *models.Session, err error)
}
