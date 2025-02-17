package repository

import (
	"context"
	"log"
	"review-service/pkg/database"
	entity "review-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductReviewsRepository struct {
	collection *mongo.Collection
}

func NewProductReviewsRepository() *ProductReviewsRepository {

	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("review-service")

	return &ProductReviewsRepository{
		collection: db.Collection("product_reviews"),
	}
}

// methods

func (repo *ProductReviewsRepository) GetProductReviews(id primitive.ObjectID) ([]entity.ProductReviews, error) {
	var product_reviews []entity.ProductReviews

	filter := bson.M{"product_id": id}

	cursor, err := repo.collection.Find(context.Background(), filter)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var reviews entity.ProductReviews
		if err := cursor.Decode(&reviews); err != nil {
			return nil, err
		}
		product_reviews = append(product_reviews, reviews)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return product_reviews, nil

}

func (repo *ProductReviewsRepository) CreateProductReview(product_review entity.ProductReviews) error {
	var found_product_reviews entity.ProductReviews

	filter := bson.M{"product_id": product_review.Product_id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&found_product_reviews)
	if err != nil {
		log.Printf("err while find product : %v\n", err)
		// Product does not exist, create a new product entry with the review
		_, err := repo.collection.InsertOne(context.Background(), product_review)
		if err != nil {
			log.Printf("Failed to insert new product review: %v\n", err)
			return err
		}
		return nil

	}
	// Product exists, check if user already has a review
	var existingReview *entity.Review

	for i, review := range found_product_reviews.Reviews {
		if review.User_id == product_review.Reviews[0].User_id {
			existingReview = &found_product_reviews.Reviews[i]
			break
		}
	}

	if existingReview != nil {
		// Update existing review
		*existingReview = product_review.Reviews[0]
		filter = bson.M{"product_id": product_review.Product_id, "reviews.user_id": product_review.Reviews[0].User_id}
		update := bson.M{"$set": bson.M{"reviews": found_product_reviews.Reviews}}
		_, err = repo.collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return err
		}
	} else {
		return repo.AddReview(context.Background(), product_review.Product_id, product_review.Reviews[0])
	}
	return nil
}

// AddReview adds a new review to the product reviews collection
func (repo *ProductReviewsRepository) AddReview(ctx context.Context, productID primitive.ObjectID, review entity.Review) error {

	filter := bson.M{"product_id": productID}
	update := bson.M{"$push": bson.M{"reviews": review}}

	_, err := repo.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Failed to add review: %v\n", err)
		return err
	}

	return nil
}
