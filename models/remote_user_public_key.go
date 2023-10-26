package models

import "sample/db"

type RemoteUserPublicKey struct {
	ID        string `gorm:"primaryKey,type:text"`
	PublicKey string `gorm:"not null;type:text"`
}

func (r *RemoteUserPublicKey) GetByID() error {
	return db.DB.First(r, "id = ?", r.ID).Error
}
