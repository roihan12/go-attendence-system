package data

import (
	"time"

	attendanceReport "github.com/roihan12/technical-test-kalbe/features/report"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) attendanceReport.AttendanceReportData {
	return &query{
		db: db,
	}
}

func (q *query) GenerateReport(startTime, endTime time.Time) ([]attendanceReport.AttendanceReport, error) {
	var reports []attendanceReport.AttendanceReport

	// Join antara tabel attendance, employee, department, position, dan location
	err := q.db.Table("attendances").
		Select("attendances.created_at, employees.employee_code, employees.employee_name, departments.department_name, positions.position_name, locations.location_name, attendances.absent_in, attendances.absent_out").
		Joins("LEFT JOIN employees ON attendances.employee_id = employees.id").
		Joins("LEFT JOIN departments ON employees.department_id = departments.id").
		Joins("LEFT JOIN positions ON employees.position_id = positions.id").
		Joins("LEFT JOIN locations ON attendances.location_id = locations.id").
		Where("attendances.created_at BETWEEN ? AND ?", startTime, endTime).
		Scan(&reports).Error

	if err != nil {
		return nil, err
	}

	return reports, nil
}
