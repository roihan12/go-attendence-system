package data

import (
	"errors"
	"log"
	"time"

	"github.com/roihan12/technical-test-kalbe/features/department"
	"github.com/roihan12/technical-test-kalbe/utils"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) department.DepartmentData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]department.DepartmentEntity, error) {
	var departments []Department

	if err := q.db.Where("deleted_at IS NULL").Find(&departments).Error; err != nil {
		return nil, err
	}

	return ListDepartmentToDepartmentEntity(departments), nil
}

func (q *query) Create(departmentEntity department.DepartmentEntity) (department.DepartmentEntity, error) {
	newDepartment := DepartmentEntityToDepartment(departmentEntity)
	if err := q.db.Create(&newDepartment).Error; err != nil {
		log.Println("create department query error", err.Error())
		return department.DepartmentEntity{}, utils.ErrInternal
	}
	return DepartmentToDepartmentEntity(newDepartment), nil
}

func (q *query) Update(departmentEntity department.DepartmentEntity, id uint) (department.DepartmentEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return department.DepartmentEntity{}, err
	}

	updateDepartment := DepartmentEntityToDepartment(departmentEntity)

	if err := q.db.Model(&Department{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updateDepartment).Error; err != nil {
		log.Println("update department query error", err.Error())
		return department.DepartmentEntity{}, utils.ErrInternal
	}
	return DepartmentToDepartmentEntity(updateDepartment), nil
}

func (q *query) GetById(id uint) (department.DepartmentEntity, error) {
	var singleDepartment Department
	if err := q.db.Where("id = ? AND deleted_at IS NULL", id).First(&singleDepartment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return department.DepartmentEntity{}, utils.ErrDataNotFound
		}
		return department.DepartmentEntity{}, err
	}

	return DepartmentToDepartmentEntity(singleDepartment), nil
}

func (q *query) Delete(id uint) error {
	if err := q.db.Model(&Department{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("soft delete department query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}

