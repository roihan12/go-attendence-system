package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/features/position"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type PositionController struct {
	PositionService position.PositionService
}

func New(srv position.PositionService) *PositionController {
	return &PositionController{
		PositionService: srv,
	}
}

// GetAll godoc
//
// @Summary		List Positions
// @Description	List Positions
// @Tags			Positions
// @Accept			json
// @Produce		json
// @Success		200	{array}		PositionResponse	"Positions retrieved"
// @Failure		500	{object}	utils.ErrorResponse	"Internal server error"
// @Router			/positions [get]
// @Security		BearerAuth
func (dc *PositionController) GetAll(c *fiber.Ctx) error {

	Positions, err := dc.PositionService.GetAll()
	if err != nil {
		return utils.HandleError(c, err)
	}

	listPositionResponse := ListPositionEntityToPositionResponse(Positions)
	return utils.HandleSuccess(c, "Get all Positions successfully", listPositionResponse)
}

// GetById godoc
//
// @Summary		Get a Position
// @Description	get a Position by id
// @Tags			Positions
// @Accept			json
// @Produce		json
// @Param			positionId	path		uint				true	"Position ID"
// @Success		200		{object}	PositionResponse	"Position retrieved"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/positions/{positionId} [get]
// @Security		BearerAuth
func (dc *PositionController) GetById(c *fiber.Ctx) error {
	Id := c.Params("positionId")
	PositionId, _ := strconv.Atoi(Id)
	Position, err := dc.PositionService.GetById(uint(PositionId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	response := PositionEntityToPositionResponse(Position)
	return utils.HandleSuccess(c, "Get by id Position successfully", response)
}

// CreatePosition godoc
//
// @Summary		Create a new Position
// @Description	create a new Position with title and description
// @Tags			Positions
// @Produce		json
// @Param			CreatePositionRequest	formData	CreatePositionRequest	true	"Create Position request"
// @Success		200					{object}	PositionResponse		"Position created"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/positions [post]
// @Security		BearerAuth
func (dc *PositionController) Create(c *fiber.Ctx) error {
	var req CreatePositionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	PositionEntity := CreatePositionRequestToPositionEntity(&req)

	response, err := dc.PositionService.Create(PositionEntity)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Create Position successfully", PositionEntityToPositionResponse(response))
}

// UpdatePosition godoc
//
// @Summary		Update a Position
// @Description	update a Position with title and description
// @Tags			Positions
// @Produce		json
// @Param			positionId			path		uint				true	"Position ID"
// @Param			UpdatePositionRequest	formData	UpdatePositionRequest	true	"Update Position request"
// @Success		200					{object}	PositionResponse		"Position updated"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/positions/{positionId} [put]
// @Security		BearerAuth
func (dc *PositionController) Update(c *fiber.Ctx) error {

	Id := c.Params("positionId")
	PositionId, _ := strconv.Atoi(Id)
	var req UpdatePositionRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	PositionEntity := UpdatePositionRequestToPositionEntity(&req)

	response, err := dc.PositionService.Update(PositionEntity, uint(PositionId))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Update Position successfully", PositionEntityToPositionResponse(response))
}

// DeletePosition godoc
//
// @Summary		Delete a Position
// @Description	delete a Position by id
// @Tags			Positions
// @Accept			json
// @Produce		json
// @Param			positionId	path		uint				true	"Position ID"
// @Success		200		{object}	utils.Response		"Position deleted"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/positions/{positionId} [delete]
// @Security		BearerAuth
func (dc *PositionController) Delete(c *fiber.Ctx) error {
	Id := c.Params("positionId")
	PositionId, _ := strconv.Atoi(Id)

	err := dc.PositionService.Delete(uint(PositionId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Your Position has been successfully deleted", nil)
}
