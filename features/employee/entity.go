package employee

import (
	"time"
)

type EmployeeEntity struct {
	EmployeeID   uint
	EmployeeCode string
	EmployeeName string
	Password     string
	DepartmentID uint
	PositionID   uint
	Superior     uint
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type EmployeeService interface {
	Login(email, password string) (string, error)
	Register(newEmployee EmployeeEntity) (EmployeeEntity, error)
	Profile(employeeID uint) (EmployeeEntity, error)
	Update(employeeID uint, updateData EmployeeEntity) (EmployeeEntity, error)
	Delete(employeeID uint) error
}

type EmployeeData interface {
	Login(employeeName string) (EmployeeEntity, error)
	Register(newEmployee EmployeeEntity) (EmployeeEntity, error)
	Profile(employeeID uint) (EmployeeEntity, error)
	Update(employeeID uint, updateData EmployeeEntity) (EmployeeEntity, error)
	GetLatestEmployee() (EmployeeEntity, error)
	Delete(employeeID uint) error
}
