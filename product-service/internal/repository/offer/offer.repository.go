package repository

import (
	"context"
	"product-service/pkg/database"
	entity "product-service/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type OfferRepository struct {
	collection *mongo.Collection
}

func NewOfferRepository() *OfferRepository {

	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("product-service")

	return &OfferRepository{
		collection: db.Collection("offers"),
	}
}

// methods

func (repo *OfferRepository) CreateOffer(offer entity.Offer) error {
	offer.Created_at = time.Now()
	offer.Updated_at = time.Now()

	_, err := repo.collection.InsertOne(context.Background(), offer)
	if err != nil {
		return err
	}

	return nil
}

func (repo *OfferRepository) GetOfferByID(id primitive.ObjectID) (entity.Offer, error) {
	var offer entity.Offer

	filter := bson.M{"_id": id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&offer)
	if err != nil {
		return entity.Offer{}, err
	}

	return offer, nil
}

func (repo *OfferRepository) UpdateOffer(offer entity.Offer) error {
	offer.Updated_at = time.Now()

	filter := bson.M{"_id": offer.ID}
	update := bson.M{"$set": offer}

	_, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *OfferRepository) DeleteOfferByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// paginate

func (repo *OfferRepository) GetOffers(page int, pageSize int) ([]entity.Offer, error) {
	var offers []entity.Offer
	// Calculate the number of documents to skip based on the page number and page size
	skip := (page - 1) * pageSize

	// Define options for pagination

	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	// Define the filter (if any)
	filter := bson.M{}

	cursor, err := repo.collection.Find(context.Background(), filter, options)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	// Decode each document from the cursor and add it to the offers

	for cursor.Next(context.Background()) {
		var offer entity.Offer
		if err := cursor.Decode(&offer); err != nil {
			return nil, err
		}
		offers = append(offers, offer)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return offers, nil

}
