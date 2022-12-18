package models

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" binding:"required"`
}

type GetDepartmentByNameURI struct {
	DepartmentName string `uri:"department_name" binding:"required"`
}

type UpdateDepartmentRequest struct {
	DepartmentName string `uri:"department_name" binding:"required"`
}

type DeleteDepartmentURI struct {
	DepartmentName string `uri:"department_name" binding:"required"`
}
