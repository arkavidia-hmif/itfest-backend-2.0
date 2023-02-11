package models

import (
	"gopkg.in/guregu/null.v4"
	"itfest-backend-2.0/types"
)

type Profile struct {
	UserID    uint                  `gorm:"not null" json:"user_id"`
	Email     null.String           `gorm:"default:null" json:"email"`
	BirthDate types.BirthDate       `gorm:"default:null" json:"birthdate"`
	Gender    null.String           `gorm:"default:null" json:"gender"`
	Interests types.CareerInterests `gorm:"type:string;default:''" json:"interests"`
	Submitted bool                  `gorm:"default:false" json:"submitted"`
}
