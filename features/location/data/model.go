package data

import (
	"time"
	"github.com/roihan12/technical-test-kalbe/features/location"
	"gorm.io/gorm"
)

type Location struct {
	ID   uint           `gorm:"primaryKey;autoIncrement"`
	LocationName string         `gorm:"type:varchar(255)"`
	CreatedBy      string         `gorm:"type:varchar(255)"`
	UpdatedBy      string         `gorm:"type:varchar(255)"`
	CreatedAt      time.Time      `gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

// Convert LocationEntity to Location
func LocationEntityToLocation(data location.LocationEntity) Location {
	return Location{
		ID:   data.LocationID,
		LocationName: data.LocationName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      gorm.DeletedAt{Time: data.DeletedAt, Valid: !data.DeletedAt.IsZero()},
	}
}

// Convert Location to LocationEntity
func LocationToLocationEntity(data Location) location.LocationEntity {
	return location.LocationEntity{
		LocationID:   data.ID,
		LocationName: data.LocationName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt.Time,
	}
}

// Convert list of Location to list of LocationEntity
func ListLocationToLocationEntity(data []Location) []location.LocationEntity {
	var LocationEntities []location.LocationEntity
	for _, v := range data {
		LocationEntities = append(LocationEntities, LocationToLocationEntity(v))
	}
	return LocationEntities
}


type LocationEntity struct {
	LocationID   uint
	LocationName string
	CreatedBy      string
	UpdatedBy      string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}
