package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/roihan12/technical-test-kalbe/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	auth := c.Get("Authorization")
	if bearerIsExist := strings.Contains(auth, "Bearer"); !bearerIsExist {
		err := utils.ErrEmptyAuthorizationHeader
		return utils.HandleAbort(c, err)
	}

	token := strings.Split(auth, " ")
	if len(token) < 2 {
		err := utils.ErrInvalidAuthorizationHeader
		return utils.HandleAbort(c, err)
	}

	claims, err := utils.VerifyAccessToken(token[1])
	if err != nil {
		return utils.HandleAbort(c, err)
	}

	c.Locals("employeeName", claims.EmployeeName)
	c.Locals("employeeID", claims.EmployeeID)

	return c.Next()
}
