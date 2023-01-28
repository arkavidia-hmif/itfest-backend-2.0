package models

import (
	"database/sql"
	"itfest-backend-2.0/types"
)

type Profile struct {
	UserID    uint                  `gorm:"not null"`
	Email     sql.NullString        `gorm:"default:null"`
	BirthDate sql.NullTime          `gorm:"default:null"`
	Gender    types.Gender          `gorm:"default:null"`
	Interests types.CareerInterests `gorm:"type:string;default:''"`
	Submitted bool                  `gorm:"default:false"`
}
