package utils

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Response represents a response body format
type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Success"`
	Data    any    `json:"data,omitempty"`
}

// NewResponse is a helper function to create a response body
func NewResponse(success bool, message string, data any) Response {
	return Response{
		Success: success,
		Message: message,
		Data:    data,
	}
}

// Meta represents metadata for a paginated response
type Meta struct {
	Total uint64 `json:"total" example:"100"`
	Limit uint64 `json:"limit" example:"10"`
	Skip  uint64 `json:"skip" example:"0"`
}

// NewMeta is a helper function to create metadata for a paginated response
func NewMeta(total, limit, skip uint64) Meta {
	return Meta{
		Total: total,
		Limit: limit,
		Skip:  skip,
	}
}

// ErrorStatusMap is a map of defined error messages and their corresponding http status codes
var ErrorStatusMap = map[error]int{
	ErrInternal:                   http.StatusInternalServerError,
	ErrDataNotFound:               http.StatusNotFound,
	ErrConflictingData:            http.StatusConflict,
	ErrDuplicateData:              http.StatusConflict,
	ErrInvalidCredentials:         http.StatusUnauthorized,
	ErrUnauthorized:               http.StatusUnauthorized,
	ErrEmptyAuthorizationHeader:   http.StatusUnauthorized,
	ErrInvalidAuthorizationHeader: http.StatusUnauthorized,
	ErrInvalidAuthorizationType:   http.StatusUnauthorized,
	ErrInvalidToken:               http.StatusUnauthorized,
	ErrExpiredToken:               http.StatusUnauthorized,
	ErrForbidden:                  http.StatusForbidden,
	ErrNoUpdatedData:              http.StatusBadRequest,
	ErrFormatFile:                 http.StatusBadRequest,
	ErrFileSize:                   http.StatusBadRequest,
	ErrFileExtension:              http.StatusBadRequest,
	ErrImageRequired:              http.StatusBadRequest,
}

// ValidationError sends an error response for some specific request validation error
func ValidationError(ctx *fiber.Ctx, err error) error {
	errMsgs := parseError(err)
	errRsp := NewErrorResponse(errMsgs)
	return ctx.Status(http.StatusBadRequest).JSON(errRsp)
}

// HandleError determines the status code of an error and returns a JSON response with the error message and status code
func HandleError(ctx *fiber.Ctx, err error) error {
	statusCode, ok := ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := NewErrorResponse(errMsg)
	return  ctx.Status(statusCode).JSON(errRsp)
}

// HandleAbort sends an error response and aborts the request with the specified status code and error message
func HandleAbort(ctx *fiber.Ctx, err error) error {
	statusCode, ok := ErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := NewErrorResponse(errMsg)
	return ctx.Status(statusCode).JSON(errRsp)
}

// ParseError parses error messages from the error object and returns a slice of error messages
func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, ErrorRequestMessages(err.Field(), err.ActualTag(), err.Param()))
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

func ErrorRequestMessages(field, tag, param string) (message string) {
	switch tag {
	case "required":
		message = fmt.Sprintf("%s is required", field)
	case "email":
		message = fmt.Sprintf("%s must be valid email", field)
	case "gt":
		message = fmt.Sprintf("%s must be grather than %s", field, param)
	case "min":
		message = fmt.Sprintf("%s minimum %s characters", field, param)
	default:
		message = "Error validated request"
	}

	return
}

// ErrorResponse represents an error response body format
type ErrorResponse struct {
	Success  bool     `json:"success" example:"false"`
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

// NewErrorResponse is a helper function to create an error response body
func NewErrorResponse(errMsgs []string) ErrorResponse {
	return ErrorResponse{
		Success:  false,
		Messages: errMsgs,
	}
}

// HandleSuccess sends a success response with the specified status code and optional data
func HandleSuccess(ctx *fiber.Ctx, message string, data any) error {
	rsp := NewResponse(true, message, data)
	return ctx.Status(http.StatusOK).JSON(rsp)
}
