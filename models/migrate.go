package models

import "sample/db"

func Migrate() {
	db.DB.AutoMigrate(
		&User{},
		&UserProfile{},
		&UserKeyPair{},
		&UserInbox{},
		&RemoteUserPublicKey{},
	)
}
