package services

import (
	"log"

	"github.com/roihan12/technical-test-kalbe/features/employee"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type employeeUseCase struct {
	query employee.EmployeeData
}

func New(query employee.EmployeeData) employee.EmployeeService {
	return &employeeUseCase{
		query: query,
	}
}

func (uuc *employeeUseCase) Register(newemployee employee.EmployeeEntity) (employee.EmployeeEntity, error) {
	// Generate hashed password
	hashed, err := utils.GeneratePassword(newemployee.Password)
	if err != nil {
		log.Println("bcrypt error ", err.Error())
		return employee.EmployeeEntity{}, utils.ErrInternal
	}
	newemployee.Password = string(hashed)

	// Get the latest employee code
	latestEmployee, err := uuc.query.GetLatestEmployee()
	if err != nil && err != utils.ErrDataNotFound {
		log.Println("get latest employee error", err.Error())
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	var latestCode string
	if latestEmployee != (employee.EmployeeEntity{}) {
		latestCode = latestEmployee.EmployeeCode
	}

	// Generate new employee code
	newEmployeeCode := utils.GenerateEmployeeCode(latestCode)
	newemployee.EmployeeCode = newEmployeeCode

	// Register new employee
	res, err := uuc.query.Register(newemployee)
	if err != nil {
		if err == utils.ErrConflictingData {
			return employee.EmployeeEntity{}, utils.ErrDuplicateData
		}
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	return res, nil
}


func (uuc *employeeUseCase) Login(name, password string) (string, error) {
	res, err := uuc.query.Login(name)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return "", utils.ErrInvalidCredentials
		}
		log.Println("error login query: ", err.Error())
		return "", utils.ErrInternal
	}

	if err := utils.CheckPassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", utils.ErrInvalidCredentials
	}

	useToken, err := utils.GenerateToken(res.EmployeeName, res.EmployeeID)
	if err != nil {
		return "", utils.ErrTokenCreation
	}
	return useToken, nil
}

func (uuc *employeeUseCase) Profile(employeeID uint) (employee.EmployeeEntity, error) {

	res, err := uuc.query.Profile(employeeID)
	if err != nil {
		log.Println("data not found")
		if err == utils.ErrDataNotFound {
			return employee.EmployeeEntity{}, err
		}
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (uuc *employeeUseCase) Update(employeeID uint, updateData employee.EmployeeEntity) (employee.EmployeeEntity, error) {
	if updateData.Password != "" {
		hashed, err := utils.GeneratePassword(updateData.Password)
		if err != nil {
			log.Println("bcrypt error ", err.Error())
			return employee.EmployeeEntity{}, utils.ErrInternal
		}

		updateData.Password = string(hashed)
	}

	res, err := uuc.query.Update(employeeID, updateData)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return employee.EmployeeEntity{}, err
		}
		return employee.EmployeeEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (uuc *employeeUseCase) Delete(employeeID uint) error {
	err := uuc.query.Delete(employeeID)
	if err != nil {
		if err == utils.ErrDataNotFound {
			return err
		}
		return utils.ErrInternal
	}

	return nil
}