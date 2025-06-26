package request

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func (u *UserRequest) Validate() error {
	return validate.Struct(u)
}

func (u *LoginRequest) Validate() error {
	return validate.Struct(u)
}

func (u *ResetUser) Validate() error {
	return validate.Struct(u)
}

func (u *VerifyToken) Validate() error {
	return validate.Struct(u)
}

func (u *NewPassword) Validate() error {
	return validate.Struct(u)
}
