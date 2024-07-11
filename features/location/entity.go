package location

import (
	"time"
)

type LocationEntity struct {
	LocationID   uint
	LocationName string
	CreatedBy       string
	UpdatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type LocationService interface {
	GetAll() ([]LocationEntity, error)
	GetById(id uint) (LocationEntity, error)
	Create(createLocation LocationEntity,) (LocationEntity, error)
	Update(updateLocation LocationEntity, id uint) (LocationEntity, error)
	Delete(id uint) error
}

type LocationData interface {
	GetAll() ([]LocationEntity, error)
	GetById(id uint) (LocationEntity, error)
	Create(createLocation LocationEntity) (LocationEntity, error)
	Update(updateLocation LocationEntity, id uint) (LocationEntity, error)
	Delete(id uint) error
}
