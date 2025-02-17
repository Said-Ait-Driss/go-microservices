package request

import (
	"fmt"
	entity "review-service/pkg/entities"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProductReviewsRequest struct {
	ID string `json:"_id"`
}

type CreateNewProductReviewRequest struct {
	Product_title string             `json:"product_title" validate:"required"`
	Product_id    primitive.ObjectID `json:"product_id" validate:"required"`
	Review        entity.Review      `json:"review" validate:"required"`
}

// Validate the request
func (r CreateNewProductReviewRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r.Review); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	return nil
}
