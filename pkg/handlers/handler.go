package handlers

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/auth"
	"github.com/DavidKorochik/issues4missions-backend/pkg/department"
	"github.com/DavidKorochik/issues4missions-backend/pkg/issue"
	"github.com/DavidKorochik/issues4missions-backend/pkg/session"
	"github.com/DavidKorochik/issues4missions-backend/pkg/user"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationHeaderType = "Bearer"
	authorizationPayloadKey = "x-auth-token"
)

type Handler struct {
	userStore       user.Store
	issueStore      issue.Store
	departmentStore department.Store
	authStore       auth.Store
	sessionStore    session.Store
}

func NewHandler(us user.Store, is issue.Store, ds department.Store, as auth.Store, s session.Store) *Handler {
	return &Handler{
		userStore:    us,
		issueStore:   is,
		authStore:    as,
		sessionStore: s,
	}
}
