package request

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type CreateCategoryRequest struct {
	Code    	string              `json:"code" validate:"required"`
	Libelle    	string              `json:"libelle" validate:"required"`
}

type UpdateCategoryRequest struct {
	ID       	string 				`bson:"_id" validate:"required"`
	Code    	string              `json:"code" validate:"required"`
	Libelle    	string              `json:"libelle" validate:"required"`
}

// Validate the request
func (r CreateCategoryRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}
	return nil
}

func (r UpdateCategoryRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}
	return nil
}

