package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetProductReviewsEndpoint   endpoint.Endpoint
	CreateProductReviewEndpoint endpoint.Endpoint

	GetStoreReviewsEndpoint   endpoint.Endpoint
	CreateStoreReviewEndpoint endpoint.Endpoint

	GetReviewsOfListOfStoresEndpoint endpoint.Endpoint

	GetDeletedReviewsByStoreEndpoint endpoint.Endpoint

	GetDeletedReviewsByUserEndpoint endpoint.Endpoint

	ChangeReviewStatusEndpoint endpoint.Endpoint
}
