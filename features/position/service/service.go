package service

import (
	"errors"
	"log"

	"github.com/roihan12/technical-test-kalbe/features/position"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type positionService struct {
	query position.PositionData
}

func New(q position.PositionData) position.PositionService {
	return &positionService{
		query: q,
	}
}

func (ds *positionService) Create(newposition position.PositionEntity) (position.PositionEntity, error) {
	resposition, err := ds.query.Create(newposition)
	if err != nil {
		log.Println("create position service error:", err)
		return position.PositionEntity{}, utils.ErrInternal
	}

	return resposition, nil
}

func (ds *positionService) GetAll() ([]position.PositionEntity, error) {
	res, err := ds.query.GetAll()
	if err != nil {
		log.Println("get all positions service error:", err)
		return nil, utils.ErrInternal
	}

	return res, nil
}

func (ds *positionService) GetById(positionID uint) (position.PositionEntity, error) {
	res, err := ds.query.GetById(positionID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return position.PositionEntity{}, err
		}
		log.Println("get position by id service error:", err)
		return position.PositionEntity{}, utils.ErrInternal
	}

	return res, nil
}

func (ds *positionService) Update(updateposition position.PositionEntity, positionID uint) (position.PositionEntity, error) {
	_, err := ds.query.GetById(positionID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return position.PositionEntity{}, err
		}
		log.Println("get position by id service error:", err)
		return position.PositionEntity{}, utils.ErrInternal
	}

	updateRes, err := ds.query.Update(updateposition, positionID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return position.PositionEntity{}, err
		}
		log.Println("update position service error:", err)
		return position.PositionEntity{}, utils.ErrInternal
	}

	return updateRes, nil
}

func (ds *positionService) Delete(positionID uint) error {
	_, err := ds.query.GetById(positionID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("get position by id service error:", err)
		return utils.ErrInternal
	}

	err = ds.query.Delete(positionID)
	if err != nil {
		if errors.Is(err, utils.ErrDataNotFound) {
			return err
		}
		log.Println("delete position service error:", err)
		return utils.ErrInternal
	}

	return nil
}
