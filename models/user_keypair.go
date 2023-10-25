package models

import "sample/db"

type UserKeyPair struct {
	ID         string `gorm:"primaryKey"`
	PublicKey  string `gorm:"not null;type:text"`
	PrivateKey string `gorm:"not null;type:text"`
}

func (u *UserKeyPair) GetByID() error {
	return db.DB.First(&u, "id = ?", u.ID).Error
}
