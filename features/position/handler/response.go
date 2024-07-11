package handler

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/position"
)

type PositionResponse struct {
	PositionID   uint      `json:"position_id"`
	PositionName string    `json:"position_name"`
	DepartmentID uint      `json:"department_id"`
	CreatedBy    string    `json:"created_by"`
	UpdatedBy    string    `json:"updated_by"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

func PositionEntityToPositionResponse(data position.PositionEntity) PositionResponse {
	PositionResponse := PositionResponse{
		PositionID:   data.PositionID,
		PositionName: data.PositionName,
		DepartmentID: data.DepartmentID,
		CreatedBy:    data.CreatedBy,
		UpdatedBy:    data.UpdatedBy,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}

	return PositionResponse
}

func ListPositionEntityToPositionResponse(data []position.PositionEntity) []PositionResponse {
	var dataResponses []PositionResponse
	for _, v := range data {
		dataResponses = append(dataResponses, PositionEntityToPositionResponse(v))
	}
	return dataResponses
}
