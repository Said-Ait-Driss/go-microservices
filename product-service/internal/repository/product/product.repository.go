package repository

import (
	"context"
	"product-service/pkg/database"
	entity "product-service/pkg/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository() *ProductRepository {

	// Get the MongoDB client instance from the database package
	client := database.Client

	// Access the specific database from the client
	db := client.Database("product-service")

	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

// methods

func (repo *ProductRepository) CreateProduct(product entity.Product) error {
	product.Created_at = time.Now()
	product.Updated_at = time.Now()

	_, err := repo.collection.InsertOne(context.Background(), product)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) GetProductByID(id primitive.ObjectID) (entity.Product, error) {
	var product entity.Product

	filter := bson.M{"_id": id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (repo *ProductRepository) UpdateProduct(product entity.Product) error {
	product.Updated_at = time.Now()

	filter := bson.M{"_id": product.ID}
	update := bson.M{"$set": product}

	_, err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (repo *ProductRepository) DeleteProductByID(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := repo.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return nil
}

// paginate

func (repo *ProductRepository) GetProducts(page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
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

	// Decode each document from the cursor and add it to the products

	for cursor.Next(context.Background()) {
		var product entity.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil

}

func (repo *ProductRepository) UpdateProductsStore(storeID string, storeName string) error {
	filter := bson.M{"store.store_id": storeID}
	update := bson.M{
		"$set": bson.M{"store.title": storeName},
	}

	_, err := repo.collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

// for consume nats messages

func (repo *ProductRepository) UpdateProductQty(product_id primitive.ObjectID, qty int, operation string) error {

	var product entity.Product

	filter := bson.M{"_id": product_id}

	err := repo.collection.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		return err
	}
	product.Updated_at = time.Now()
	if operation == "-" {
		product.Quantity -= qty
	} else {
		product.Quantity += qty
	}
	update := bson.M{"$set": product}

	_, _err := repo.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return _err
	}

	return nil
}

func (repo *ProductRepository) GetProductsByStore(storeId string) ([]entity.Product, error) {
	var products []entity.Product
	filterById := bson.M{"store.store_id": storeId}

	cursor, err := repo.collection.Find(context.Background(), filterById)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product entity.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil

}

// getting product count of a store
func (repo *ProductRepository) GetProductsCount(storeId string) (int64, error) {
	filter := bson.M{"store.store_id": storeId}
	count, err := repo.collection.CountDocuments(context.Background(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// getting products that has offers
func (repo *ProductRepository) GetProductsThatHasOffers(page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
	// Calculate the number of documents to skip based on the page number and page size
	skip := (page - 1) * pageSize

	// Define options for pagination

	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	// Define the filter (if any)
	filter := bson.M{"current_offer.offer_id": bson.M{"$ne": ""}}

	cursor, err := repo.collection.Find(context.Background(), filter, options)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	// Decode each document from the cursor and add it to the products

	for cursor.Next(context.Background()) {
		var product entity.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// getting products that has offers
func (repo *ProductRepository) GetProductsOfStoreThatHasOffers(storeId string, page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
	// Calculate the number of documents to skip based on the page number and page size
	skip := (page - 1) * pageSize

	// Define options for pagination

	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	// Define the filter (if any)
	filter := bson.M{"store.store_id": storeId, "current_offer.offer_id": bson.M{"$ne": ""}}

	cursor, err := repo.collection.Find(context.Background(), filter, options)

	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	// Decode each document from the cursor and add it to the products

	for cursor.Next(context.Background()) {
		var product entity.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

// Get Filtered Products
func (repo *ProductRepository) GetStoreProductsFiltered(storeId string, name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
	filter := bson.M{"store.store_id": storeId}

	skip := (page - 1) * pageSize

	// Define options for pagination

	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"} // Case-insensitive search
	}

	if minPrice >= 0 && maxPrice >= 0 {
		filter["original_price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	}

	// Add date range filter
	if !startDate.IsZero() && !endDate.IsZero() {
		filter["created_at"] = bson.M{"$gte": startDate, "$lte": endDate}
	}

	// Find products based on the filter
	cursor, err := repo.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}

	return products, nil

}

func (repo *ProductRepository) GetProductsFilter(name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, neighborhood string, storeTitle string, page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
	filter := bson.M{"store.title": storeTitle}

	skip := (page - 1) * pageSize

	// Define options for pagination

	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"} // Case-insensitive search
	}

	if minPrice >= 0 && maxPrice >= 0 {
		filter["original_price"] = bson.M{"$gte": minPrice, "$lte": maxPrice}
	}

	// Add date range filter
	if !startDate.IsZero() && !endDate.IsZero() {
		filter["created_at"] = bson.M{"$gte": startDate, "$lte": endDate}
	}

	// Find products based on the filter
	cursor, err := repo.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}

	return products, nil

}

func (repo *ProductRepository) GetProductsByCategory(category_id string, page int, pageSize int) ([]entity.Product, error) {
	var products []entity.Product
	filter := bson.M{"category_id": category_id}
	skip := (page - 1) * pageSize
	options := options.Find()
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))
	cursor, err := repo.collection.Find(context.Background(), filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err := cursor.All(context.Background(), &products); err != nil {
		return nil, err
	}
	return products, nil

}
