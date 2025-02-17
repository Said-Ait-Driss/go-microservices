package httpHandler

import (
	categoryCtr "product-service/internal/controller/category"
	offerCtr "product-service/internal/controller/offer"
	productCtr "product-service/internal/controller/product"

	"github.com/gin-gonic/gin"
)

func ProductHandler(router *gin.Engine) {
	router.GET("/products/:product_id", productCtr.GetProductById)

	router.GET("/products", productCtr.GetProducts)

	router.POST("/products", productCtr.CreateProduct)

	router.PUT("/products", productCtr.UpdateProduct)

	router.DELETE("/products/:product_id", productCtr.DeleteProduct)

	// products that has offers
	router.GET("/products/has-offers", productCtr.GetProductsThatHasOffers)

	router.GET("/products/store/:store_id/has-offers", productCtr.GetProductsOfStoreThatHasOffers)

	router.GET("/store/products/:store_id", productCtr.GetProductsByStore)

	// getting products count of store
	router.GET("/store/products/count/:store_id", productCtr.GetProductsCountByStore)

	// filter all products based on : name, price range, date ( start, end )
	router.GET("/products/filter/:store_id", productCtr.GetStoreProductsFilter)
	router.GET("/products/filter", productCtr.GetProductsFilter)

	// get products by category
	router.GET("/products/category/:category_id", productCtr.GetProductsByCategory)

}

// for offers

func OfferHandler(router *gin.Engine) {
	router.GET("/offers/:offer_id", offerCtr.GetOfferById)
	router.GET("/offers", offerCtr.GetOffers)
	router.POST("/offers", offerCtr.CreateOffer)

	router.PUT("/offers", offerCtr.UpdateOffer)

	router.DELETE("/offers/:offer_id", offerCtr.DeleteOffer)

}

// for category

func CategoryHandler(router *gin.Engine) {
	router.GET("/categories/:category_id", categoryCtr.GetCategoryById)
	router.GET("/categories", categoryCtr.GetCategories)
	router.POST("/categories", categoryCtr.CreateCategory)

	router.PUT("/categories", categoryCtr.UpdateCategory)

	router.DELETE("/categories/:category_id", categoryCtr.DeleteCategory)
}
