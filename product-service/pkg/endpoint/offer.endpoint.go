package endpoint

import (
	"context"
	service "product-service/internal/service/offer"
	entity "product-service/pkg/entities"
	"product-service/pkg/mapper"
	request "product-service/pkg/transport/http/request"
	response "product-service/pkg/transport/http/response"

	"github.com/go-kit/kit/endpoint"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MakeOfferEndpoints(s service.OfferService) Endpoints {
	return Endpoints{
		CreateOfferEndpoint: MakeCreateOfferEndpoint(s),
		GetOffersEndpoit:    MakeGetOffersEndpoint(s),
		GetOfferEndpoint:    MakeGetOfferEndpoit(s),
		UpdateOfferEndpoint: MakeUpdateOfferEndpoint(s),
		DeleteOfferEndpoint: MakeDeleteOfferEndpoint(s),
	}
}

func MakeCreateOfferEndpoint(srv service.OfferService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.CreateOfferRequest)

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.CreateOfferResponse{Offer: entity.Offer{}, Err: err}, nil
		}

		// Map req to entity.Offer
		offer := mapper.MapCreateOfferRequest(req)
		_offer, err := srv.CreateOffer(offer)

		return response.CreateOfferResponse{Offer: _offer, Err: err}, nil
	}
}

func MakeGetOffersEndpoint(srv service.OfferService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetOffersRequest)

		offers, err := srv.GetOffers(req.Page, req.PageSize)
		return response.GetOffersResponse{Offers: offers, Err: err}, nil
	}
}

func MakeGetOfferEndpoit(srv service.OfferService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.GetOfferRequest)

		// Parse the string ID into primitive.ObjectID
		offer_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.GetOfferResponse{Offer: entity.Offer{}, Err: err}, err
		}

		offer, err := srv.GetOffer(offer_id)
		return response.GetOfferResponse{Offer: offer, Err: err}, nil
	}
}

func MakeUpdateOfferEndpoint(srv service.OfferService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.UpdateOfferRequest)

		// Parse the string ID into primitive.ObjectID
		offer_id, err := primitive.ObjectIDFromHex(primitive.ObjectID.Hex(req.ID))
		if err != nil {
			return response.UpdateOfferResponse{Offer: entity.Offer{}, Err: err}, err
		}

		// Validate the request
		if err := req.Validate(); err != nil {
			return response.UpdateOfferResponse{Offer: entity.Offer{}, Err: err}, nil
		}

		// Map req to entity.Offer
		offer := mapper.MapUpdateOfferRequest(req)

		_offer, err := srv.UpdateOffer(offer, offer_id)

		return response.UpdateOfferResponse{Offer: _offer, Err: err}, nil
	}
}

func MakeDeleteOfferEndpoint(srv service.OfferService) endpoint.Endpoint {
	return func(ctx context.Context, _request interface{}) (interface{}, error) {
		req := _request.(request.DeleteOfferRequest)

		// Parse the string ID into primitive.ObjectID
		offer_id, err := primitive.ObjectIDFromHex(req.ID)
		if err != nil {
			return response.DeletOfferResponse{Offer: entity.Offer{}, Err: err}, err
		}

		err = srv.DeleteOffer(offer_id)
		return response.DeletOfferResponse{Offer: entity.Offer{ID: offer_id}, Err: err}, nil
	}
}
