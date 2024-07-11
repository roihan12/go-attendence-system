package absent

import (
	"time"
)

type AttendanceEntity struct {
	AttendanceID     uint
	EmployeeID  uint
	LocationID uint
	AbsentIn   time.Time
	AbsentOut  time.Time
	CreatedBy       string
	UpdatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type AttendanceService interface {
	GetAll() ([]AttendanceEntity, error)
	GetById(id uint) (AttendanceEntity, error)
	Create(createPosition AttendanceEntity,) (AttendanceEntity, error)
	Update(updatePosition AttendanceEntity, id uint) (AttendanceEntity, error)
	Delete(id uint) error
}

type AttendanceData interface {
	GetAll() ([]AttendanceEntity, error)
	GetById(id uint) (AttendanceEntity, error)
	Create(createPosition AttendanceEntity) (AttendanceEntity, error)
	Update(updatePosition AttendanceEntity, id uint) (AttendanceEntity, error)
	Delete(id uint) error
}


