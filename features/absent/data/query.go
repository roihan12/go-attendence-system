package data

import (
	"errors"
	"log"
	"time"

	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
	"github.com/roihan12/technical-test-kalbe/utils"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) attendance.AttendanceData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]attendance.AttendanceEntity, error) {
	var attendances []Attendance

	if err := q.db.Where("deleted_at IS NULL").Find(&attendances).Error; err != nil {
		return nil, err
	}

	return ListAttendanceToAttendanceEntity(attendances), nil
}

func (q *query) Create(attendanceEntity attendance.AttendanceEntity) (attendance.AttendanceEntity, error) {
	newAttendance := AttendanceEntityToAttendance(attendanceEntity)
	if err := q.db.Create(&newAttendance).Error; err != nil {
		log.Println("create Attendance query error", err.Error())
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}
	return AttendanceToAttendanceEntity(newAttendance), nil
}

func (q *query) Update(attendanceEntity attendance.AttendanceEntity, id uint) (attendance.AttendanceEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return attendance.AttendanceEntity{}, err
	}

	updateAttendance := AttendanceEntityToAttendance(attendanceEntity)

	if err := q.db.Model(&Attendance{}).Where("attendance_id = ? AND deleted_at IS NULL", id).Updates(&updateAttendance).Error; err != nil {
		log.Println("update Attendance query error", err.Error())
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}
	return AttendanceToAttendanceEntity(updateAttendance), nil
}

func (q *query) GetById(id uint) (attendance.AttendanceEntity, error) {
	var singleAttendance Attendance
	if err := q.db.Where("attendance_id = ? AND deleted_at IS NULL", id).First(&singleAttendance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return attendance.AttendanceEntity{}, utils.ErrDataNotFound
		}
		return attendance.AttendanceEntity{}, err
	}

	return AttendanceToAttendanceEntity(singleAttendance), nil
}

func (q *query) Delete(id uint) error {
	if err := q.db.Model(&Attendance{}).Where("attendance_id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("soft delete Attendance query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}
