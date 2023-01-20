package types

import (
	"database/sql/driver"
)

type Role string

const (
	Admin   Role = "admin"
	User    Role = "user"
	Startup Role = "startup"
)

func (role *Role) Scan(value interface{}) error {
	*role = Role(value.(string))
	return nil
}

func (role Role) Value() (driver.Value, error) {
	return string(role), nil
}

func (Role) GormDataType() string {
	return "string"
}
