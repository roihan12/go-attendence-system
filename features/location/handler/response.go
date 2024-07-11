package handler

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/location"
)

type LocationResponse struct {
	LocationID   uint             `json:"location_id"`
	LocationName string           `json:"location_name"`
	CreatedBy      string           `json:"created_by"`
	UpdatedBy      string           `json:"updated_by"`
	CreatedAt      time.Time        `json:"created_at,omitempty"`
	UpdatedAt      time.Time        `json:"updated_at,omitempty"`
	DeletedAt      time.Time       `json:"deleted_at,omitempty"`
}

func LocationEntityToLocationResponse(data location.LocationEntity) LocationResponse {
	LocationResponse := LocationResponse{
		LocationID:   data.LocationID,
		LocationName: data.LocationName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt,
	}

	return LocationResponse
}

func ListLocationEntityToLocationResponse(data []location.LocationEntity) []LocationResponse {
	var dataResponses []LocationResponse
	for _, v := range data {
		dataResponses = append(dataResponses, LocationEntityToLocationResponse(v))
	}
	return dataResponses
}
