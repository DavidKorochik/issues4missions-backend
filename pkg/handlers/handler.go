package handlers

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/issue"
	"github.com/DavidKorochik/issues4missions-backend/pkg/user"
)

type HandlerRepo struct {
	userStore  user.Store
	issueStore issue.Store
}

func NewHandlerRepo(us user.Store, is issue.Store) *HandlerRepo {
	return &HandlerRepo{
		userStore:  us,
		issueStore: is,
	}
}
