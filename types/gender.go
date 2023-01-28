package types

import (
	"database/sql/driver"
	"errors"
)

type Gender struct {
	String string
	Valid  bool
}

const (
	Male   string = "male"
	Female string = "female"
)

func (gender *Gender) Scan(input interface{}) error {
	value := input.(string)

	if input == "" || input == nil {
		gender.String = ""
		gender.Valid = false
	} else if !(value == Male || value == Female) {
		return errors.New("invalid gender. Should be male or female")
	} else {
		gender.String = value
		gender.Valid = true
	}

	return nil
}

func (gender Gender) Value() (driver.Value, error) {
	if gender.Valid {
		return gender.String, nil
	}

	return nil, nil
}

func (Gender) GormDataType() string {
	return "string"
}
