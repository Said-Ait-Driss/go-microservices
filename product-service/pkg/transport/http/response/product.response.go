package response

import entity "product-service/pkg/entities"

type GetProductResponse struct {
	Product entity.Product `json:"product"`
	Err     error          `json:"error,omitempty"`
}

type GetProductsResponse struct {
	Products []entity.Product `json:"products"`
	Err      error            `json:"error,omitempty"`
}

type CreateProductResponse struct {
	Product entity.Product `json:"product"`
	Err     error          `json:"error,omitempty"`
}

type UpdateProductResponse struct {
	Product entity.Product `json:"product"`
	Err     error          `json:"error,omitempty"`
}

type DeletProductResponse struct {
	Product entity.Product `json:"product"`
	Err     error          `json:"error,omitempty"`
}

// getting products count of a store
type GetProductsCountResponse struct {
	Count int64 `json:"count"`
	Err   error `json:"error,omitempty"`
}
