package data

import (
	"log"
	"time"

	"github.com/roihan12/technical-test-kalbe/features/employee"
	"github.com/roihan12/technical-test-kalbe/utils"
	"gorm.io/gorm"
)

type employeeQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) employee.EmployeeData {
	return &employeeQuery{
		db: db,
	}
}

func (eq *employeeQuery) CheckName(newEmployee employee.EmployeeEntity) error {
	e := Employee{}
	eq.db.Where("employee_name = ?", newEmployee.EmployeeName).First(&e)
	if e.ID != 0 {
		if e.EmployeeName == newEmployee.EmployeeName {
			return utils.ErrConflictingData
		}
	}
	return nil
}

func (eq *employeeQuery) GetLatestEmployee() (employee.EmployeeEntity, error) {
	latestEmployee := Employee{}
	err := eq.db.Order("employee_code desc").First(&latestEmployee).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return employee.EmployeeEntity{}, utils.ErrDataNotFound
		}
		return employee.EmployeeEntity{}, err
	}
	return ToEmployeeEntity(latestEmployee), nil
}

func (eq *employeeQuery) Login(email string) (employee.EmployeeEntity, error) {
	res := Employee{}
	if err := eq.db.Where("employee_name = ? AND deleted_at IS NULL", email).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return employee.EmployeeEntity{}, utils.ErrDataNotFound
	}

	return ToEmployeeEntity(res), nil
}

func (eq *employeeQuery) Register(newEmployee employee.EmployeeEntity) (employee.EmployeeEntity, error) {
	// Check Employee
	if err := eq.CheckName(newEmployee); err != nil {
		log.Println("error create new employee: ", err.Error())
		return employee.EmployeeEntity{}, utils.ErrConflictingData
	}

	cnv := EmployeeEntityToEmployee(newEmployee)
	err := eq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	newEmployee.EmployeeID = cnv.ID
	return newEmployee, nil
}

func (eq *employeeQuery) Profile(employeeID uint) (employee.EmployeeEntity, error) {
	res := Employee{}
	if err := eq.db.Where("id = ? AND deleted_at IS NULL", employeeID).First(&res).Error; err != nil {
		log.Println("get profile query error", err.Error())
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	return ToEmployeeEntity(res), nil
}

func (eq *employeeQuery) Update(employeeID uint, updateData employee.EmployeeEntity) (employee.EmployeeEntity, error) {
	cnv := EmployeeEntityToEmployee(updateData)
	res := Employee{}
	qry := eq.db.Model(&res).Where("id = ? AND deleted_at IS NULL", employeeID).Updates(&cnv)

	if qry.RowsAffected <= 0 {
		log.Println("\tupdate employee query error: data not found")
		return employee.EmployeeEntity{}, utils.ErrDataNotFound
	}

	if err := qry.Error; err != nil {
		log.Println("\tupdate employee query error: ", err.Error())
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	return ToEmployeeEntity(cnv), nil
}

func (eq *employeeQuery) Delete(employeeID uint) error {
	res := Employee{}

	// Soft delete: update the DeletedAt field
	qry := eq.db.Model(&res).Where("id = ?", employeeID).Update("deleted_at", time.Now())

	if qry.RowsAffected <= 0 {
		log.Println("\tDelete employee query error: data not found")
		return utils.ErrDataNotFound
	}
	if err := qry.Error; err != nil {
		log.Println("\tDelete employee query error: ", err.Error())
		return utils.ErrInternal
	}

	return nil
}
