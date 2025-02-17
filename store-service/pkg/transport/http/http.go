package http

import (
	"github.com/gin-gonic/gin"
	storeController "store-service/internal/controller/store"
)

var storeControllerInstance = storeController.NewStoreController()

// function to handle store api
func StoreHandler(group *gin.RouterGroup) {
	group.GET("/get/:id", storeControllerInstance.GetStoreController)
	group.GET("/get-all", storeControllerInstance.GetAllStoresController)
	group.POST("/create", storeControllerInstance.CreateStoreController)
	group.POST("/update", storeControllerInstance.UpdateStoreController)
	group.POST("/change-location", storeControllerInstance.UpdateStoreLocationController)
	group.POST("/change-status", storeControllerInstance.UpdateStoreStatusController)
}

func CategoryHandler(group *gin.RouterGroup) {
	group.POST("/create", storeControllerInstance.CreateCategoryController)
	group.GET("/get-all", storeControllerInstance.GetAllCategoriesController)
	group.POST("/update", storeControllerInstance.UpdateCategoryController)
}
