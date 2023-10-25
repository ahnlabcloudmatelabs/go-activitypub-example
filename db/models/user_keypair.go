package models

type UserKeyPair struct {
	ID         string `gorm:"primaryKey"`
	PublicKey  string `gorm:"not null;type:text"`
	PrivateKey string `gorm:"not null;type:text"`
}
