package report

import (
	"time"
)

type AttendanceReport struct {
	Date          time.Time `json:"date"`
	EmployeeCode  string    `json:"employee_code"`
	EmployeeName  string    `json:"employee_name"`
	DepartmentName string   `json:"department_name"`
	PositionName  string    `json:"position_name"`
	LocationName  string    `json:"location_name"`
	AbsentIn      time.Time `json:"absent_in"`
	AbsentOut     time.Time `json:"absent_out"`
}

type AttendanceReportService interface {
	GenerateReport(startDate, endDate time.Time) ([]AttendanceReport, error)
}

type AttendanceReportData interface {
	GenerateReport(startDate, endDate time.Time) ([]AttendanceReport, error)
}
