package response

import (
	entity "review-service/pkg/entities"
)

type GetProductReviewsResponse struct {
	ProductReviews []entity.ProductReviews `json:"product_reviews"`
	Err            error
}

type CreateNewProductReviewResponse struct {
	ProductReviews []entity.ProductReviews `json:"product_reviews"`
	Err            error
}
