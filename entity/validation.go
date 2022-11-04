package entity

import "github.com/go-playground/validator/v10"

func (r *UserRegisterRequest) Validate() error {
	v := validator.New()

	if err := v.Struct(r); err != nil {
		return err
	}

	return nil
}
