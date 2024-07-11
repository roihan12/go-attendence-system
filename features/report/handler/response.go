package handler

import (
	"time"

	"github.com/roihan12/technical-test-kalbe/features/report"
)

type AttendanceReportResponse struct {
	Date          time.Time `json:"date"`
	EmployeeCode  string    `json:"employee_code"`
	EmployeeName  string    `json:"employee_name"`
	DepartmentName string    `json:"department_name"`
	PositionName  string    `json:"position_name"`
	LocationName  string    `json:"location_name"`
	AbsentIn      time.Time `json:"absent_in"`
	AbsentOut     time.Time `json:"absent_out"`
}

func AttendanceEntityToResponse(data report.AttendanceReport) AttendanceReportResponse {
	return AttendanceReportResponse{
		Date:          data.Date,
		EmployeeCode:  data.EmployeeCode,
		EmployeeName:  data.EmployeeName,
		DepartmentName: data.DepartmentName,
		PositionName:  data.PositionName,
		LocationName:  data.LocationName,
		AbsentIn:      data.AbsentIn,
		AbsentOut:     data.AbsentOut,
	}
}

func ListAttendanceEntityToResponse(data []report.AttendanceReport) []AttendanceReportResponse {
	var responses []AttendanceReportResponse
	for _, v := range data {
		responses = append(responses, AttendanceEntityToResponse(v))
	}
	return responses
}
