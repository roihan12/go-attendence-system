package handler

import (
	"time"

	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
)

type CreateAttendanceRequest struct {
	EmployeeID uint      `json:"employee_id" form:"employee_id" binding:"required" example:"1"`
	LocationID uint      `json:"location_id" form:"location_id" binding:"required" example:"1"`
	AbsentIn   time.Time `json:"absent_in" form:"absent_in" binding:"required" example:"2024-07-11T08:00:00Z"`
	AbsentOut  time.Time `json:"absent_out" form:"absent_out" binding:"required" example:"2024-07-11T17:00:00Z"`
	CreatedBy  string    `json:"created_by" form:"created_by" binding:"required" example:"Admin"`
}

type UpdateAttendanceRequest struct {
	EmployeeID uint      `json:"employee_id" form:"employee_id" binding:"required" example:"1"`
	LocationID uint      `json:"location_id" form:"location_id" binding:"required" example:"1"`
	AbsentIn   time.Time `json:"absent_in" form:"absent_in" binding:"required" example:"2024-07-11T08:00:00Z"`
	AbsentOut  time.Time `json:"absent_out" form:"absent_out" binding:"required" example:"2024-07-11T17:00:00Z"`
	UpdatedBy  string    `json:"updated_by" form:"updated_by" binding:"required" example:"Admin"`
}

func CreateAttendanceRequestToAttendanceEntity(request *CreateAttendanceRequest) attendance.AttendanceEntity {
	return attendance.AttendanceEntity{
		EmployeeID: request.EmployeeID,
		LocationID: request.LocationID,
		AbsentIn:   request.AbsentIn,
		AbsentOut:  request.AbsentOut,
		CreatedBy:  request.CreatedBy,
	}
}

func UpdateAttendanceRequestToAttendanceEntity(request *UpdateAttendanceRequest) attendance.AttendanceEntity {
	return attendance.AttendanceEntity{
		EmployeeID: request.EmployeeID,
		LocationID: request.LocationID,
		AbsentIn:   request.AbsentIn,
		AbsentOut:  request.AbsentOut,
		UpdatedBy:  request.UpdatedBy,
	}
}
