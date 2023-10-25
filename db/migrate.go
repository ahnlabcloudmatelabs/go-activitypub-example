package db

import "sample/db/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.UserProfile{},
		&models.UserKeyPair{},
		&models.UserInbox{},
	)
}
