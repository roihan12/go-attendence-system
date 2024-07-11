package data

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/department"
	"gorm.io/gorm"
)

type Department struct {
	ID             uint           `gorm:"primaryKey;autoIncrement"`
	DepartmentName string         `gorm:"type:varchar(255)"`
	CreatedBy      string         `gorm:"type:varchar(255)"`
	UpdatedBy      string         `gorm:"type:varchar(255)"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Convert DepartmentEntity to Department
func DepartmentEntityToDepartment(data department.DepartmentEntity) Department {
	return Department{
		ID:             data.DepartmentID,
		DepartmentName: data.DepartmentName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      gorm.DeletedAt{Time: data.DeletedAt, Valid: !data.DeletedAt.IsZero()},
	}
}

// Convert Department to DepartmentEntity
func DepartmentToDepartmentEntity(data Department) department.DepartmentEntity {
	return department.DepartmentEntity{
		DepartmentID:   data.ID,
		DepartmentName: data.DepartmentName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt.Time,
	}
}

// Convert list of Department to list of DepartmentEntity
func ListDepartmentToDepartmentEntity(data []Department) []department.DepartmentEntity {
	var departmentEntities []department.DepartmentEntity
	for _, v := range data {
		departmentEntities = append(departmentEntities, DepartmentToDepartmentEntity(v))
	}
	return departmentEntities
}

type DepartmentEntity struct {
	DepartmentID   uint
	DepartmentName string
	CreatedBy      string
	UpdatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
