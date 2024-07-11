package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/features/report"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type AttendanceReportController struct {
	AttendanceReportService report.AttendanceReportService
}

func New(service report.AttendanceReportService) *AttendanceReportController {
	return &AttendanceReportController{
		AttendanceReportService: service,
	}
}

// GetAttendanceReports godoc
//
// @Summary		Get Attendance Reports
// @Description	Get attendance reports within a time range
// @Tags			Attendance
// @Accept			json
// @Produce		json
// @Param			startTime	query		string	true	"Start time (RFC3339 format)"
// @Param			endTime		query		string	true	"End time (RFC3339 format)"
// @Success		200		{array}		AttendanceReportResponse	"Attendance reports"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/attendance/reports [get]
func (ctrl *AttendanceReportController) GetAttendanceReports(c *fiber.Ctx) error {
	startTimeString := c.Query("startTime")
	endTimeString := c.Query("endTime")

	startTime, err := time.Parse(time.RFC3339, startTimeString)
	if err != nil {
		return utils.ValidationError(c, err)
	}

	endTime, err := time.Parse(time.RFC3339, endTimeString)
	if err != nil {
		return utils.ValidationError(c, err)
	}

	reports, err := ctrl.AttendanceReportService.GenerateReport(startTime, endTime)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Get attendance reports successfully", ListAttendanceEntityToResponse(reports))
}
