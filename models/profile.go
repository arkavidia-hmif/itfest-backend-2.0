package models

import (
	"time"

	"itfest-backend-2.0/types"
)

type Profile struct {
	UserID    int `gorm:"not null"`
	Email     string
	BirthDate time.Time
	Gender    string
	Interests types.CareerInterests `gorm:"type:string"`
	Submitted bool
}
