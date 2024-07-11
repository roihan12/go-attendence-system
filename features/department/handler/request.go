package handler

import (
	"github.com/roihan12/technical-test-kalbe/features/department"
)

type CreateDepartmentRequest struct {
	DepartmentName string `json:"department_name" form:"department_name" binding:"required" example:"Finance"`
	CreatedBy      string `json:"created_by" form:"created_by" binding:"required" example:"Admin"`
}

type UpdateDepartmentRequest struct {
	DepartmentName string `json:"department_name" form:"department_name" binding:"required" example:"Finance"`
	UpdatedBy      string `json:"updated_by" form:"updated_by" binding:"required" example:"Admin"`
}

func CreateDepartmentRequestToDepartmentEntity(request *CreateDepartmentRequest) department.DepartmentEntity {
	return department.DepartmentEntity{
		DepartmentName: request.DepartmentName,
		CreatedBy:      request.CreatedBy,
	}
}

func UpdateDepartmentRequestToDepartmentEntity(request *UpdateDepartmentRequest) department.DepartmentEntity {
	return department.DepartmentEntity{
		DepartmentName: request.DepartmentName,
		UpdatedBy:      request.UpdatedBy,
	}
}
