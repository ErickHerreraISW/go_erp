package products

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type CreateProductDTO struct {
	Name        string  `json:"name" validate:"required,min=2"`
	SKU         string  `json:"sku" validate:"omitempty,alphanum"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"gte=0"`
	Description string  `json:"description" validate:"omitempty"`
}

type UpdateProductDTO struct {
	Name        string  `json:"name" validate:"omitempty,min=2"`
	SKU         string  `json:"sku" validate:"omitempty,alphanum"`
	Price       float64 `json:"price" validate:"omitempty,gt=0"`
	Stock       int     `json:"stock" validate:"omitempty,gte=0"`
	Description string  `json:"description" validate:"omitempty"`
}

func (d *CreateProductDTO) Validate() error { return validate.Struct(d) }
func (d *UpdateProductDTO) Validate() error { return validate.Struct(d) }
