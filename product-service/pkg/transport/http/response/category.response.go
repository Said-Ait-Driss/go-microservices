package response

import entity "product-service/pkg/entities"

type GetCategoryResponse struct {
	Category entity.Category `json:"category"`
	Err      error           `json:"error,omitempty"`
}

type CreateCategoryResponse struct {
	Category entity.Category `json:"category"`
	Err      error           `json:"error,omitempty"`
}

type UpdateCategoryResponse struct {
	Category entity.Category `json:"category"`
	Err      error           `json:"error,omitempty"`
}

type DeletCategoryResponse struct {
	Category entity.Category `json:"category"`
	Err      error           `json:"error,omitempty"`
}

type GetCategoriesResponse struct {
	Categories []entity.Category `json:"categories"`
	Err        error             `json:"error,omitempty"`
}
