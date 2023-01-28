package configs

import (
	"github.com/go-playground/validator/v10"
)

type RequestValidator struct {
	Validator *validator.Validate
}

func (rv *RequestValidator) Validate(i interface{}) error {
	if err := rv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
