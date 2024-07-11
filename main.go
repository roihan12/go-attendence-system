package main

import (
	"fmt"
	"log/slog"
	"os"

	_ "github.com/roihan12/technical-test-kalbe/docs"

	dData "github.com/roihan12/technical-test-kalbe/features/department/data"
	dHandler "github.com/roihan12/technical-test-kalbe/features/department/handler"
	dService "github.com/roihan12/technical-test-kalbe/features/department/service"

	lData "github.com/roihan12/technical-test-kalbe/features/location/data"
	lHandler "github.com/roihan12/technical-test-kalbe/features/location/handler"
	lService "github.com/roihan12/technical-test-kalbe/features/location/service"

	pData "github.com/roihan12/technical-test-kalbe/features/position/data"
	pHandler "github.com/roihan12/technical-test-kalbe/features/position/handler"
	pService "github.com/roihan12/technical-test-kalbe/features/position/service"

	eData "github.com/roihan12/technical-test-kalbe/features/employee/data"
	eHandler "github.com/roihan12/technical-test-kalbe/features/employee/handler"
	eService "github.com/roihan12/technical-test-kalbe/features/employee/service"

	aData "github.com/roihan12/technical-test-kalbe/features/absent/data"
	aHandler "github.com/roihan12/technical-test-kalbe/features/absent/handler"
	aService "github.com/roihan12/technical-test-kalbe/features/absent/service"

	rData "github.com/roihan12/technical-test-kalbe/features/report/data"
	rHandler "github.com/roihan12/technical-test-kalbe/features/report/handler"
	rService "github.com/roihan12/technical-test-kalbe/features/report/service"

	"github.com/roihan12/technical-test-kalbe/app/config"
	"github.com/roihan12/technical-test-kalbe/app/database"
	"github.com/roihan12/technical-test-kalbe/app/router"
)

// @title			CRM API
// @version		1.0
// @description	This is a simple RESTful CRM Service API written in Go using fiber web framework, PostgreSQL database
//
// @contact.name	Roihan Sori
// @contact.url
// @contact.email	roihansori12@gmail.com
//
// @license.name	MIT
// @license.url
//
// @host
// @BasePath					/v1
// @schemes					http https
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the access token.
func main() {
	cfg, _ := config.NewEnv()

	db := database.InitDBPostgres(*cfg)
	database.InitMigration(db)

	// SETUP DOMAIN
	departmentData := dData.New(db)
	departmentService := dService.New(departmentData)
	departmentHandler := dHandler.New(departmentService)

	positionData := pData.New(db)
	positionService := pService.New(positionData)
	positionHandler := pHandler.New(positionService)

	locationData := lData.New(db)
	locationService := lService.New(locationData)
	locationHandler := lHandler.New(locationService)

	employeeData := eData.New(db)
	employeeService := eService.New(employeeData)
	employeeHandler := eHandler.New(employeeService)

	absentData := aData.New(db)
	absentService := aService.New(absentData)
	absentHandler := aHandler.New(absentService)

	reportData := rData.New(db)
	reportService := rService.New(reportData)
	reportHandler := rHandler.New(reportService)

	// Init router
	router, err := router.NewRouter(
		*cfg,
		*departmentHandler,
		*positionHandler,
		*locationHandler,
		*employeeHandler,
		*absentHandler,
		*reportHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	listenAddr := fmt.Sprintf("%s:%s", cfg.URL, cfg.Port)
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
