package request

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetOfferRequest struct {
	ID string `json:"_id"`
}

type CreateOfferRequest struct {
	ID         primitive.ObjectID `bson:"_id"`
	Price      float64            `json:"price" validate:"required,numeric"`
	Discount   float64            `json:"discount" validate:"required,numeric"`
	Start_date string             `json:"start_date" validate:"required,customStartDateValidator"`
	End_date   string             `json:"end_date" validate:"required,customEndDateValidator"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

type UpdateOfferRequest struct {
	ID primitive.ObjectID `bson:"_id" validate:"required"`
	CreateOfferRequest
}

type GetOffersRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type DeleteOfferRequest struct {
	ID string `json:"_id" validate:"required"`
}

// Validate the request

// CustomStartDateValidator is a custom validator for start date validation
func CustomStartDateValidator(fl validator.FieldLevel) bool {
	startDateStr := fl.Field().String()
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return false
	}
	// Example validation: Ensure start date is not in the past
	return !startDate.Before(time.Now())
}

// CustomEndDateValidator is a custom validator for end date validation
func CustomEndDateValidator(fl validator.FieldLevel) bool {
	endDateStr := fl.Field().String()
	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return false
	}
	// Example validation: Ensure end date is after the current time
	return endDate.After(time.Now())
}

func (r CreateOfferRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("customStartDateValidator", CustomStartDateValidator)
	validate.RegisterValidation("customEndDateValidator", CustomEndDateValidator)

	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	return nil
}

func (r UpdateOfferRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("customStartDateValidator", CustomStartDateValidator)
	validate.RegisterValidation("customEndDateValidator", CustomEndDateValidator)

	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	return r.CreateOfferRequest.Validate()
}
