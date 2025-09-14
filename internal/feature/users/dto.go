package users

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required,min=2"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=admin user"`
}

type CreateUserResponseDTO struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	ErpKey string `json:"erp_key"`
}

type UpdateUserDTO struct {
	Name string `json:"name" validate:"omitempty,min=2"`
	Role string `json:"role" validate:"omitempty,oneof=admin user"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (d *CreateUserDTO) Validate() error { return validate.Struct(d) }
func (d *UpdateUserDTO) Validate() error { return validate.Struct(d) }
func (d *LoginDTO) Validate() error      { return validate.Struct(d) }
