package models

import (
	"gorm.io/gorm"
	"itfest-backend-2.0/types"
)

type Account struct {
	gorm.Model
	Name        string                `gorm:"not null"`
	Username    string                `gorm:"not null;unique"`
	Password    types.EncryptedString `gorm:"not null"`
	Role        types.Role            `gorm:"not null"`
	// User        User                  `gorm:"foreignKey:UserId;references:ID"`
	// Log         Log                   `gorm:"foreignKey:From,To;references:ID"`
	// Merchandise Merchandise           `gorm:"foreignKey:MerchandiseID;references:ID"`
	// Clue        Clue                  `gorm:"foreignKey:ClueID;references:ID"`
}
