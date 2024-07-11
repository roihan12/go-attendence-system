package data

import (
	"time"

	department "github.com/roihan12/technical-test-kalbe/features/department/data"
	"github.com/roihan12/technical-test-kalbe/features/employee"
	position "github.com/roihan12/technical-test-kalbe/features/position/data"
	"gorm.io/gorm"
)

type Employee struct {
	ID   uint           `gorm:"primaryKey;autoIncrement"`
	EmployeeCode string         `gorm:"type:varchar(255)"`
	EmployeeName string         `gorm:"type:varchar(255)"`
	Password     string         `gorm:"type:varchar(255)"`
	DepartmentID uint           `gorm:"type:integer;not null"`
	PositionID   uint           `gorm:"type:integer;not null"`
	Superior     uint           `gorm:"type:integer"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	CreatedBy    string         `gorm:"type:varchar(255)"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	UpdatedBy    string         `gorm:"type:varchar(255)"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Department department.Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Position   position.Position     `gorm:"foreignKey:PositionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func ToEmployeeEntity(data Employee) employee.EmployeeEntity {
	return employee.EmployeeEntity{
		EmployeeID:   data.ID,
		EmployeeCode: data.EmployeeCode,
		EmployeeName: data.EmployeeName,
		Password:     data.Password,
		DepartmentID: data.DepartmentID,
		PositionID:   data.PositionID,
		Superior:     data.Superior,
		CreatedAt:    data.CreatedAt,
		CreatedBy:    data.CreatedBy,
		UpdatedAt:    data.UpdatedAt,
		UpdatedBy:    data.UpdatedBy,
		DeletedAt:    data.DeletedAt.Time,
	}
}

func EmployeeEntityToEmployee(employeeEntity employee.EmployeeEntity) Employee {
	return Employee{
		ID:   employeeEntity.EmployeeID,
		EmployeeCode: employeeEntity.EmployeeCode,
		EmployeeName: employeeEntity.EmployeeName,
		Password:     employeeEntity.Password,
		DepartmentID: employeeEntity.DepartmentID,
		PositionID:   employeeEntity.PositionID,
		Superior:     employeeEntity.Superior,
		CreatedAt:    employeeEntity.CreatedAt,
		CreatedBy:    employeeEntity.CreatedBy,
		UpdatedAt:    employeeEntity.UpdatedAt,
		UpdatedBy:    employeeEntity.UpdatedBy,
		DeletedAt:    gorm.DeletedAt{Time: employeeEntity.DeletedAt, Valid: !employeeEntity.DeletedAt.IsZero()},
	}
}

func ListEmployeeToEmployeeEntity(data []Employee) []employee.EmployeeEntity {
	var employeeEntities []employee.EmployeeEntity
	for _, v := range data {
		employeeEntities = append(employeeEntities, ToEmployeeEntity(v))
	}
	return employeeEntities
}
