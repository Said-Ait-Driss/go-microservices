package repository

import (
	"context"
	"product-service/pkg/database"
	entity "product-service/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository struct {
	collection *mongo.Collection
}

func NewCategoryRepository() *CategoryRepository {

	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("product-service")

	return &CategoryRepository{
		collection: db.Collection("categories"),
	}
}

// methods

func (repo *CategoryRepository) CreateCategory(category entity.Category) error {
	category.Created_at = time.Now()
	category.Updated_at = time.Now()

	_, err := repo.collection.InsertOne(context.Background(), category)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) GetCategories(page int, pageSize int) ([]entity.Category, error) {
	var categories []entity.Category
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

	// Decode each document from the cursor and add it to the categories

	for cursor.Next(context.Background()) {
		var category entity.Category
		if err := cursor.Decode(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *CategoryRepository) GetCategoryByID(id primitive.ObjectID) (entity.Category, error) {
	var category entity.Category

	filter := bson.M{"_id": id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&category)
	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (repo *CategoryRepository) UpdateCategory(category entity.Category) error {
	category.Updated_at = time.Now()

	filter := bson.M{"_id": category.ID}
	update := bson.M{"$set": category}

	_, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) DeleteCategory(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}
