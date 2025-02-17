package service

import (
	repository "product-service/internal/repository/offer"
	entity "product-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *repository.OfferRepository

func init() {

	repo = repository.NewOfferRepository()
}

type OfferService interface {
	CreateOffer(offer entity.Offer) (entity.Offer, error)
	GetOffers(page int, pageSize int) ([]entity.Offer, error)
	GetOffer(id primitive.ObjectID) (entity.Offer, error)
	UpdateOffer(offer entity.Offer, id primitive.ObjectID) (entity.Offer, error)
	DeleteOffer(id primitive.ObjectID) error
}

// implementations of offer service
type OfferServiceImpl struct{}

func (service *OfferServiceImpl) GetOffer(id primitive.ObjectID) (entity.Offer, error) {

	offer, err := repo.GetOfferByID(id)

	if err != nil {
		return entity.Offer{}, err
	}

	return offer, nil
}

func (service *OfferServiceImpl) CreateOffer(offer entity.Offer) (entity.Offer, error) {

	if err := repo.CreateOffer(offer); err != nil {
		return entity.Offer{}, err
	}

	return offer, nil
}

func (service *OfferServiceImpl) UpdateOffer(offer entity.Offer, id primitive.ObjectID) (entity.Offer, error) {
	if err := repo.UpdateOffer(offer); err != nil {
		return entity.Offer{}, err
	}
	return offer, nil
}

func (service *OfferServiceImpl) GetOffers(page int, pageSize int) ([]entity.Offer, error) {
	offers, err := repo.GetOffers(page, pageSize)

	if err != nil {
		return []entity.Offer{}, err
	}

	return offers, nil
}

func (service *OfferServiceImpl) DeleteOffer(id primitive.ObjectID) error {
	if err := repo.DeleteOfferByID(id); err != nil {
		return err
	}
	return nil
}
