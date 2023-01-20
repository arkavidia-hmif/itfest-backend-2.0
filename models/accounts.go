package models

import "itfest-backend-2.0/types"

type Accounts struct {
	Id       string                `gorm:"primaryKey"`
	Username string                `gorm:"not null;unique"`
	Password types.EncryptedString `gorm:"not null"`
	Role     types.Role            `gorm:"not null"`
}
