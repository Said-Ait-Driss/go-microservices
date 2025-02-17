package service

import (
	repository "review-service/internal/repository/store_review"
	entity "review-service/pkg/entities"
)

var repo *repository.StoreReviewsRepository

func init() {
	repo = repository.NewStoreReviewsRepository()
}

type StoreReviewsService interface {
	GetStoreReviews(id string) ([]entity.StoreReviews, error)
	CreateStoreReview(store_review entity.StoreReviews) ([]entity.StoreReviews, error)
	GetReviewsOfListOfStores(ids []string) ([]entity.StoreReviews, error)

	GetDeletedReviewsByStore(id string) (entity.StoreReviews, error)

	GetDeletedReviewsByUser(user_id string) ([]entity.StoreReviews, error)

	ChangeReviewStatus(store_id string, review_id string, new_Status string) (entity.StoreReviews, error)

	// DeleteReviewByStore(req request.DeleteReviewByStoreRequest) (response.DeleteReviewByStoreResponse, error)
}

// implementations of product service
type StoreReviewsServiceImpl struct{}

func (service *StoreReviewsServiceImpl) GetStoreReviews(id string) ([]entity.StoreReviews, error) {

	reviews, err := repo.GetStoreReviews(id)

	if err != nil {
		return []entity.StoreReviews{}, err
	}

	return reviews, nil
}

func (service *StoreReviewsServiceImpl) CreateStoreReview(store_review entity.StoreReviews) ([]entity.StoreReviews, error) {

	store_id := store_review.Store_id

	err := repo.CreateStoreReview(store_review)

	if err != nil {
		return []entity.StoreReviews{}, err
	}

	return service.GetStoreReviews(store_id)
}

func (service *StoreReviewsServiceImpl) GetReviewsOfListOfStores(ids []string) ([]entity.StoreReviews, error) {

	reviews, err := repo.GetReviewsOfListOfStores(ids)

	if err != nil {
		return []entity.StoreReviews{}, err
	}

	return reviews, nil
}

func (service *StoreReviewsServiceImpl) GetDeletedReviewsByStore(store_id string) (entity.StoreReviews, error) {

	deletedReviewsOfAStore, err := repo.GetDeletedReviewsByStore(store_id)

	if err != nil {
		return entity.StoreReviews{}, err
	}

	return deletedReviewsOfAStore, nil
}

// get deleted reviews by user
func (service *StoreReviewsServiceImpl) GetDeletedReviewsByUser(user_id string) ([]entity.StoreReviews, error) {
	reviews, err := repo.GetDeletedReviewsByUser(user_id)
	if err != nil {
		return []entity.StoreReviews{}, err
	}
	return reviews, nil
}

func (service *StoreReviewsServiceImpl) ChangeReviewStatus(storeId string, reviewId string, newStatus string) (entity.StoreReviews, error) {

	storeReview, err := repo.ChangeReviewStatus(storeId, reviewId, newStatus)
	if err != nil {
		return entity.StoreReviews{}, err
	}

	return storeReview, nil
}
