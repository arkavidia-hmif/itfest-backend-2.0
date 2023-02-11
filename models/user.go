package models

import (
	"gorm.io/gorm"
	"itfest-backend-2.0/types"
)

type User struct {
	gorm.Model
	Usercode string                `gorm:"not null;unique"`
	Name     string                `gorm:"not null"`
	Username string                `gorm:"not null;unique"`
	Password types.EncryptedString `gorm:"not null"`
	Role     types.Role            `gorm:"not null"`
	Point    uint                  `gorm:"not null"`
}
