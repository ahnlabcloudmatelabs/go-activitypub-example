package models

import (
	"time"

	"github.com/google/uuid"
)

type UserInbox struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	From      string    `gorm:"type:text"`
	To        string    `gorm:"index"`
	Type      string
	Content   string `gorm:"type:text"`
	CreatedAt time.Time
}
