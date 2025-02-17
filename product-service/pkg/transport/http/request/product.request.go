package request

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetProductRequest struct {
	ID string `json:"_id"`
}

type GetProductsRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type GetProductsOfStoreThatHasOffersRequest struct {
	Store_id string `json:"store_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}

type GetStoreProductsFilterRequest struct {
	Store_id  string    `json:"store_id"`
	Name      string    `json:"name"`
	MinPrice  float64   `json:"minPrice"`
	MaxPrice  float64   `json:"maxPrice"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Page      int       `json:"page"`
	PageSize  int       `json:"pageSize"`
}

type GetProductsFilterRequest struct {
	Name         string    `json:"name"`
	MinPrice     float64   `json:"minPrice"`
	MaxPrice     float64   `json:"maxPrice"`
	StartDate    time.Time `json:"startDate"`
	EndDate      time.Time `json:"endDate"`
	Neighborhood string    `json:"neighborhood"`
	StoreTitle   string    `json:"storeTitle"`
	Page         int       `json:"page"`
	PageSize     int       `json:"pageSize"`
}

type GetProductByCategoryRequest struct {
	CategoryId string `json:"category_id"`
}

type GetProductsByStoreRequest struct {
	Store_id string `json:"store_id"`
}

type CreateProductRequest struct {
	Name           string  `json:"name" validate:"required"`
	Original_price float64 `json:"original_price" validate:"required"`
	Store          struct {
		Store_id string `json:"store_id" validate:"required"`
		Title    string `json:"title" validate:"required"`
	}
	Image         string   `json:"image" validate:"required"`
	Images        []string `json:"images"`
	Quantity      int      `json:"quantity" validate:"required"`
	Current_offer struct {
		Offer_id string  `json:"offer_id" validate:"required"`
		Price    float64 `json:"price" validate:"required,numeric"`
		Discount float64 `json:"discount" validate:"required,numeric"`
	}
}

type UpdateProductRequest struct {
	ID primitive.ObjectID `bson:"_id" validate:"required"`
	CreateProductRequest
}

type DeleteProductRequest struct {
	ID string `json:"_id" validate:"required"`
}

// getting products count of a store
type GetProductsCountByStoreRequest struct {
	Store_id string `json:"store_id" validate:"required"`
}

type ValidateProductRequest interface {
	Validate() error
}

// Validate the request
func (r CreateProductRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r.Current_offer); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	if err := validate.Struct(r.Store); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	return nil
}

func (r UpdateProductRequest) Validate() error {
	validate := validator.New()

	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field '%s' '%s'", _err.Field(), _err.Tag())
	}

	return r.CreateProductRequest.Validate()
}
