package service

import (
	"errors"
	"log"

	"github.com/roihan12/technical-test-kalbe/features/department"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type departmentService struct {
	query department.DepartmentData
}

func New(q department.DepartmentData) department.DepartmentService {
	return &departmentService{
		query: q,
	}
}

func (ds *departmentService) Create(newDepartment department.DepartmentEntity) (department.DepartmentEntity, error) {
	resDepartment, err := ds.query.Create(newDepartment)
	if err != nil {
		log.Println("create department service error:", err)
		return department.DepartmentEntity{}, utils.ErrInternal
	}

	return resDepartment, nil
}

func (ds *departmentService) GetAll() ([]department.DepartmentEntity, error) {
	res, err := ds.query.GetAll()
	if err != nil {
		log.Println("get all departments service error:", err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (ds *departmentService) GetById(departmentID uint) (department.DepartmentEntity, error) {
	res, err := ds.query.GetById(departmentID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return department.DepartmentEntity{}, err
		}
		log.Println("get department by id service error:", err)
		return department.DepartmentEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (ds *departmentService) Update(updateDepartment department.DepartmentEntity, departmentID uint) (department.DepartmentEntity, error) {
	_, err := ds.query.GetById(departmentID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return department.DepartmentEntity{}, err
		}
		log.Println("get department by id service error:", err)
		return department.DepartmentEntity{}, utils.ErrInternal
	}

	updateRes, err := ds.query.Update(updateDepartment, departmentID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return department.DepartmentEntity{}, err
		}
		log.Println("update department service error:", err)
		return department.DepartmentEntity{}, utils.ErrInternal
	}

	return updateRes, nil
}

func (ds *departmentService) Delete(departmentID uint) error {
	_, err := ds.query.GetById(departmentID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("get department by id service error:", err)
		return utils.ErrInternal
	}

	err = ds.query.Delete(departmentID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("delete department service error:", err)
		return utils.ErrInternal
	}

	return nil
}
