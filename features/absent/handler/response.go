package handler

import (
	"time"

	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
)

type AttendanceResponse struct {
	AttendanceID uint      `json:"attendance_id"`
	EmployeeID   uint      `json:"employee_id"`
	LocationID   uint      `json:"location_id"`
	AbsentIn     time.Time `json:"absent_in"`
	AbsentOut    time.Time `json:"absent_out"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	DeletedAt    time.Time `json:"deleted_at,omitempty"`
}

func AttendanceEntityToAttendanceResponse(data attendance.AttendanceEntity) AttendanceResponse {
	return AttendanceResponse{
		AttendanceID: data.AttendanceID,
		EmployeeID:   data.EmployeeID,
		LocationID:   data.LocationID,
		AbsentIn:     data.AbsentIn,
		AbsentOut:    data.AbsentOut,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
		DeletedAt:    data.DeletedAt,
	}
}

func ListAttendanceEntityToAttendanceResponse(data []attendance.AttendanceEntity) []AttendanceResponse {
	var dataResponses []AttendanceResponse
	for _, v := range data {
		dataResponses = append(dataResponses, AttendanceEntityToAttendanceResponse(v))
	}
	return dataResponses
}
