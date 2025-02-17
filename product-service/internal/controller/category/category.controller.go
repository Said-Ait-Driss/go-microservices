package controller

import (
	"net/http"
	service "product-service/internal/service/category"
	"product-service/pkg/endpoint"
	"product-service/pkg/transport/http/request"
	"product-service/pkg/transport/http/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

var categoryEndpoints endpoint.Endpoints

func init() {

	categoryService := &service.CategoryServiceImpl{}

	categoryEndpoints = endpoint.MakeCategoryEndpoints(categoryService)

}

// create category function
func CreateCategory(ctx *gin.Context) {
	var request request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := categoryEndpoints.CreateCategoryEndpoint(ctx, request)

	res := _response.(response.CreateCategoryResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"category": res.Category})
}

func GetCategoryById(ctx *gin.Context) {

	category_id := ctx.Param("category_id")

	_request := request.GetCategoryRequest{ID: category_id}
	_response, _ := categoryEndpoints.GetCategoryEndpoint(ctx, _request)

	res := _response.(response.GetCategoryResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"category": res.Category})
}

func GetCategories(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	pageSizeParam := ctx.DefaultQuery("pageSize", "10")

	// Convert string to int
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeParam)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	_request := request.GetCategoriesRequest{Page: page, PageSize: pageSize}

	_response, _ := categoryEndpoints.GetCategoriesEndpoit(ctx, _request)

	res := _response.(response.GetCategoriesResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"categories": res.Categories})
}

func UpdateCategory(ctx *gin.Context) {
	var request request.UpdateCategoryRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := categoryEndpoints.UpdateCategoryEndpoint(ctx, request)

	res := _response.(response.UpdateCategoryResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"category": res.Category})
}

func DeleteCategory(ctx *gin.Context) {
	category_id := ctx.Param("category_id")
	_request := request.DeleteCategoryRequest{ID: category_id}
	_response, _ := categoryEndpoints.DeleteCategoryEndpoint(ctx, _request)

	res := _response.(response.DeletCategoryResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"category_id": category_id})
}
