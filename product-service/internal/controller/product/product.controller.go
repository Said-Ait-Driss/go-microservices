package controller

import (
	"fmt"
	"net/http"
	service "product-service/internal/service/product"
	"product-service/pkg/endpoint"
	"product-service/pkg/transport/http/request"
	"product-service/pkg/transport/http/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var productEndpoints endpoint.Endpoints

func init() {

	productService := &service.ProductServiceImpl{}

	productEndpoints = endpoint.MakeProductEndpoints(productService)

}

/**
 * get product by id
 * @func
 * @param {Context} ctx current context
 */
func GetProductById(ctx *gin.Context) {

	product_id := ctx.Param("product_id")

	_request := request.GetProductRequest{ID: product_id}
	_response, _ := productEndpoints.GetProductEndpoint(ctx, _request)

	res := _response.(response.GetProductResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product": res.Product})
}

func GetProducts(ctx *gin.Context) {
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

	_request := request.GetProductsRequest{Page: page, PageSize: pageSize}

	_response, _ := productEndpoints.GetProductsEndpoit(ctx, _request)

	res := _response.(response.GetProductsResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})

}

// create product function
func CreateProduct(ctx *gin.Context) {
	var request request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := productEndpoints.CreateProductEndpoint(ctx, request)

	res := _response.(response.CreateProductResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product": res.Product})

}

// update product function
func UpdateProduct(ctx *gin.Context) {
	var request request.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := productEndpoints.UpdateProductEndpoint(ctx, request)

	res := _response.(response.UpdateProductResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product": res.Product})
}

// delete product function

func DeleteProduct(ctx *gin.Context) {
	product_id := ctx.Param("product_id")
	_request := request.DeleteProductRequest{ID: product_id}
	_response, _ := productEndpoints.DeleteProductEndpoint(ctx, _request)

	res := _response.(response.DeletProductResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"product_id": product_id})
}

func GetProductsByStore(ctx *gin.Context) {
	storeId := ctx.Param("store_id")

	_request := request.GetProductsByStoreRequest{Store_id: storeId}
	_response, _ := productEndpoints.GetProductsByStoreEndpoint(ctx, _request)

	res := _response.(response.GetProductsResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})

}

// get products count of store
func GetProductsCountByStore(ctx *gin.Context) {
	storeId := ctx.Param("store_id")
	_request := request.GetProductsCountByStoreRequest{Store_id: storeId}
	_response, _ := productEndpoints.GetProductsCountByStoreEndpoint(ctx, _request)
	res := _response.(response.GetProductsCountResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"count": res.Count})
}

// Get Products That Has Offers
func GetProductsThatHasOffers(ctx *gin.Context) {
	pageParam := ctx.DefaultQuery("page", "1")
	pageSizeParam := ctx.DefaultQuery("page_size", "10")

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

	_request := request.GetProductsRequest{Page: page, PageSize: pageSize}
	_response, _ := productEndpoints.GetProductsThatHasOffersEndpoint(ctx, _request)
	res := _response.(response.GetProductsResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})
}

// Get Products of store that Has Offers
func GetProductsOfStoreThatHasOffers(ctx *gin.Context) {
	storeId := ctx.Param("store_id")
	pageParam := ctx.DefaultQuery("page", "1")
	pageSizeParam := ctx.DefaultQuery("page_size", "10")

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

	_request := request.GetProductsOfStoreThatHasOffersRequest{Store_id: storeId, Page: page, PageSize: pageSize}
	_response, _ := productEndpoints.GetProductsOfStoreThatHasOffersEndpoint(ctx, _request)
	res := _response.(response.GetProductsResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})
}

func GetStoreProductsFilter(ctx *gin.Context) {
	storeId := ctx.Param("store_id")

	nameQuery := ctx.DefaultQuery("name", "")
	minPriceQuery := ctx.DefaultQuery("minPrice", "0")
	maxPriceQuery := ctx.DefaultQuery("minPrice", "1000000000000")
	startDateQuery := ctx.DefaultQuery("startDate", "2006-01-02")
	endDateQuery := ctx.DefaultQuery("endDate", "2090-01-02")
	pageParam := ctx.DefaultQuery("page", "1")
	pageSizeParam := ctx.DefaultQuery("page_size", "10")

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

	minPrice, err := strconv.ParseFloat(minPriceQuery, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceQuery, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	// Convert string to time.Time
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateQuery)
	if err != nil {
		fmt.Println("Error converting string to time:", err)
		return
	}

	endDate, err := time.Parse(layout, endDateQuery)
	if err != nil {
		fmt.Println("Error converting string to time:", err)
		return
	}

	_request := request.GetStoreProductsFilterRequest{Store_id: storeId, Name: nameQuery, MinPrice: minPrice, MaxPrice: maxPrice, StartDate: startDate, EndDate: endDate, Page: page, PageSize: pageSize}
	_response, _ := productEndpoints.GetStoreProductsFilterEndpoint(ctx, _request)
	res := _response.(response.GetProductsResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})
}

func GetProductsFilter(ctx *gin.Context) {
	nameQuery := ctx.DefaultQuery("name", "")
	minPriceQuery := ctx.DefaultQuery("minPrice", "0")
	maxPriceQuery := ctx.DefaultQuery("minPrice", "1000000000000")
	startDateQuery := ctx.DefaultQuery("startDate", "2006-01-02")
	endDateQuery := ctx.DefaultQuery("endDate", "2090-01-02")
	neighborhoodQuery := ctx.DefaultQuery("neighborhood", "")
	storeTitleQuery := ctx.DefaultQuery("storeTitle", "")
	pageQuery := ctx.DefaultQuery("page", "1")
	pageSizeQuery := ctx.DefaultQuery("page_size", "10")

	// Convert string to int
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeQuery)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	minPrice, err := strconv.ParseFloat(minPriceQuery, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	maxPrice, err := strconv.ParseFloat(maxPriceQuery, 64)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": err})
		return
	}

	// Convert string to time.Time
	layout := "2006-01-02"
	startDate, err := time.Parse(layout, startDateQuery)
	if err != nil {
		fmt.Println("Error converting string to time:", err)
		return
	}

	endDate, err := time.Parse(layout, endDateQuery)
	if err != nil {
		fmt.Println("Error converting string to time:", err)
		return
	}

	_request := request.GetProductsFilterRequest{Name: nameQuery, MinPrice: minPrice, MaxPrice: maxPrice, StartDate: startDate, EndDate: endDate, Neighborhood: neighborhoodQuery, StoreTitle: storeTitleQuery, Page: page, PageSize: pageSize}
	_response, _ := productEndpoints.GetProductsFilterEndpoint(ctx, _request)
	res := _response.(response.GetProductsResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})
}

// get products by category
func GetProductsByCategory(ctx *gin.Context) {

	categoryId := ctx.Param("category_id")

	_request := request.GetProductByCategoryRequest{CategoryId: categoryId}
	_response, _ := productEndpoints.GetProductsByCategoryEndpoint(ctx, _request)
	res := _response.(response.GetProductsResponse)
	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"products": res.Products})
}
