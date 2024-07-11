package data

import (
	"errors"
	"log"
	"time"

	"github.com/roihan12/technical-test-kalbe/features/location"
	"github.com/roihan12/technical-test-kalbe/utils"
	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) location.LocationData {
	return &query{
		db: db,
	}
}

func (q *query) GetAll() ([]location.LocationEntity, error) {
	var locations []Location

	if err := q.db.Where("deleted_at IS NULL").Find(&locations).Error; err != nil {
		return nil, err
	}

	return ListLocationToLocationEntity(locations), nil
}

func (q *query) Create(LocationEntity location.LocationEntity) (location.LocationEntity, error) {
	newLocation := LocationEntityToLocation(LocationEntity)
	if err := q.db.Create(&newLocation).Error; err != nil {
		log.Println("create Location query error", err.Error())
		return location.LocationEntity{}, utils.ErrInternal
	}
	return LocationToLocationEntity(newLocation), nil
}

func (q *query) Update(LocationEntity location.LocationEntity, id uint) (location.LocationEntity, error) {
	_, err := q.GetById(id)
	if err != nil {
		return location.LocationEntity{}, err
	}

	updateLocation := LocationEntityToLocation(LocationEntity)

	if err := q.db.Model(&Location{}).Where("id = ? AND deleted_at IS NULL", id).Updates(&updateLocation).Error; err != nil {
		log.Println("update Location query error", err.Error())
		return location.LocationEntity{}, utils.ErrInternal
	}
	return LocationToLocationEntity(updateLocation), nil
}

func (q *query) GetById(id uint) (location.LocationEntity, error) {
	var singleLocation Location
	if err := q.db.Where("id = ? AND deleted_at IS NULL", id).First(&singleLocation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return location.LocationEntity{}, utils.ErrDataNotFound
		}
		return location.LocationEntity{}, err
	}

	return LocationToLocationEntity(singleLocation), nil
}

func (q *query) Delete(id uint) error {
	if err := q.db.Model(&Location{}).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println("soft delete Location query error", err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.ErrDataNotFound
		}
		return utils.ErrInternal
	}

	return nil
}

