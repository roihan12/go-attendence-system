package database

import (
	"fmt"

	absent "github.com/roihan12/technical-test-kalbe/features/absent/data"
	department "github.com/roihan12/technical-test-kalbe/features/department/data"
	employee "github.com/roihan12/technical-test-kalbe/features/employee/data"
	location "github.com/roihan12/technical-test-kalbe/features/location/data"
	position "github.com/roihan12/technical-test-kalbe/features/position/data"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	err := db.AutoMigrate(
		&department.Department{},
		&position.Position{},
		&employee.Employee{},
		&location.Location{},
		&absent.Attendance{},
	)

	if err != nil {
		panic("Error Migration")
	}
	fmt.Println("Migration Done")
}
