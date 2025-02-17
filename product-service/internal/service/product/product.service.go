package service

import (
	repository "product-service/internal/repository/product"
	entity "product-service/pkg/entities"
	request "product-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *repository.ProductRepository

func init() {
	repo = repository.NewProductRepository()
}

type ProductService interface {
	GetProduct(id primitive.ObjectID) (entity.Product, error)
	GetProducts(page int, pageSize int) ([]entity.Product, error)
	CreateProduct(product entity.Product) (entity.Product, error)
	UpdateProduct(product entity.Product, id primitive.ObjectID) (entity.Product, error)
	DeleteProduct(id primitive.ObjectID) error
	UpdateProductsStore(storeInfo request.StoreInfo) error

	// getting products count of a store
	GetProductsCount(store_id string) (int64, error)

	// getting products that has offers
	GetProductsThatHasOffers(page int, pageSize int) ([]entity.Product, error)
	// products of store that has offers
	GetProductsOfStoreThatHasOffers(store_id string, page int, pageSize int) ([]entity.Product, error)

	GetStoreProductsFilter(store_id string, name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, page int, pageSize int) ([]entity.Product, error)
	GetProductsFilter(name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, neighborhood string, storeTitle string, page int, pageSize int) ([]entity.Product, error)
	GetProductsByCategory(category_id string) ([]entity.Product, error)

	GetProductsByStore(storeId string) ([]entity.Product, error)
	// for consumed messages
	SubtractProductQty(product_id primitive.ObjectID, qty int) error
	AddProductQty(product_id primitive.ObjectID, qty int) error
}

// implementations of product service
type ProductServiceImpl struct{}

func (service *ProductServiceImpl) GetProduct(id primitive.ObjectID) (entity.Product, error) {

	product, err := repo.GetProductByID(id)

	if err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (service *ProductServiceImpl) GetProducts(page int, pageSize int) ([]entity.Product, error) {
	products, err := repo.GetProducts(page, pageSize)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

func (service *ProductServiceImpl) CreateProduct(product entity.Product) (entity.Product, error) {

	if err := repo.CreateProduct(product); err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (service *ProductServiceImpl) UpdateProduct(product entity.Product, id primitive.ObjectID) (entity.Product, error) {

	if err := repo.UpdateProduct(product); err != nil {
		return entity.Product{}, err
	}
	return product, nil

}

func (service *ProductServiceImpl) DeleteProduct(id primitive.ObjectID) error {
	if err := repo.DeleteProductByID(id); err != nil {
		return err
	}
	return nil

}

func (service *ProductServiceImpl) UpdateProductsStore(storeInfo request.StoreInfo) error {
	if err := repo.UpdateProductsStore(storeInfo.ID, storeInfo.Name); err != nil {
		return err
	}
	return nil
}

// for consumed messages
func (service *ProductServiceImpl) SubtractProductQty(product_id primitive.ObjectID, qty int) error {
	if err := repo.UpdateProductQty(product_id, qty, "-"); err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) AddProductQty(product_id primitive.ObjectID, qty int) error {
	if err := repo.UpdateProductQty(product_id, qty, "+"); err != nil {
		return err
	}
	return nil
}

func (service *ProductServiceImpl) GetProductsByStore(storeId string) ([]entity.Product, error) {
	products, err := repo.GetProductsByStore(storeId)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

// getting products count of a store
func (service *ProductServiceImpl) GetProductsCount(storeId string) (int64, error) {
	count, err := repo.GetProductsCount(storeId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// getting products that has offers
func (service *ProductServiceImpl) GetProductsThatHasOffers(page int, pageSize int) ([]entity.Product, error) {
	products, err := repo.GetProductsThatHasOffers(page, pageSize)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

// get products of store that has offers
func (service *ProductServiceImpl) GetProductsOfStoreThatHasOffers(store_id string, page int, pageSize int) ([]entity.Product, error) {
	products, err := repo.GetProductsOfStoreThatHasOffers(store_id, page, pageSize)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

// get filtered products
func (service *ProductServiceImpl) GetStoreProductsFilter(store_id string, name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, page int, pageSize int) ([]entity.Product, error) {
	products, err := repo.GetStoreProductsFiltered(store_id, name, minPrice, maxPrice, startDate, endDate, page, pageSize)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

func (service *ProductServiceImpl) GetProductsFilter(name string, minPrice float64, maxPrice float64, startDate time.Time, endDate time.Time, neighborhood string, storeTitle string, page int, pageSize int) ([]entity.Product, error) {
	products, err := repo.GetProductsFilter(name, minPrice, maxPrice, startDate, endDate, neighborhood, storeTitle, page, pageSize)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}

func (service *ProductServiceImpl) GetProductsByCategory(category_id string) ([]entity.Product, error) {
	products, err := repo.GetProductsByCategory(category_id, 0, 6)

	if err != nil {
		return []entity.Product{}, err
	}

	return products, nil
}
