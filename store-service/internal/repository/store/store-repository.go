package storeRepository

import (
	"context"
	"errors"
	"fmt"
	"store-service/pkg/database"
	entity "store-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreRepository struct {
	storeCollection    *mongo.Collection
	categoryCollection *mongo.Collection
}

func NewStoreRepository() *StoreRepository {
	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("store_db")

	return &StoreRepository{
		storeCollection:    db.Collection("stores"),
		categoryCollection: db.Collection("categories"),
	}
}

func (or *StoreRepository) GetStoreByID(storeID string) (entity.Store, error) {
	objectID, error := primitive.ObjectIDFromHex(storeID)
	if error != nil {
		return entity.Store{}, error
	}

	filter := bson.M{"_id": objectID}
	var store entity.Store
	err := or.storeCollection.FindOne(context.Background(), filter).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return entity.Store{}, errors.New("Store not found")
	}
	return store, err
}

func (or *StoreRepository) GetAllStores() ([]entity.Store, error) {
	var stores []entity.Store

	cursor, err := or.categoryCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	fmt.Printf("**************** STORES :", len(stores))
	for cursor.Next(context.Background()) {
		var store entity.Store
		if err := cursor.Decode(&store); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return stores, nil
}

func (or *StoreRepository) CreateStore(store entity.Store) error {
	_, err := or.storeCollection.InsertOne(context.Background(), store)
	return err
}

func (or *StoreRepository) UpdateStore(updatedStore entity.Store) error {
	filter := bson.M{"_id": updatedStore.ID}
	update := bson.M{"$set": updatedStore}

	_, err := or.storeCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (or *StoreRepository) ChangeStoreLogo(storeID string, fileName string) error {
	objectID, error := primitive.ObjectIDFromHex(storeID)
	if error != nil {
		return error
	}
	
	filterById := bson.M{"_id": objectID}
	var store entity.Store
	err := or.storeCollection.FindOne(context.Background(), filterById).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return errors.New("Store not found")
	}

	store.Logo_url = fileName
	filter := bson.M{"_id": store.ID}
	update := bson.M{"$set": store}
	_, err = or.storeCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (or *StoreRepository) ChangeStoreCover(coverID string,storeID string, fileName string) error {
	objectID, error := primitive.ObjectIDFromHex(storeID)
	if error != nil {
		return error
	}
	fmt.Println(objectID)
	filterById := bson.M{"_id": objectID}
	var store entity.Store
	err := or.storeCollection.FindOne(context.Background(), filterById).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return errors.New("Store not found")
	}

	// find cover by coverId
	coverObjectID, error := primitive.ObjectIDFromHex(coverID)
	if error != nil {
		return error
	}
	coverFound := false
	for i := range store.Covers {
		if store.Covers[i].ID == coverObjectID {
			store.Covers[i].Url = fileName
			coverFound = true
			break
		}
	}

	if !coverFound {
		return errors.New("Cover not found")
	}

	filter := bson.M{"_id": store.ID}
	update := bson.M{"$set": store}
	_, err = or.storeCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (or *StoreRepository) UpdateStoreLocation(updatedStore entity.Store) error {
	filterById := bson.M{"_id": updatedStore.ID}
	var store entity.Store
	err := or.storeCollection.FindOne(context.Background(), filterById).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return errors.New("Store not found")
	}

	store.Longitude = store.Longitude
	store.Latitude = store.Latitude
	filter := bson.M{"_id": store.ID}
	update := bson.M{"$set": store}
	_, err = or.storeCollection.UpdateOne(context.Background(), filter, update)
	return err
}

func (or *StoreRepository) UpdateStoreStatus(updatedStore entity.Store) error {
	filterById := bson.M{"_id": updatedStore.ID}
	var store entity.Store
	err := or.storeCollection.FindOne(context.Background(), filterById).Decode(&store)
	if err == mongo.ErrNoDocuments {
		return errors.New("Store not found")
	}

	store.Status = store.Status
	filter := bson.M{"_id": store.ID}
	update := bson.M{"$set": store}
	_, err = or.storeCollection.UpdateOne(context.Background(), filter, update)
	return err
}

// start function sof category

func (or *StoreRepository) GetAllCategories() ([]entity.Categories, error) {
	var categories []entity.Categories

	cursor, err := or.categoryCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var categorie entity.Categories
		if err := cursor.Decode(&categorie); err != nil {
			return nil, err
		}
		categories = append(categories, categorie)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (or *StoreRepository) CreateCategory(category entity.Categories) error {
	_, err := or.categoryCollection.InsertOne(context.Background(), category)
	return err
}

func (or *StoreRepository) UpdateCategory(updatedCategory entity.Categories) error {
	filter := bson.M{"_id": updatedCategory.ID}
	update := bson.M{"$set": updatedCategory}

	_, err := or.categoryCollection.UpdateOne(context.Background(), filter, update)
	return err
}
