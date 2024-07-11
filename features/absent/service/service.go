package service

import (
	"errors"
	"log"

	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type attendanceService struct {
	query attendance.AttendanceData
}

func New(q attendance.AttendanceData) attendance.AttendanceService {
	return &attendanceService{
		query: q,
	}
}

func (as *attendanceService) Create(newAttendance attendance.AttendanceEntity) (attendance.AttendanceEntity, error) {
	resAttendance, err := as.query.Create(newAttendance)
	if err != nil {
		log.Println("create attendance service error:", err)
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}

	return resAttendance, nil
}

func (as *attendanceService) GetAll() ([]attendance.AttendanceEntity, error) {
	res, err := as.query.GetAll()
	if err != nil {
		log.Println("get all attendances service error:", err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (as *attendanceService) GetById(attendanceID uint) (attendance.AttendanceEntity, error) {
	res, err := as.query.GetById(attendanceID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return attendance.AttendanceEntity{}, err
		}
		log.Println("get attendance by id service error:", err)
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (as *attendanceService) Update(updateAttendance attendance.AttendanceEntity, attendanceID uint) (attendance.AttendanceEntity, error) {
	_, err := as.query.GetById(attendanceID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return attendance.AttendanceEntity{}, err
		}
		log.Println("get attendance by id service error:", err)
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}

	updateRes, err := as.query.Update(updateAttendance, attendanceID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return attendance.AttendanceEntity{}, err
		}
		log.Println("update attendance service error:", err)
		return attendance.AttendanceEntity{}, utils.ErrInternal
	}

	return updateRes, nil
}

func (as *attendanceService) Delete(attendanceID uint) error {
	_, err := as.query.GetById(attendanceID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("get attendance by id service error:", err)
		return utils.ErrInternal
	}

	err = as.query.Delete(attendanceID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("delete attendance service error:", err)
		return utils.ErrInternal
	}

	return nil
}
