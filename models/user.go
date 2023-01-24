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
	Point    int32                 `gorm:"not null"`
	Profile  Profile               `gorm:"foreignKey:UserID"`
	LogsFrom []Log                 `gorm:"foreignKey:From"`
	LogsTo   []Log                 `gorm:"foreignKey:To"`
	Game     Game                  `gorm:"foreignKey:UserID"`
}
