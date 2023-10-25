package models

import (
	"sample/db"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string      `gorm:"primaryKey"`
	Profile   UserProfile `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	KeyPair   UserKeyPair `gorm:"foreignKey:ID;constraint:OnDelete:CASCADE;"`
	Inbox     UserInbox   `gorm:"foreignKey:To;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (u User) Exists() bool {
	var count int64
	db.DB.Model(&User{}).Where("id = ?", u.ID).Count(&count)
	return count > 0
}
