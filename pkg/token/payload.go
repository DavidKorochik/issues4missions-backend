package token

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	IssuesAt  time.Time `json:"issues_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(userID uuid.UUID, duration time.Duration) *Payload {
	return &Payload{
		ID:        uuid.New(),
		UserID:    userID,
		IssuesAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return fmt.Errorf("The time %d of the token expiration has passed", p.ExpiredAt)
	}

	return nil
}
