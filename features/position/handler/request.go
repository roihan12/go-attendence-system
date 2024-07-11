package handler
import (
	"github.com/roihan12/technical-test-kalbe/features/position"
)
type CreatePositionRequest struct {
	PositionName string `json:"position_name" form:"position_name" binding:"required" example:"Finance"`
	CreatedBy    string `json:"created_by" form:"created_by" binding:"required" example:"Admin"`
	DepartmentID uint `json:"department_id" form:"department_id" binding:"required" example:"1"`
}

type UpdatePositionRequest struct {
	PositionName string `json:"position_name" form:"position_name" binding:"required" example:"Finance"`
	UpdatedBy    string `json:"updated_by" form:"updated_by" binding:"required" example:"Admin"`
	DepartmentID uint `json:"department_id" form:"department_id" binding:"required" example:"1"`
}

func CreatePositionRequestToPositionEntity(request *CreatePositionRequest) position.PositionEntity {
	return position.PositionEntity{
		PositionName: request.PositionName,
		CreatedBy:    request.CreatedBy,
		DepartmentID: request.DepartmentID,
	}
}

func UpdatePositionRequestToPositionEntity(request *UpdatePositionRequest) position.PositionEntity {
	return position.PositionEntity{
		PositionName: request.PositionName,
		UpdatedBy:    request.UpdatedBy,
	}
}
