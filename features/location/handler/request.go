package handler

import (
	"github.com/roihan12/technical-test-kalbe/features/location"
)

type CreateLocationRequest struct {
	LocationName string `json:"location_name" form:"location_name" binding:"required" example:"Finance"`
	CreatedBy      string `json:"created_by" form:"created_by" binding:"required" example:"Admin"`
}

type UpdateLocationRequest struct {
	LocationName string `json:"location_name" form:"location_name" binding:"required" example:"Finance"`
	UpdatedBy      string `json:"updated_by" form:"updated_by" binding:"required" example:"Admin"`
}

func CreateLocationRequestToLocationEntity(request *CreateLocationRequest) location.LocationEntity {
	return location.LocationEntity{
		LocationName: request.LocationName,
		CreatedBy:      request.CreatedBy,
	}
}

func UpdateLocationRequestToLocationEntity(request *UpdateLocationRequest) location.LocationEntity {
	return location.LocationEntity{
		LocationName: request.LocationName,
		UpdatedBy:      request.UpdatedBy,
	}
}
