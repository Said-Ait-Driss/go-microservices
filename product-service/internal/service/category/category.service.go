package service

import (
	repository "product-service/internal/repository/category"
	entity "product-service/pkg/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var repo *repository.CategoryRepository

func init() {

	repo = repository.NewCategoryRepository()
}

type CategoryService interface {
	GetCategory(id primitive.ObjectID) (entity.Category, error)
	CreateCategory(category entity.Category) (entity.Category, error)
	UpdateCategory(category entity.Category, category_id primitive.ObjectID) (entity.Category, error)
	DeleteCategory(category_id primitive.ObjectID) error
	GetCategories(page int, pageSize int) ([]entity.Category, error)
}

// implementations of category service
type CategoryServiceImpl struct{}

func (service *CategoryServiceImpl) GetCategory(id primitive.ObjectID) (entity.Category, error) {

	category, err := repo.GetCategoryByID(id)

	if err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (service *CategoryServiceImpl) GetCategories(page int, pageSize int) ([]entity.Category, error) {
	categories, err := repo.GetCategories(page, pageSize)

	if err != nil {
		return []entity.Category{}, err
	}

	return categories, nil
}

func (service *CategoryServiceImpl) CreateCategory(category entity.Category) (entity.Category, error) {

	if err := repo.CreateCategory(category); err != nil {
		return entity.Category{}, err
	}

	return category, nil
}

func (service *CategoryServiceImpl) UpdateCategory(category entity.Category, category_id primitive.ObjectID) (entity.Category, error) {
	if err := repo.UpdateCategory(category); err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (service *CategoryServiceImpl) DeleteCategory(category_id primitive.ObjectID) error {
	if err := repo.DeleteCategory(category_id); err != nil {
		return err
	}
	return nil
}
