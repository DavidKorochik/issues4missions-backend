package issue

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
)

type Store interface {
	CreateIssue(issue models.Issue) (err error)
	GetIssues() (issues *[]models.Issue, err error)
	GetByID(id uuid.UUID) (issue *models.Issue, err error)
	UpdateIssue(id uuid.UUID, issueUpdates models.UpdateIssueRequest) (updatedIssue *models.Issue, err error)
	DeleteIssue(id uuid.UUID) (deletedIssue *models.Issue, err error)
}
