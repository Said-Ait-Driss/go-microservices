package repository

import (
	"context"
	"errors"
	"fmt"
	"log"
	"review-service/pkg/database"
	entity "review-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StoreReviewsRepository struct {
	collection *mongo.Collection
}

func NewStoreReviewsRepository() *StoreReviewsRepository {

	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("review-service")

	return &StoreReviewsRepository{
		collection: db.Collection("store_reviews"),
	}
}

// methods

func (repo *StoreReviewsRepository) GetStoreReviews(id string) ([]entity.StoreReviews, error) {
	var store_reviews []entity.StoreReviews
	filter := bson.M{"store_id": id}

	cursor, err := repo.collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var reviews entity.StoreReviews
		if err := cursor.Decode(&reviews); err != nil {
			return nil, err
		}
		store_reviews = append(store_reviews, reviews)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return store_reviews, nil

}

func (repo *StoreReviewsRepository) CreateStoreReview(store_review entity.StoreReviews) error {
	var found_store_reviews entity.StoreReviews

	filter := bson.M{"store_id": store_review.Store_id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&found_store_reviews)
	if err != nil {
		log.Printf("err while find store : %v\n", err)
		// Store does not exist, create a new product entry with the review
		_, err := repo.collection.InsertOne(context.Background(), store_review)
		if err != nil {
			log.Printf("Failed to insert new store review: %v\n", err)
			return err
		}
		return nil

	}
	// Store exists, check if user already has a review
	var existingReview *entity.Review

	for i, review := range found_store_reviews.Reviews {
		if review.User_id == store_review.Reviews[0].User_id {
			existingReview = &found_store_reviews.Reviews[i]
			break
		}
	}

	if existingReview != nil {
		// Update existing review
		*existingReview = store_review.Reviews[0]
		filter = bson.M{"store_id": store_review.Store_id, "reviews.user_id": store_review.Reviews[0].User_id}
		update := bson.M{"$set": bson.M{"reviews": found_store_reviews.Reviews}}
		_, err = repo.collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}
	} else {
		return repo.AddReview(context.Background(), store_review.Store_id, store_review.Reviews[0])
	}
	return nil
}

// AddReview adds a new review to the product reviews collection
func (repo *StoreReviewsRepository) AddReview(ctx context.Context, store_id string, review entity.Review) error {

	filter := bson.M{"store_id": store_id}
	update := bson.M{"$push": bson.M{"reviews": review}}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to add review: %v\n", err)
		return err
	}

	return nil
}

// get reviews of a list of stores
func (repo *StoreReviewsRepository) GetReviewsOfListOfStores(ids []string) ([]entity.StoreReviews, error) {
	var store_reviews []entity.StoreReviews

	// Loop through each store ID in the request
	for _, storeID := range ids {
		var reviews entity.StoreReviews

		// Find the reviews for the current store ID
		filter := bson.M{"store_id": storeID}
		err := repo.collection.FindOne(context.Background(), filter).Decode(&reviews)
		if err == mongo.ErrNoDocuments {
			log.Printf("No reviews found for store ID: %s\n", storeID)
			continue
		} else if err != nil {
			return nil, err
		}

		// Append the reviews to the store_reviews slice
		store_reviews = append(store_reviews, reviews)
	}

	return store_reviews, nil

}

func (repo *StoreReviewsRepository) DeleteReviewByStore(store_id string, review_id string) error {

	// Filter to find the specific review within the StoreReviews document
	filter := bson.M{
		"store_id":    store_id,
		"reviews._id": review_id,
	}

	// Update operation to set the new status
	update := bson.M{
		"$set": bson.M{
			"reviews.$.status": "pending-to-delete",
		},
	}

	// Perform the update operation
	_, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Failed to update review status: %v\n", err)
		return err
	}

	return nil
}

func (repo *StoreReviewsRepository) GetDeletedReviewsByStore(storeId string) (entity.StoreReviews, error) {
	var storeReview entity.StoreReviews
	filter := bson.M{"store_id": storeId}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&storeReview)
	if err == mongo.ErrNoDocuments {
		log.Printf("No reviews found for store ID: %s\n", storeId)
	} else if err != nil {
		return entity.StoreReviews{}, err
	}

	if storeReview.Store_id == "" {
		return entity.StoreReviews{}, errors.New("store ID not found")
	}

	var deletedReviews []entity.Review
	// Iterate over the cursor
	for _, review := range storeReview.Reviews {

		if review.Status == "pending_to_delete_by_store" {
			deletedReviews = append(deletedReviews, review)
		}
	}
	return entity.StoreReviews{
		Store_id: storeId,
		Title:    storeReview.Title,
		Reviews:  deletedReviews,
	}, nil

}

// get deleted reviews by user
func (repo *StoreReviewsRepository) GetDeletedReviewsByUser(userId string) ([]entity.StoreReviews, error) {
	var deletedReviews []entity.StoreReviews

	filter := bson.M{
		"reviews.status":  "pending_to_delete_by_user",
		"reviews.user_id": userId,
	}

	// Use the aggregation framework to unwind the reviews array and filter
	cursor, err := repo.collection.Aggregate(context.Background(), mongo.Pipeline{
		{{"$match", filter}},
		{{"$unwind", "$reviews"}},
		{{"$match", bson.M{"reviews.status": "pending_to_delete_by_user", "reviews.user_id": userId}}},
		{{"$group", bson.M{
			"_id":        "$_id",
			"reviews":    bson.M{"$push": "$reviews"},
			"created_at": bson.M{"$first": "$created_at"},
			"updated_at": bson.M{"$first": "$updated_at"},
			"title":      bson.M{"$first": "$title"},
			"store_id":   bson.M{"$first": "$store_id"},
		}}},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate through the cursor and decode the results
	for cursor.Next(context.Background()) {
		var review entity.StoreReviews
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		deletedReviews = append(deletedReviews, review)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return deletedReviews, nil
}

func (repo *StoreReviewsRepository) ChangeReviewStatus(storeId string, reviewID string, newStatus string) (entity.StoreReviews, error) {

	// Define the filter to find the specific review within the StoreReviews document
	objectID, errObjectId := primitive.ObjectIDFromHex(reviewID)
	if errObjectId != nil {
		return entity.StoreReviews{}, fmt.Errorf("invalid review ID: %v", errObjectId)
	}

	filterToUpdate := bson.M{
		"store_id":    storeId,
		"reviews._id": objectID,
	}

	// Define the update operation to set the new status
	update := bson.M{
		"$set": bson.M{
			"reviews.$.status": newStatus,
		},
	}

	// Perform the update operation
	_, errToUpdate := repo.collection.UpdateOne(context.Background(), filterToUpdate, update)
	if errToUpdate != nil {
		return entity.StoreReviews{}, errToUpdate
	}

	var updatedStoreReview entity.StoreReviews
	filter := bson.M{"store_id": storeId}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&updatedStoreReview)
	if err == mongo.ErrNoDocuments {
		log.Printf("No reviews found for store ID: %s\n", storeId)
		return entity.StoreReviews{}, err
	} else if err != nil {
		return entity.StoreReviews{}, err
	}

	return updatedStoreReview, nil
}
