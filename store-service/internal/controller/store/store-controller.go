package storeController

import (
	"net/http"
	storeService "store-service/internal/service/store"
	"store-service/pkg/transport/http/exception"
	"store-service/pkg/transport/http/request"

	"github.com/gin-gonic/gin"
)

type StoreController struct {
	storeServiceInstance *storeService.StoreService
}

func NewStoreController() *StoreController {
	return &StoreController{
		storeServiceInstance: storeService.NewStoreService(),
	}
}

func (cs *StoreController) GetStoreController(c *gin.Context) {
	var request request.GetStoreRequest
	id := c.Param("id")

	request.Store_id = id
	responce, err := cs.storeServiceInstance.GetStoreByID(request)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) GetAllStoresController(c *gin.Context) {
	responce, err := cs.storeServiceInstance.GetAllStores()
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) CreateStoreController(c *gin.Context) {
	var req request.CreateStoreRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.CreateStore(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) UpdateStoreController(c *gin.Context) {
	var req request.UpdateStoreRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.UpdateStore(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

// start functions of category

func (cs *StoreController) GetAllCategoriesController(c *gin.Context) {
	responce, err := cs.storeServiceInstance.GetAllCategories()
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) CreateCategoryController(c *gin.Context) {
	var req request.CreateCategoryRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.CreateCategory(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) UpdateCategoryController(c *gin.Context) {
	var req request.UpdateCategoryRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.UpdateCategory(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) UpdateStoreLocationController(c *gin.Context) {
	var req request.UpdateStoreLocationRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.UpdateStoreLocation(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}

func (cs *StoreController) UpdateStoreStatusController(c *gin.Context) {
	var req request.UpdateStoreStatusRequest
	if exception := exception.HandleBindUriError(c, c.ShouldBindJSON(&req)); exception != nil {
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}

	responce, err := cs.storeServiceInstance.UpdateStoreStatus(req)
	if err != nil {
		exception := exception.BadException{
			Message: err.Error(),
		}
		c.JSON(http.StatusBadRequest, exception)
		return
	}
	c.JSON(http.StatusOK, responce)
}
