package response

import (
	entity "review-service/pkg/entities"
)

type GetStoreReviewsResponse struct {
	StoreReviews []entity.StoreReviews `json:"store_reviews"`
	Err          error
}

type GetSingleStoreReviewsResponse struct {
	StoreReviews entity.StoreReviews `json:"store_reviews"`
	Err          error
}

type GetDeletedReviewsByStoreResponse struct {
	StoreReview entity.StoreReviews `json:"store_reviews"`
	Err         error
}

type GetDeletedReviewsByUserResponse struct {
	StoreReview []entity.StoreReviews `json:"store_reviews"`
	Err         error
}

type CreateNewStoreReviewResponse struct {
	StoreReviews []entity.StoreReviews `json:"store_reviews"`
	Err          error
}

type DeleteReviewByStoreResponse struct {
	Statue_changing bool `json:"statue_changing"`
	Err             error
}
