package request

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetCategoryRequest struct {
	ID string `json:"_id"`
}

type CreateCategoryRequest struct {
	ID         primitive.ObjectID `bson:"_id"`
	Title      string             `json:"title" validate:"required"`
	Image      string             `json:"image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type UpdateCategoryRequest struct {
	ID primitive.ObjectID `bson:"_id" validate:"required"`
	CreateCategoryRequest
}

type DeleteCategoryRequest struct {
	ID string `json:"_id"`
}

type GetCategoriesRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

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

	return r.CreateCategoryRequest.Validate()
}
