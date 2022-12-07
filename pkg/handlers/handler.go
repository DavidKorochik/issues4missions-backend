package handlers

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/auth"
	"github.com/DavidKorochik/issues4missions-backend/pkg/issue"
	"github.com/DavidKorochik/issues4missions-backend/pkg/user"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationHeaderType = "Bearer"
	authorizationPayloadKey = "x-auth-token"
)

type Handler struct {
	userStore  user.Store
	issueStore issue.Store
	authStore  auth.Store
}

func NewHandler(us user.Store, is issue.Store, as auth.Store) *Handler {
	return &Handler{
		userStore:  us,
		issueStore: is,
		authStore:  as,
	}
}
