package db

import (
	"fmt"
	"os"
	"sample/constants"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(
		dialector(),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}
}

func dialector() gorm.Dialector {
	if os.Getenv("DB_TYPE") == "mssql" {
		dsn := fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s",
			constants.DB_USER,
			constants.DB_PASSWORD,
			constants.DB_HOST,
			constants.DB_PORT,
			constants.DB_NAME,
		)
		return sqlserver.Open(dsn)
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		constants.DB_HOST,
		constants.DB_USER,
		constants.DB_PASSWORD,
		constants.DB_NAME,
		constants.DB_PORT,
	)
	return postgres.Open(dsn)
}
