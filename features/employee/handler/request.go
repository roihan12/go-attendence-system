package handler

import (
	"github.com/roihan12/technical-test-kalbe/features/employee"
)

// LoginRequest represents the request body for logging in an employee
type LoginRequest struct {
	EmployeeName string `json:"employee_name" form:"employee_name" binding:"required" example:"John Doe"`
	Password string `json:"password" form:"password" binding:"required,min=6" example:"123456" minLength:"6"`
}

// RegisterRequest represents the request body for registering an employee
type RegisterRequest struct {
	EmployeeName string `json:"employee_name" form:"employee_name" binding:"required" example:"John Doe"`
	Password     string `json:"password" form:"password" binding:"required,min=6" example:"123456" minLength:"6"`
	DepartmentID uint   `json:"department_id" form:"department_id" binding:"required" example:"1"`
	PositionID   uint   `json:"position_id" form:"position_id" binding:"required" example:"1"`
	Superior     uint   `json:"superior" form:"superior" binding:"required" example:"1"`
	CreatedBy    string `json:"created_by" form:"created_by" binding:"required" example:"admin"`
}

// UpdateRequest represents the request body for updating an employee
type UpdateRequest struct {
	EmployeeName string `json:"employee_name" form:"employee_name" binding:"omitempty,required" example:"John Doe"`
	Password     string `json:"password" form:"password" binding:"omitempty,required,min=6" example:"123456"`
	DepartmentID uint   `json:"department_id" form:"department_id" binding:"omitempty,required" example:"1"`
	PositionID   uint   `json:"position_id" form:"position_id" binding:"omitempty,required" example:"1"`
	Superior     uint   `json:"superior" form:"superior" binding:"omitempty,required" example:"1"`
	UpdatedBy    string `json:"updated_by" form:"updated_by" binding:"required" example:"admin"`
}

func ReqToCore(data interface{}) *employee.EmployeeEntity {
	res := employee.EmployeeEntity{}

	switch docs := data.(type) {
	case LoginRequest:
		cnv := docs
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := data.(RegisterRequest)
		res.EmployeeName = cnv.EmployeeName
		res.Password = cnv.Password
		res.DepartmentID = cnv.DepartmentID
		res.PositionID = cnv.PositionID
		res.Superior = cnv.Superior
		res.CreatedBy = cnv.CreatedBy
	case UpdateRequest:
		cnv := data.(UpdateRequest)
		res.EmployeeName = cnv.EmployeeName
		res.Password = cnv.Password
		res.DepartmentID = cnv.DepartmentID
		res.PositionID = cnv.PositionID
		res.Superior = cnv.Superior
		res.UpdatedBy = cnv.UpdatedBy
	default:
		return nil
	}

	return &res
}
