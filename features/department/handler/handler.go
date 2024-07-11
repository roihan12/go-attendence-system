package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/features/department"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type DepartmentController struct {
	departmentService department.DepartmentService
}

func New(srv department.DepartmentService) *DepartmentController {
	return &DepartmentController{
		departmentService: srv,
	}
}

// GetAll godoc
//
// @Summary		List departments
// @Description	List departments
// @Tags			Departments
// @Accept			json
// @Produce		json
// @Success		200	{array}		DepartmentResponse	"Departments retrieved"
// @Failure		500	{object}	utils.ErrorResponse	"Internal server error"
// @Router			/departments [get]
// @Security		BearerAuth
func (dc *DepartmentController) GetAll(c *fiber.Ctx) error {

	departments, err := dc.departmentService.GetAll()
	if err != nil {
		return utils.HandleError(c, err)
	}

	listDepartmentResponse := ListDepartmentEntityToDepartmentResponse(departments)
	return utils.HandleSuccess(c, "Get all departments successfully", listDepartmentResponse)
}

// GetById godoc
//
// @Summary		Get a department
// @Description	get a department by id
// @Tags			Departments
// @Accept			json
// @Produce		json
// @Param			departmentId	path		uint				true	"Department ID"
// @Success		200		{object}	DepartmentResponse	"Department retrieved"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/departments/{departmentId} [get]
// @Security		BearerAuth
func (dc *DepartmentController) GetById(c *fiber.Ctx) error {
	Id := c.Params("departmentId")
	departmentId, _ := strconv.Atoi(Id)
	department, err := dc.departmentService.GetById(uint(departmentId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	response := DepartmentEntityToDepartmentResponse(department)
	return utils.HandleSuccess(c, "Get by id department successfully", response)
}

// CreateDepartment godoc
//
// @Summary		Create a new department
// @Description	create a new department with title and description
// @Tags			Departments
// @Produce		json
// @Param			CreateDepartmentRequest	formData	CreateDepartmentRequest	true	"Create department request"
// @Success		200					{object}	DepartmentResponse		"Department created"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/departments [post]
// @Security		BearerAuth
func (dc *DepartmentController) Create(c *fiber.Ctx) error {
	var req CreateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	departmentEntity := CreateDepartmentRequestToDepartmentEntity(&req)

	response, err := dc.departmentService.Create(departmentEntity)
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Create department successfully", DepartmentEntityToDepartmentResponse(response))
}

// UpdateDepartment godoc
//
// @Summary		Update a department
// @Description	update a department with title and description
// @Tags			Departments
// @Produce		json
// @Param			departmentId			path		uint				true	"Department ID"
// @Param			UpdateDepartmentRequest	formData	UpdateDepartmentRequest	true	"Update department request"
// @Success		200					{object}	DepartmentResponse		"Department updated"
// @Failure		400					{object}	utils.ErrorResponse	"Validation error"
// @Failure		404					{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500					{object}	utils.ErrorResponse	"Internal server error"
// @Router			/departments/{departmentId} [put]
// @Security		BearerAuth
func (dc *DepartmentController) Update(c *fiber.Ctx) error {

	Id := c.Params("departmentId")
	departmentId, _ := strconv.Atoi(Id)
	var req UpdateDepartmentRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	departmentEntity := UpdateDepartmentRequestToDepartmentEntity(&req)

	response, err := dc.departmentService.Update(departmentEntity, uint(departmentId))
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Update department successfully", DepartmentEntityToDepartmentResponse(response))
}

// DeleteDepartment godoc
//
// @Summary		Delete a department
// @Description	delete a department by id
// @Tags			Departments
// @Accept			json
// @Produce		json
// @Param			departmentId	path		uint				true	"Department ID"
// @Success		200		{object}	utils.Response		"Department deleted"
// @Failure		400		{object}	utils.ErrorResponse	"Validation error"
// @Failure		404		{object}	utils.ErrorResponse	"Data not found error"
// @Failure		500		{object}	utils.ErrorResponse	"Internal server error"
// @Router			/departments/{departmentId} [delete]
// @Security		BearerAuth
func (dc *DepartmentController) Delete(c *fiber.Ctx) error {
	Id := c.Params("departmentId")
	departmentId, _ := strconv.Atoi(Id)

	err := dc.departmentService.Delete(uint(departmentId))
	if err != nil {
		return utils.HandleError(c, err)
	}
	return utils.HandleSuccess(c, "Your department has been successfully deleted", nil)
}
