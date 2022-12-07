package models

import "github.com/google/uuid"

type CreateIssueRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type GetIssueByIDUri struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

type UpdateIssueRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=open closed"`
	IsCompleted bool   `json:"is_completed" binding:"required,oneof=true false"`
}

type UpdateIssueUri struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}

type DeleteIssueUri struct {
	ID uuid.UUID `uri:"id" binding:"required,min=1"`
}
