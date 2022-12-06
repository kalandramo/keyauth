package user

import "github.com/go-playground/validator/v10"

var validate = validator.New()

func (q *QueryAccountRequest) Validate() error {
	return nil
}

func (c *CreateAccountRequest) Validate() error {
	return validate.Struct(c)
}
