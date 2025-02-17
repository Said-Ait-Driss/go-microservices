package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetProductEndpoint        endpoint.Endpoint
	GetProductsEndpoit        endpoint.Endpoint
	CreateProductEndpoint     endpoint.Endpoint
	UpdateProductEndpoint     endpoint.Endpoint
	DeleteProductEndpoint     endpoint.Endpoint
	GetProductsByStoreEndpoit endpoint.Endpoint
	// offers
	CreateOfferEndpoint endpoint.Endpoint
	GetOffersEndpoit    endpoint.Endpoint
	GetOfferEndpoint    endpoint.Endpoint
	UpdateOfferEndpoint endpoint.Endpoint
	DeleteOfferEndpoint endpoint.Endpoint

	// category
	CreateCategoryEndpoint endpoint.Endpoint
	GetCategoryEndpoint    endpoint.Endpoint
	UpdateCategoryEndpoint endpoint.Endpoint
	DeleteCategoryEndpoint endpoint.Endpoint
	GetCategoriesEndpoit   endpoint.Endpoint

	// get products count
	GetProductsCountByStoreEndpoint endpoint.Endpoint

	// get products that has offers
	GetProductsThatHasOffersEndpoint endpoint.Endpoint

	// get products of store that has offers
	GetProductsOfStoreThatHasOffersEndpoint endpoint.Endpoint

	// get store products filter
	GetStoreProductsFilterEndpoint endpoint.Endpoint

	GetProductsFilterEndpoint endpoint.Endpoint

	GetProductsByCategoryEndpoint endpoint.Endpoint

	GetProductsByStoreEndpoint endpoint.Endpoint
}
