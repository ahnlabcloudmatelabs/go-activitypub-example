package models

import (
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
