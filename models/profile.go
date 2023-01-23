package models

import (
	"time"

	"itfest-backend-2.0/types"
)

type Profile struct {
	UserId    string `gorm:"not null"`
	Email     string
	BirthDate time.Time
	Gender    string
	Interests types.CareerInterest
	Submitted bool
}
