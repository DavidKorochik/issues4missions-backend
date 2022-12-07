package auth

import "github.com/DavidKorochik/issues4missions-backend/pkg/models"

type Store interface {
	AuthUser(personalNumber string) (user *models.User, err error)
}
