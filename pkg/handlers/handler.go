package handlers

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/issue"
	"github.com/DavidKorochik/issues4missions-backend/pkg/user"
)

type Handler struct {
	userStore  user.Store
	issueStore issue.Store
}

func NewHandlerRepo(us user.Store, is issue.Store) *Handler {
	return &Handler{
		userStore:  us,
		issueStore: is,
	}
}
