package models

type UpdateIssueRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required,oneof=open closed"`
	IsCompleted bool   `json:"is_completed" binding:"required,oneof=true false"`
}
