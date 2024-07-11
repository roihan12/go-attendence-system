package service

import (
	"errors"
	"log"

	"github.com/roihan12/technical-test-kalbe/features/location"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type LocationService struct {
	query location.LocationData
}

func New(q location.LocationData) location.LocationService {
	return &LocationService{
		query: q,
	}
}

func (ds *LocationService) Create(newLocation location.LocationEntity) (location.LocationEntity, error) {
	resLocation, err := ds.query.Create(newLocation)
	if err != nil {
		log.Println("create Location service error:", err)
		return location.LocationEntity{}, utils.ErrInternal
	}

	return resLocation, nil
}

func (ds *LocationService) GetAll() ([]location.LocationEntity, error) {
	res, err := ds.query.GetAll()
	if err != nil {
		log.Println("get all Locations service error:", err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (ds *LocationService) GetById(locationID uint) (location.LocationEntity, error) {
	res, err := ds.query.GetById(locationID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return location.LocationEntity{}, err
		}
		log.Println("get Location by id service error:", err)
		return location.LocationEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (ds *LocationService) Update(updateLocation location.LocationEntity, locationID uint) (location.LocationEntity, error) {
	_, err := ds.query.GetById(locationID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return location.LocationEntity{}, err
		}
		log.Println("get Location by id service error:", err)
		return location.LocationEntity{}, utils.ErrInternal
	}

	updateRes, err := ds.query.Update(updateLocation, locationID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return location.LocationEntity{}, err
		}
		log.Println("update Location service error:", err)
		return location.LocationEntity{}, utils.ErrInternal
	}

	return updateRes, nil
}

func (ds *LocationService) Delete(locationID uint) error {
	_, err := ds.query.GetById(locationID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("get Location by id service error:", err)
		return utils.ErrInternal
	}

	err = ds.query.Delete(locationID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("delete Location service error:", err)
		return utils.ErrInternal
	}

	return nil
}
