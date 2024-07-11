package data

import (
	"time"

	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
	employee "github.com/roihan12/technical-test-kalbe/features/employee/data"
	location "github.com/roihan12/technical-test-kalbe/features/location/data"
	"gorm.io/gorm"
)

type Attendance struct {
	ID         uint              `gorm:"primaryKey;autoIncrement"`
	EmployeeID uint              `gorm:"type:integer;not null"`
	LocationID uint              `gorm:"type:integer;not null"`
	AbsentIn   time.Time        
	AbsentOut  time.Time         
	CreatedAt  time.Time         `gorm:"autoCreateTime"`
	CreatedBy  string            `gorm:"type:varchar(255)"`
	UpdatedAt  time.Time         `gorm:"autoUpdateTime"`
	UpdatedBy  string            `gorm:"type:varchar(255)"`
	DeletedAt  gorm.DeletedAt    `gorm:"index"`
	Employee   employee.Employee `gorm:"foreignKey:EmployeeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"` 
	Location   location.Location `gorm:"foreignKey:LocationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Convert AttendanceEntity to Attendance
func AttendanceEntityToAttendance(data attendance.AttendanceEntity) Attendance {
	return Attendance{
		ID:         data.AttendanceID,
		EmployeeID: data.EmployeeID,
		LocationID: data.LocationID,
		AbsentIn:   data.AbsentIn,
		AbsentOut:  data.AbsentOut,
		CreatedAt:  data.CreatedAt,
		CreatedBy:  data.CreatedBy,
		UpdatedAt:  data.UpdatedAt,
		UpdatedBy:  data.UpdatedBy,
		DeletedAt:  gorm.DeletedAt{Time: data.DeletedAt, Valid: !data.DeletedAt.IsZero()},
	}
}

// Convert Attendance to AttendanceEntity
func AttendanceToAttendanceEntity(data Attendance) attendance.AttendanceEntity {
	return attendance.AttendanceEntity{
		AttendanceID: data.ID,
		EmployeeID:   data.EmployeeID,
		LocationID:   data.LocationID,
		AbsentIn:     data.AbsentIn,
		AbsentOut:    data.AbsentOut,
		CreatedAt:    data.CreatedAt,
		CreatedBy:    data.CreatedBy,
		UpdatedAt:    data.UpdatedAt,
		UpdatedBy:    data.UpdatedBy,
		DeletedAt:    data.DeletedAt.Time,
	}
}

// Convert list of Attendance to list of AttendanceEntity
func ListAttendanceToAttendanceEntity(data []Attendance) []attendance.AttendanceEntity {
	var attendanceEntities []attendance.AttendanceEntity
	for _, v := range data {
		attendanceEntities = append(attendanceEntities, AttendanceToAttendanceEntity(v))
	}
	return attendanceEntities
}
