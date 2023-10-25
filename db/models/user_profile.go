package models

import "time"

type UserProfile struct {
	ID        string  `gorm:"primaryKey"`
	Name      string  `gorm:"index,unique,not null"`
	Bio       *string `gorm:"type:text"`
	Icon      *string `gorm:"type:text"`
	Image     *string `gorm:"type:text"`
	UpdatedAt time.Time
}
