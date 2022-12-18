package store

import (
	"github.com/DavidKorochik/issues4missions-backend/pkg/models"
	"gorm.io/gorm"
)

type DepartmentStore struct {
	db             *gorm.DB
	DepartmentNil  *models.Department
	DepartmentsNil *[]models.Department
}

func NewDepartmentStore(db *gorm.DB) *DepartmentStore {
	return &DepartmentStore{
		db:             db,
		DepartmentNil:  &models.Department{},
		DepartmentsNil: &[]models.Department{},
	}
}

func (ds *DepartmentStore) CreateDepartment(department models.Department) (err error) {
	return ds.db.Create(&department).Error
}

func (ds *DepartmentStore) GetDepartments() (departments *[]models.Department, err error) {
	if err := ds.db.Find(&departments).Error; err != nil {
		return ds.DepartmentsNil, err
	}

	return
}

func (ds *DepartmentStore) GetDepartmentByName(departmentName string) (department *models.Department, err error) {
	if err := ds.db.First(&department, "department_name = ?", departmentName).Error; err != nil {
		return ds.DepartmentNil, err
	}

	return
}

func (ds *DepartmentStore) UpdateDepartment(departmentUpdates models.Department) (updatedDepartment *models.Department, err error) {
	if err := ds.db.Where("department_name = ?", departmentUpdates.DepartmentName).Updates(&departmentUpdates).Error; err != nil {
		return ds.DepartmentNil, err
	}

	return ds.GetDepartmentByName(departmentUpdates.DepartmentName)
}

func (ds *DepartmentStore) DeleteDepartment(departmentName string) (deletedDepartment *models.Department, err error) {
	deletedDepartment, err = ds.GetDepartmentByName(departmentName)

	if err != nil {
		return ds.DepartmentNil, err
	}

	if err := ds.db.Delete(&models.Department{}, "department_name = ?", departmentName).Error; err != nil {
		return ds.DepartmentNil, err
	}

	return
}
