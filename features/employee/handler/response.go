package handler

import "github.com/roihan12/technical-test-kalbe/features/employee"

type EmployeeResponse struct {
	ID           uint   `json:"id,omitempty"`
	EmployeeCode string `json:"employee_code,omitempty"`
	EmployeeName string `json:"employee_name,omitempty"`
	Email        string `json:"email,omitempty"`
	DepartmentID uint   `json:"department_id,omitempty"`
	PositionID   uint   `json:"position_id,omitempty"`
	Superior     uint   `json:"superior,omitempty"`
}

// authResponse represents an authentication response body
type authResponse struct {
	AccessToken string `json:"token" example:"eGdh5kiOTyyaQ3_bNykYDeYHO21Jg2..."`
}

// newAuthResponse is a helper function to create a response body for handling authentication data
func newAuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

func ToResponse(data employee.EmployeeEntity) EmployeeResponse {
	return EmployeeResponse{
		ID:           data.EmployeeID,
		EmployeeCode: data.EmployeeCode,
		EmployeeName: data.EmployeeName,
		DepartmentID: data.DepartmentID,
		PositionID:   data.PositionID,
		Superior:     data.Superior,
	}
}

func GetToResponse(data employee.EmployeeEntity) EmployeeResponse {
	return EmployeeResponse{
		ID:           data.EmployeeID,
		EmployeeCode: data.EmployeeCode,
		EmployeeName: data.EmployeeName,
		DepartmentID: data.DepartmentID,
		PositionID:   data.PositionID,
		Superior:     data.Superior,
	}
}
