package endpoint

import (
	"context"
	service "review-service/internal/service/store_review"
	entity "review-service/pkg/entities"
	"review-service/pkg/mapper"
	request "review-service/pkg/transport/http/request"
	response "review-service/pkg/transport/http/response"

	"github.com/go-kit/kit/endpoint"
)

func MakeStoreReviewsEndpoints(s service.StoreReviewsService) Endpoints {
	return Endpoints{
		GetStoreReviewsEndpoint:          MakeGetStoreReviewsEndpoint(s),
		CreateStoreReviewEndpoint:        MakeCreateStoreReviewEndpoint(s),
		GetDeletedReviewsByStoreEndpoint: MakeGetDeletedReviewsByStoreEndpoint(s),

		GetDeletedReviewsByUserEndpoint: MakeGetDeletedReviewsByUserEndpoint(s),

		ChangeReviewStatusEndpoint: MakeChangeReviewStatusEndpoint(s),

		GetReviewsOfListOfStoresEndpoint: MakeGetReviewsOfListOfStoresEndpoint(s),
	}
}

func MakeGetStoreReviewsEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {

		req := _request.(request.GetStoreReviewsRequest)

		reviews, err := srv.GetStoreReviews(req.ID)
		return response.GetStoreReviewsResponse{StoreReviews: reviews, Err: err}, nil
	}
}

func MakeGetReviewsOfListOfStoresEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetReviewsOfListOfStoresRequest)
		reviews, err := srv.GetReviewsOfListOfStores(req.Store_ids)
		return response.GetStoreReviewsResponse{StoreReviews: reviews, Err: err}, nil
	}
}

func MakeCreateStoreReviewEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.CreateNewStoreReviewRequest)

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.CreateNewStoreReviewResponse{StoreReviews: []entity.StoreReviews{}, Err: err}, nil
		}

		// Map req to entity.Store
		store_review := mapper.MapCreateStoreReviewRequest(req)
		reviews, err := srv.CreateStoreReview(store_review)

		return response.CreateNewStoreReviewResponse{StoreReviews: reviews, Err: err}, nil
	}
}

// get deleted reviews by store
func MakeGetDeletedReviewsByStoreEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.DeleteReviewByStoreRequest)

		reviews, err := srv.GetDeletedReviewsByStore(req.Store_id)

		return response.GetDeletedReviewsByStoreResponse{StoreReview: reviews, Err: err}, nil
	}
}

// get deleted reviews by user
func MakeGetDeletedReviewsByUserEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.DeleteReviewByUserRequest)
		reviews, err := srv.GetDeletedReviewsByUser(req.User_id)
		return response.GetDeletedReviewsByUserResponse{StoreReview: reviews, Err: err}, nil
	}
}

// change review status
func MakeChangeReviewStatusEndpoint(srv service.StoreReviewsService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.ChangeReviewStatusRequest)

		storeReviews, err := srv.ChangeReviewStatus(req.Store_id, req.Review_id, req.New_Status)

		return response.GetSingleStoreReviewsResponse{StoreReviews: storeReviews, Err: err}, nil
	}
}
