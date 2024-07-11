package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/features/location"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type LocationController struct {
	LocationService location.LocationService
}

func New(srv location.LocationService) *LocationController {
	return &LocationController{
		LocationService: srv,
	}
}

// GetAll godoc
//
// @Summary		List Locations
// @Description	List Locations
// @Tags			Locations
// @Accept			json
// @Produce		json
// @Success		200	{array}		LocationResponse	"Locations retrieved"
// @Failure		500	{object}	utils.ErrorResponse	"Internal server error"
// @Router			/locations [get]
// @Security		BearerAuth
func (dc *LocationController) GetAll(c *fiber.Ctx) error {

	locations, err := dc.LocationService.GetAll()
	if err != nil {
		return utils.HandleError(c, err)
	}

	listLocationResponse := ListLocationEntityToLocationResponse(locations)
	return utils.HandleSuccess(c, "Get all Locations successfully", listLocationResponse)
}

// GetById godoc
//
// @Summary		Get a Location
// @Description	get a Location by id
// @Tags			Locations
// @Accept			json
// @Produce		json
// @Param			locationId	path		uint				true	"Location ID"
// @Success		200		{object}	LocationResponse	"Location retrieved"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/locations/{locationId} [get]
// @Security		BearerAuth
func (dc *LocationController) GetById(c *fiber.Ctx) error {
	Id := c.Params("locationId")
	LocationId, _ := strconv.Atoi(Id)
	Location, err := dc.LocationService.GetById(uint(LocationId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	response := LocationEntityToLocationResponse(Location)
	return utils.HandleSuccess(c, "Get by id Location successfully", response)
}

// CreateLocation godoc
//
// @Summary		Create a new Location
// @Description	create a new Location with title and description
// @Tags			Locations
// @Produce		json
// @Param			CreateLocationRequest	formData	CreateLocationRequest	true	"Create Location request"
// @Success		200					{object}	LocationResponse		"Location created"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/locations [post]
// @Security		BearerAuth
func (dc *LocationController) Create(c *fiber.Ctx) error {
	var req CreateLocationRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	locationEntity := CreateLocationRequestToLocationEntity(&req)

	response, err := dc.LocationService.Create(locationEntity)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Create Location successfully", LocationEntityToLocationResponse(response))
}

// UpdateLocation godoc
//
// @Summary		Update a Location
// @Description	update a Location with title and description
// @Tags			Locations
// @Produce		json
// @Param			locationId			path		uint				true	"Location ID"
// @Param			UpdateLocationRequest	formData	UpdateLocationRequest	true	"Update Location request"
// @Success		200					{object}	LocationResponse		"Location updated"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/locations/{locationId} [put]
// @Security		BearerAuth
func (dc *LocationController) Update(c *fiber.Ctx) error {

	Id := c.Params("locationId")
	LocationId, _ := strconv.Atoi(Id)
	var req UpdateLocationRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	locationEntity := UpdateLocationRequestToLocationEntity(&req)

	response, err := dc.LocationService.Update(locationEntity, uint(LocationId))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Update Location successfully", LocationEntityToLocationResponse(response))
}

// DeleteLocation godoc
//
// @Summary		Delete a Location
// @Description	delete a Location by id
// @Tags			Locations
// @Accept			json
// @Produce		json
// @Param			locationId	path		uint				true	"Location ID"
// @Success		200		{object}	utils.Response		"Location deleted"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/locations/{locationId} [delete]
// @Security		BearerAuth
func (dc *LocationController) Delete(c *fiber.Ctx) error {
	Id := c.Params("locationId")
	locationId, _ := strconv.Atoi(Id)

	err := dc.LocationService.Delete(uint(locationId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Your Location has been successfully deleted", nil)
}
