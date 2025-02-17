package storeService

import (
	storeRepository "store-service/internal/repository/store"
	mapper "store-service/pkg/mapper"
	"store-service/pkg/transport/http/request"
	"store-service/pkg/transport/http/response"
	producer "store-service/pkg/transport/nats/producer"
)

type StoreService struct {
	storeRepo *storeRepository.StoreRepository
}

func NewStoreService() *StoreService {
	return &StoreService{
		storeRepo: storeRepository.NewStoreRepository(),
	}
}

func (us *StoreService) GetStoreByID(req request.GetStoreRequest) (response.DataResponse, error) {
	store, err := us.storeRepo.GetStoreByID(req.Store_id)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    store,
		Message: "Store retrieved successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) GetAllStores() (response.DataResponse, error) {
	stores, err := us.storeRepo.GetAllStores()
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    stores,
		Message: "",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) CreateStore(req request.CreateStoreRequest) (response.DataResponse, error) {
	newStore := mapper.MapCreateStoreRequest(req)

	err := us.storeRepo.CreateStore(newStore)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    newStore,
		Message: "Store created successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) UpdateStore(req request.UpdateStoreRequest) (response.DataResponse, error) {
	store := mapper.MapUpdateStoreRequest(req)

	err := us.storeRepo.UpdateStore(store)
	if err != nil {
		return response.DataResponse{}, err
	}
	storeInfo := request.StoreInfo{
		ID:   store.ID.Hex(),
		Name: store.Name,
	}

	erro := producer.UpdateProductsStore(storeInfo)
	if erro != nil {
		return response.DataResponse{}, erro
	}

	response := response.DataResponse{
		Data:    true,
		Message: "Store updated successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) ChangeStoreLogo(req request.FileRequest) error {
	err := us.storeRepo.ChangeStoreLogo(req.ID, req.FileName)
	return err
}

func (us *StoreService) ChangeStoreCover(req request.CoverRequest) error {
	err := us.storeRepo.ChangeStoreCover(req.ID, req.IdStore, req.FileName)
	return err
}

// start functions of category

func (us *StoreService) GetAllCategories() (response.DataResponse, error) {
	categories, err := us.storeRepo.GetAllCategories()
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    categories,
		Message: "Store retrieved successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) CreateCategory(req request.CreateCategoryRequest) (response.DataResponse, error) {
	newCategory := mapper.MapCreateCategoryRequest(req)

	err := us.storeRepo.CreateCategory(newCategory)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    newCategory,
		Message: "Category created successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) UpdateCategory(req request.UpdateCategoryRequest) (response.DataResponse, error) {
	Category := mapper.MapUpdateCategoryRequest(req)

	err := us.storeRepo.UpdateCategory(Category)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    true,
		Message: "Category updated successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) UpdateStoreLocation(req request.UpdateStoreLocationRequest) (response.DataResponse, error) {
	store := mapper.MapUpdateStoreLocationRequest(req)

	err := us.storeRepo.UpdateStoreLocation(store)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    true,
		Message: "Store location updated successfully",
		Status:  "Success",
	}
	return response, nil
}

func (us *StoreService) UpdateStoreStatus(req request.UpdateStoreStatusRequest) (response.DataResponse, error) {
	store := mapper.MapUpdateStoreStatusRequest(req)

	err := us.storeRepo.UpdateStoreStatus(store)
	if err != nil {
		return response.DataResponse{}, err
	}

	response := response.DataResponse{
		Data:    true,
		Message: "Store status updated successfully",
		Status:  "Success",
	}
	return response, nil
}
