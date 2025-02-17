package request

import (
	"fmt"
	entity "store-service/pkg/entities"

	"github.com/go-playground/validator/v10"
)

type GetStoreRequest struct {
	Store_id string `json:"store_id" validate:"required"`
}

type CreateStoreRequest struct {
	Name         string            `json:"name" validate:"required"`
	Description  string            `json:"description" validate:"required"`
	Phone        string            `json:"phone" validate:"required"`
	Logo_url     string            `json:"logo_url"`
	Covers       []entity.Covers   `json:"covers"`
	Longitude    string            `json:"longitude" validate:"required"`
	Latitude     string            `json:"latitude" validate:""`
	Client_id    string            `json:"client_id" validate:"required"`
	Adresse      string            `json:"adresse" validate:"required"`
	City         string            `json:"city" validate:"required"`
	Neighborhood string            `json:"neighborhood" validate:"required"`
	Category     entity.Categories `json:"category" validate:"required"`
	Status       string            `json:"status" validate:"required"`
	Open_from    string            `json:"open_from" validate:"required"`
	Open_to      string            `json:"open_to" validate:"required"`
}

type UpdateStoreRequest struct {
	ID           string            `json:"id" validate:"required"`
	Name         string            `json:"name" validate:"required"`
	Description  string            `json:"description" validate:"required"`
	Phone        string            `json:"phone" validate:"required"`
	Covers       []entity.Covers   `json:"covers"`
	Longitude    string            `json:"longitude" validate:"required"`
	Latitude     string            `json:"latitude" validate:""`
	Client_id    string            `json:"client_id" validate:"required"`
	Adresse      string            `json:"adresse" validate:"required"`
	City         string            `json:"city" validate:"required"`
	Neighborhood string            `json:"neighborhood" validate:"required"`
	Category     entity.Categories `json:"category" validate:"required"`
	Status       string            `json:"status" validate:"required"`
	Open_from    string            `json:"open_from" validate:"required"`
	Open_to      string            `json:"open_to" validate:"required"`
}

type ChangeStoreLogoRequest struct {
	ID       string `json:"id" validate:"required"`
	Logo_url string `json:"logo_url" validate:"required"`
}

type FileRequest struct {
	ID       string `json:"id" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
}

type CoverRequest struct {
	ID       string `json:"id" validate:"required"`
	IdStore  string `json:"id_store" validate:"required"`
	FileName string `json:"file_name" validate:"required"`
}

type UpdateStoreLocationRequest struct {
	ID        string `json:"id" validate:"required"`
	Longitude string `json:"longitude" validate:"required"`
	Latitude  string `json:"latitude" validate:""`
}

type UpdateStoreStatusRequest struct {
	ID        string `json:"id" validate:"required"`
	Status string `json:"status" validate:"required"`
}

// Validate the request
func (r CreateStoreRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field %s %s", _err.Field(), _err.Tag())
	}
	return nil
}

func (r UpdateStoreRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field %s %s", _err.Field(), _err.Tag())
	}
	return nil
}

func (r ChangeStoreLogoRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field %s %s", _err.Field(), _err.Tag())
	}
	return nil
}

func (r UpdateStoreLocationRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field %s %s", _err.Field(), _err.Tag())
	}
	return nil
}

func (r UpdateStoreStatusRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		_err := err.(validator.ValidationErrors)[0]
		return fmt.Errorf("Field %s %s", _err.Field(), _err.Tag())
	}
	return nil
}
