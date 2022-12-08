package store

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionStore struct {
	db         *gorm.DB
	sessionNil *models.Session
}

func NewSessionStore(db *gorm.DB) *SessionStore {
	return &SessionStore{
		db:         db,
		sessionNil: &models.Session{},
	}
}

// I won't call it ss because in jewish lol
func (s *SessionStore) CreateSession(session *models.Session) error {
	return s.db.Create(&session).Error
}

func (s *SessionStore) GetSessionByID(id uuid.UUID) (session *models.Session, err error) {
	if err = s.db.First(&session, "session_id = ?", id).Error; err != nil {
		return s.sessionNil, err
	}

	return
}
