package department

import "github.com/DavidKorochik/issues4missions-backend/pkg/models"

type Store interface {
	CreateDepartment(department models.Department) (err error)
	GetDepartments() (departments *[]models.Department, err error)
	GetDepartmentByName(departmentName string) (department *models.Department, err error)
	UpdateDepartment(departmentUpdates models.Department) (updatedDepartment *models.Department, err error)
	DeleteDepartment(departmentName string) (deletedDepartment *models.Department, err error)
}
