package data

import (
	"time"

	departmentModel "github.com/roihan12/technical-test-kalbe/features/department/data"
	"github.com/roihan12/technical-test-kalbe/features/position"
	"gorm.io/gorm"
)

// Position represents the position model
type Position struct {
	ID   uint           `gorm:"primaryKey;autoIncrement"`
	DepartmentID uint           `gorm:"type:integer;not null"`
	PositionName string         `gorm:"type:varchar(255)"`
	CreatedBy    string         `gorm:"type:varchar(255)"`
	UpdatedBy    string         `gorm:"type:varchar(255)"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`

	Department departmentModel.Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}

// Convert PositionEntity to Position
func PositionEntityToPosition(data position.PositionEntity) Position {
	return Position{
		ID:   data.PositionID,
		DepartmentID: data.DepartmentID,
		PositionName: data.PositionName,
		CreatedBy:    data.CreatedBy,
		UpdatedBy:    data.UpdatedBy,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    gorm.DeletedAt{Time: data.DeletedAt, Valid: !data.DeletedAt.IsZero()},
	}
}

// Convert Position to PositionEntity
func PositionToPositionEntity(data Position) position.PositionEntity {
	return position.PositionEntity{
		PositionID:   data.ID,
		DepartmentID: data.DepartmentID,
		PositionName: data.PositionName,
		CreatedBy:    data.CreatedBy,
		UpdatedBy:    data.UpdatedBy,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt.Time,
	}
}

// Convert list of Position to list of PositionEntity
func ListPositionToPositionEntity(data []Position) []position.PositionEntity {
	var positionEntities []position.PositionEntity
	for _, v := range data {
		positionEntities = append(positionEntities, PositionToPositionEntity(v))
	}
	return positionEntities
}

type PositionEntity struct {
	PositionID   uint
	DepartmentID uint
	PositionName string
	CreatedBy    string
	UpdatedBy    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
