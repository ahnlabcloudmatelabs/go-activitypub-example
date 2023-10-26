package models

import "time"

type UserFollower struct {
	ID        string `gorm:"primaryKey;uniqueIndex:idx_user_follower;"`
	Follower  string `gorm:"type:text;uniqueIndex:idx_user_follower;"`
	CreatedAt time.Time
}
