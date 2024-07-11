package router

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/roihan12/technical-test-kalbe/app/config"
	"github.com/roihan12/technical-test-kalbe/app/middleware"
	aHandler "github.com/roihan12/technical-test-kalbe/features/absent/handler"
	dHandler "github.com/roihan12/technical-test-kalbe/features/department/handler"
	eHandler "github.com/roihan12/technical-test-kalbe/features/employee/handler"
	lHandler "github.com/roihan12/technical-test-kalbe/features/location/handler"
	pHandler "github.com/roihan12/technical-test-kalbe/features/position/handler"
	rHandler "github.com/roihan12/technical-test-kalbe/features/report/handler"
)

// Router is a wrapper for HTTP router
type Router struct {
	*fiber.App
}

// NewRouter creates a new HTTP router
func NewRouter(
	app config.AppConfig,
	departmentHandler dHandler.DepartmentController,
	positionsHandler pHandler.PositionController,
	locationsHandler lHandler.LocationController,
	employeeHandler eHandler.EmployeeController,
	attendanceHandler aHandler.AttendanceController,
	reportHandler rHandler.AttendanceReportController,
) (*Router, error) {
	// Create a new Fiber app
	fiberApp := fiber.New()

	// CORS
	fiberApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Logging and recovery middleware
	fiberApp.Use(func(c *fiber.Ctx) error {
		slog.Info("Request: ", c.Method(), c.Path())
		return c.Next()
	})
	fiberApp.Use(recover.New())

	// Swagger
	fiberApp.Get("/docs/*", swagger.HandlerDefault)

	// ROUTE
	v1 := fiberApp.Group("/v1")
	employees := v1.Group("/employees")

	// Unauthenticated routes
	employees.Post("/register", employeeHandler.Register)
	employees.Post("/login", employeeHandler.Login)

	// Authenticated routes using middleware
	authUser := employees.Group("/").Use(middleware.AuthMiddleware)
	{
		authUser.Get("/", employeeHandler.Profile)
		authUser.Put("/", employeeHandler.Update)
		authUser.Delete("/", employeeHandler.Delete)
	}
	departments := v1.Group("/departments").Use(middleware.AuthMiddleware)
	{
		departments.Post("/", departmentHandler.Create)
		departments.Get("/", departmentHandler.GetAll)
		departments.Get("/:departmentId", departmentHandler.GetById)
		departments.Put("/:departmentId", departmentHandler.Update)
		departments.Delete("/:departmentId", departmentHandler.Delete)
	}
	locations := v1.Group("/locations").Use(middleware.AuthMiddleware)
	{
		locations.Post("/", locationsHandler.Create)
		locations.Get("/", locationsHandler.GetAll)
		locations.Get("/:locationId", locationsHandler.GetById)
		locations.Put("/:locationId", locationsHandler.Update)
		locations.Delete("/:locationId", locationsHandler.Delete)
	}
	positions := v1.Group("/positions").Use(middleware.AuthMiddleware)
	{
		positions.Post("/", positionsHandler.Create)
		positions.Get("/", positionsHandler.GetAll)
		positions.Get("/:positionId", positionsHandler.GetById)
		positions.Put("/:positionId", positionsHandler.Update)
		positions.Delete("/:positionId", positionsHandler.Delete)
	}

	attendances := v1.Group("/attendances").Use(middleware.AuthMiddleware)
	{
		attendances.Post("/", attendanceHandler.Create)
		attendances.Get("/", attendanceHandler.GetAll)
		attendances.Get("/:attendanceId", attendanceHandler.GetById)
		attendances.Put("/:attendanceId", attendanceHandler.Update)
		attendances.Delete("/:attendanceId", attendanceHandler.Delete)
	}
	reportAttendances := v1.Group("/attendance/reports").Use(middleware.AuthMiddleware)
	{

		reportAttendances.Get("/", reportHandler.GetAttendanceReports)

	}

	return &Router{
		App: fiberApp,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Listen(listenAddr)
}
