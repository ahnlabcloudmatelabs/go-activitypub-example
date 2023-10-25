package db

import "sample/db/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
	)
}
