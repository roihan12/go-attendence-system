package handler

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/department"
)

type DepartmentResponse struct {
	DepartmentID   uint             `json:"department_id"`
	DepartmentName string           `json:"department_name"`
	CreatedBy      string           `json:"created_by"`
	UpdatedBy      string           `json:"updated_by"`
	CreatedAt      time.Time        `json:"created_at,omitempty"`
	UpdatedAt      time.Time        `json:"updated_at,omitempty"`
	DeletedAt      time.Time       `json:"deleted_at,omitempty"`
}

func DepartmentEntityToDepartmentResponse(data department.DepartmentEntity) DepartmentResponse {
	departmentResponse := DepartmentResponse{
		DepartmentID:   data.DepartmentID,
		DepartmentName: data.DepartmentName,
		CreatedBy:      data.CreatedBy,
		UpdatedBy:      data.UpdatedBy,
		CreatedAt:      data.CreatedAt,
		UpdatedAt:      data.UpdatedAt,
		DeletedAt:      data.DeletedAt,
	}

	return departmentResponse
}

func ListDepartmentEntityToDepartmentResponse(data []department.DepartmentEntity) []DepartmentResponse {
	var dataResponses []DepartmentResponse
	for _, v := range data {
		dataResponses = append(dataResponses, DepartmentEntityToDepartmentResponse(v))
	}
	return dataResponses
}
