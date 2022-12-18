package models

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" binding:"required"`
}
