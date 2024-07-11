package department

import (
	"time"
)

type DepartmentEntity struct {
	DepartmentID   uint
	DepartmentName string
	CreatedBy       string
	UpdatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type DepartmentService interface {
	GetAll() ([]DepartmentEntity, error)
	GetById(id uint) (DepartmentEntity, error)
	Create(createDepartment DepartmentEntity,) (DepartmentEntity, error)
	Update(updateDepartment DepartmentEntity, id uint) (DepartmentEntity, error)
	Delete(id uint) error
}

type DepartmentData interface {
	GetAll() ([]DepartmentEntity, error)
	GetById(id uint) (DepartmentEntity, error)
	Create(createDepartment DepartmentEntity) (DepartmentEntity, error)
	Update(updateDepartment DepartmentEntity, id uint) (DepartmentEntity, error)
	Delete(id uint) error
}
