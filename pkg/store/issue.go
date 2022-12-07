package store

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssueStore struct {
	db        *gorm.DB
	issueNil  *models.Issue
	issuesNil *[]models.Issue
}

func NewIssueStore(db *gorm.DB) *IssueStore {
	return &IssueStore{
		db:        db,
		issueNil:  &models.Issue{},
		issuesNil: &[]models.Issue{},
	}
}

func (is *IssueStore) CreateIssue(issue models.Issue) (err error) {
	return is.db.Create(&issue).Error
}

func (is *IssueStore) GetIssues() (issues *[]models.Issue, err error) {
	if err = is.db.Preload("DepartmentRefer").Preload("UserID").Find(&issues).Error; err != nil {
		return is.issuesNil, err
	}

	return
}

func (is *IssueStore) GetIssueByID(id uuid.UUID) (issue *models.Issue, err error) {
	if err = is.db.Preload("DepartmentRefer").Preload("UserID").First(&issue, "issue_id = ?", id).Error; err != nil {
		return is.issueNil, err
	}

	return
}

func (is *IssueStore) UpdateIssue(id uuid.UUID, issueUpdates models.UpdateIssueRequest) (updatedIssue *models.Issue, err error) {
	if err = is.db.Model(&updatedIssue).Where("issue_id = ?", id).Updates(issueUpdates).Error; err != nil {
		return is.issueNil, err
	}

	return is.GetIssueByID(id)
}

func (is *IssueStore) DeleteIssue(id uuid.UUID) (deletedIssue *models.Issue, err error) {
	deletedIssue, err = is.GetIssueByID(id)

	if err != nil {
		return is.issueNil, err
	}

	if err = is.db.Delete(&deletedIssue, "issue_id = ?", id).Error; err != nil {
		return is.issueNil, err
	}

	return
}
