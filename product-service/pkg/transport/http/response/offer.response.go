package response

import entity "product-service/pkg/entities"

type GetOfferResponse struct {
	Offer entity.Offer `json:"offer"`
	Err   error        `json:"error,omitempty"`
}

type CreateOfferResponse struct {
	Offer entity.Offer `json:"offer"`
	Err   error        `json:"error,omitempty"`
}

type UpdateOfferResponse struct {
	Offer entity.Offer `json:"offer"`
	Err   error        `json:"error,omitempty"`
}

type GetOffersResponse struct {
	Offers []entity.Offer `json:"offers"`
	Err    error          `json:"error,omitempty"`
}

type DeletOfferResponse struct {
	Offer entity.Offer `json:"offer"`
	Err   error        `json:"error,omitempty"`
}
