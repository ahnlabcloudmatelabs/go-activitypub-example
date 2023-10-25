package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string      `gorm:"primaryKey"`
	Profile   UserProfile `gorm:"foreignKey:ID"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
