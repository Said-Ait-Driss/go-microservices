package request

import (
	"fmt"
	entity "review-service/pkg/entities"

	"github.com/go-playground/validator/v10"
)

type GetStoreReviewsRequest struct {
	ID string `json:"_id"`
}

type DeleteReviewByStoreRequest struct {
	Store_id string `json:"store_id"`
}

type DeleteReviewByUserRequest struct {
	User_id string `json:"user_id"`
}

type GetReviewsOfListOfStoresRequest struct {
	Store_ids []string `json:"store_ids"`
}

type ChangeReviewStatusRequest struct {
	Store_id   string `json:"store_id" validate:"required"`
	Review_id  string `json:"review_id" validate:"required"`
	New_Status string `json:"new_Status" validate:"required"`
}

type CreateNewStoreReviewRequest struct {
	Store_title string        `json:"title" validate:"required"`
	Store_id    string        `json:"store_id" validate:"required"`
	Review      entity.Review `json:"review" validate:"required"`
}

// Validate the request
func (r CreateNewStoreReviewRequest) Validate() error {
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
