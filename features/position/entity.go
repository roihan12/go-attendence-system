package position

import (
	"time"
)

type PositionEntity struct {
	PositionID     uint
	DepartmentID  uint
	PositionName string
	CreatedBy       string
	UpdatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type PositionService interface {
	GetAll() ([]PositionEntity, error)
	GetById(id uint) (PositionEntity, error)
	Create(createPosition PositionEntity,) (PositionEntity, error)
	Update(updatePosition PositionEntity, id uint) (PositionEntity, error)
	Delete(id uint) error
}

type PositionData interface {
	GetAll() ([]PositionEntity, error)
	GetById(id uint) (PositionEntity, error)
	Create(createPosition PositionEntity) (PositionEntity, error)
	Update(updatePosition PositionEntity, id uint) (PositionEntity, error)
	Delete(id uint) error
}


