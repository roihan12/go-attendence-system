package data

import (
	"errors"
	"log"
	"time"

	"github.com/roihan12/technical-test-kalbe/features/position"
	"github.com/roihan12/technical-test-kalbe/utils"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) position.PositionData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]position.PositionEntity, error) {
	var positions []Position

	if err := q.db.Where("deleted_at IS NULL").Find(&positions).Error; err != nil {
		return nil, err
	}

	return ListPositionToPositionEntity(positions), nil
}



func (q *query) Create(positionEntity position.PositionEntity) (position.PositionEntity, error) {
	newPosition := PositionEntityToPosition(positionEntity)
	if err := q.db.Create(&newPosition).Error; err != nil {
		log.Println("create Position query error", err.Error())
		return position.PositionEntity{}, utils.ErrInternal
	}
	return PositionToPositionEntity(newPosition), nil
}

func (q *query) Update(positionEntity position.PositionEntity, id uint) (position.PositionEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return position.PositionEntity{}, err
	}

	updatePosition := PositionEntityToPosition(positionEntity)

	if err := q.db.Model(&Position{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updatePosition).Error; err != nil {
		log.Println("update Position query error", err.Error())
		return position.PositionEntity{}, utils.ErrInternal
	}
	return PositionToPositionEntity(updatePosition), nil
}

func (q *query) GetById(id uint) (position.PositionEntity, error) {
	var singlePosition Position
	if err := q.db.Where("id = ? AND deleted_at IS NULL", id).First(&singlePosition).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return position.PositionEntity{}, utils.ErrDataNotFound
		}
		return position.PositionEntity{}, err
	}

	return PositionToPositionEntity(singlePosition), nil
}

func (q *query) Delete(id uint) error {
	if err := q.db.Model(&Position{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("soft delete Position query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}

