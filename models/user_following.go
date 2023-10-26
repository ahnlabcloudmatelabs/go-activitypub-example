package models

import "time"

type UserFollowing struct {
	ID        string `gorm:"primaryKey;uniqueIndex:idx_user_following;"`
	Following string `gorm:"type:text;uniqueIndex:idx_user_following;"`
	CreatedAt time.Time
}
