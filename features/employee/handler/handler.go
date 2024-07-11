package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/features/employee"
	"github.com/roihan12/technical-test-kalbe/utils"
)

type EmployeeController struct {
	srv employee.EmployeeService
}

func New(srv employee.EmployeeService) *EmployeeController {
	return &EmployeeController{
		srv: srv,
	}
}

// Login godoc
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered employee and returns an access token if the credentials are valid.
//	@Tags			Employees
//	@Accept			json
//	@Produce		json
//	@Param			request	body		LoginRequest		true	"Login request body"
//	@Success		200		{object}	authResponse		"Succesfully logged in"
//	@Failure		400		{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401		{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		500		{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/employees/login [post]
func (ec *EmployeeController) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	token, err := ec.srv.Login(req.EmployeeName, req.Password)
	if err != nil {
		return utils.HandleError(c, err)
	}

	rsp := newAuthResponse(token)
	return utils.HandleSuccess(c, "Login employee successfully", rsp)
}

// Register godoc
//
//	@Summary		Register a new employee
//	@Description	create a new employee account
//	@Tags			Employees
//	@Accept			json
//	@Produce		json
//	@Param			registerRequest	body		RegisterRequest		true	"Register request"
//	@Success		200				{object}	EmployeeResponse	"Employee created"
//	@Failure		400				{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401				{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		404				{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		409				{object}	utils.ErrorResponse	"Data conflict error"
//	@Failure		500				{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/employees/register [post]
func (ec *EmployeeController) Register(c *fiber.Ctx) error {
	req := RegisterRequest{}
	if err := c.BodyParser(&req); err != nil {
		return utils.ValidationError(c, err)
	}

	res, err := ec.srv.Register(*ReqToCore(req))
	if err != nil {
		return utils.HandleError(c, err)
	}

	rsp := ToResponse(res)
	return utils.HandleSuccess(c, "Register employee successfully", rsp)
}

// ProfileEmployee godoc
//
//	@Summary		Get an employee
//	@Description	Get a profile employee
//	@Tags			Employees
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	EmployeeResponse	"Profile Employee displayed"
//	@Failure		400	{object}	utils.ErrorResponse	"Validation error"
//	@Failure		404	{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/employees [get]
//	@Security		BearerAuth
func (ec *EmployeeController) Profile(c *fiber.Ctx) error {
	userID, userIDIsExist := c.Locals("employeeID").(uint)
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		return utils.HandleAbort(c, err)
	}

	res, err := ec.srv.Profile(userID)
	if err != nil {
		return utils.HandleError(c, err)
	}

	rsp := ToResponse(res)
	return utils.HandleSuccess(c, "Get profile employee successfully", rsp)
}

// UpdateEmployee godoc
//
//	@Summary		Update an employee
//	@Description	Update an employee's username, email, password, age
//	@Tags			Employees
//	@Accept			json
//	@Produce		json
//	@Param			updateEmployeeRequest	body		UpdateRequest		true	"Update employee request"
//	@Success		200						{object}	EmployeeResponse	"Employee updated"
//	@Failure		400						{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401						{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403						{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404						{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500						{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/employees [put]
//	@Security		BearerAuth
func (ec *EmployeeController) Update(c *fiber.Ctx) error {
	userID, userIDIsExist := c.Locals("employeeID").(uint)
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		return utils.HandleAbort(c, err)
	}

	input := UpdateRequest{}
	if err := c.BodyParser(&input); err != nil {
		return utils.ValidationError(c, err)
	}

	res, err := ec.srv.Update(userID, *ReqToCore(input))
	if err != nil {
		return utils.HandleError(c, err)
	}

	rsp := ToResponse(res)
	return utils.HandleSuccess(c, "Update employee successfully", rsp)
}

// DeleteEmployee godoc
//
//	@Summary		Delete an employee
//	@Description	Delete an employee
//	@Tags			Employees
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	utils.Response		"Employee deleted"
//	@Failure		400	{object}	utils.ErrorResponse	"Validation error"
//	@Failure		401	{object}	utils.ErrorResponse	"Unauthorized error"
//	@Failure		403	{object}	utils.ErrorResponse	"Forbidden error"
//	@Failure		404	{object}	utils.ErrorResponse	"Data not found error"
//	@Failure		500	{object}	utils.ErrorResponse	"Internal server error"
//	@Router			/employees [delete]
//	@Security		BearerAuth
func (ec *EmployeeController) Delete(c *fiber.Ctx) error {
	userID, userIDIsExist := c.Locals("employeeID").(uint)
	if !userIDIsExist {
		err := utils.ErrUnauthorized
		return utils.HandleAbort(c, err)
	}

	err := ec.srv.Delete(userID)
	if err != nil {
		return utils.HandleError(c, err)
	}

	return utils.HandleSuccess(c, "Your account has been deleted successfully", "")
}
