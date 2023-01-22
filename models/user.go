package models

type User struct {
	UserId  string
	Point   int32
	// Profile Profile `gorm:"foreignKey:UserId;references:UserId"`
}
