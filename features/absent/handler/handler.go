package handler

import (
	"strconv"
	"github.com/gofiber/fiber/v2"
	attendance "github.com/roihan12/technical-test-kalbe/features/absent"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type AttendanceController struct {
	AttendanceService attendance.AttendanceService
}

func New(srv attendance.AttendanceService) *AttendanceController {
	return &AttendanceController{
		AttendanceService: srv,
	}
}

// GetAll godoc
//
// @Summary		List Attendances
// @Description	List Attendances
// @Tags			Attendances
// @Accept			json
// @Produce		json
// @Success		200	{array}		AttendanceResponse	"Attendances retrieved"
// @Failure		500	{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendances [get]
// @Security		BearerAuth
func (ac *AttendanceController) GetAll(c *fiber.Ctx) error {
	attendances, err := ac.AttendanceService.GetAll()
	if err != nil {
		return utils.HandleError(c, err)
	}

	response := ListAttendanceEntityToAttendanceResponse(attendances)
	return utils.HandleSuccess(c, "Get all Attendances successfully", response)
}

// GetById godoc
//
// @Summary		Get an Attendance
// @Description	get an Attendance by id
// @Tags			Attendances
// @Accept			json
// @Produce		json
// @Param			attendanceId	path		uint				true	"Attendance ID"
// @Success		200		{object}	AttendanceResponse	"Attendance retrieved"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendances/{attendanceId} [get]
// @Security		BearerAuth
func (ac *AttendanceController) GetById(c *fiber.Ctx) error {
	id := c.Params("attendanceId")
	attendanceID, _ := strconv.Atoi(id)

	attendance, err := ac.AttendanceService.GetById(uint(attendanceID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	response := AttendanceEntityToAttendanceResponse(attendance)
	return utils.HandleSuccess(c, "Get Attendance by id successfully", response)
}

// CreateAttendance godoc
//
// @Summary		Create a new Attendance
// @Description	create a new Attendance with employee_id, location_id, absent_in, and absent_out
// @Tags			Attendances
// @Produce		json
// @Param			CreateAttendanceRequest	formData	CreateAttendanceRequest	true	"Create Attendance request"
// @Success		200					{object}	AttendanceResponse		"Attendance created"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendances [post]
// @Security		BearerAuth
func (ac *AttendanceController) Create(c *fiber.Ctx) error {
	var req CreateAttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	attendanceEntity := CreateAttendanceRequestToAttendanceEntity(&req)

	response, err := ac.AttendanceService.Create(attendanceEntity)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Create Attendance successfully", AttendanceEntityToAttendanceResponse(response))
}

// UpdateAttendance godoc
//
// @Summary		Update an Attendance
// @Description	update an Attendance with employee_id, location_id, absent_in, and absent_out
// @Tags			Attendances
// @Produce		json
// @Param			attendanceId			path		uint				true	"Attendance ID"
// @Param			UpdateAttendanceRequest	formData	UpdateAttendanceRequest	true	"Update Attendance request"
// @Success		200					{object}	AttendanceResponse		"Attendance updated"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendances/{attendanceId} [put]
// @Security		BearerAuth
func (ac *AttendanceController) Update(c *fiber.Ctx) error {
	id := c.Params("attendanceId")
	attendanceID, _ := strconv.Atoi(id)

	var req UpdateAttendanceRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	attendanceEntity := UpdateAttendanceRequestToAttendanceEntity(&req)

	response, err := ac.AttendanceService.Update(attendanceEntity, uint(attendanceID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Update Attendance successfully", AttendanceEntityToAttendanceResponse(response))
}

// DeleteAttendance godoc
//
// @Summary		Delete an Attendance
// @Description	delete an Attendance by id
// @Tags			Attendances
// @Accept			json
// @Produce		json
// @Param			attendanceId	path		uint				true	"Attendance ID"
// @Success		200		{object}	utils.Response		"Attendance deleted"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendances/{attendanceId} [delete]
// @Security		BearerAuth
func (ac *AttendanceController) Delete(c *fiber.Ctx) error {
	id := c.Params("attendanceId")
	attendanceID, _ := strconv.Atoi(id)

	err := ac.AttendanceService.Delete(uint(attendanceID))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Attendance has been successfully deleted", nil)
}
