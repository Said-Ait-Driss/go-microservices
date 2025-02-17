package controller

import (
	"net/http"
	service "product-service/internal/service/offer"
	"product-service/pkg/endpoint"
	"product-service/pkg/transport/http/request"
	"product-service/pkg/transport/http/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

var offerEndpoints endpoint.Endpoints

func init() {

	offerService := &service.OfferServiceImpl{}

	offerEndpoints = endpoint.MakeOfferEndpoints(offerService)

}

// create offer function
func CreateOffer(ctx *gin.Context) {
	var request request.CreateOfferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := offerEndpoints.CreateOfferEndpoint(ctx, request)

	res := _response.(response.CreateOfferResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"offer": res.Offer})
}

func GetOfferById(ctx *gin.Context) {

	offer_id := ctx.Param("offer_id")

	_request := request.GetOfferRequest{ID: offer_id}
	_response, _ := offerEndpoints.GetOfferEndpoint(ctx, _request)

	res := _response.(response.GetOfferResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"offer": res.Offer})
}

func GetOffers(ctx *gin.Context) {
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

	_request := request.GetOffersRequest{Page: page, PageSize: pageSize}

	_response, _ := offerEndpoints.GetOffersEndpoit(ctx, _request)

	res := _response.(response.GetOffersResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"offers": res.Offers})
}

func UpdateOffer(ctx *gin.Context) {
	var request request.UpdateOfferRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	_response, _ := offerEndpoints.UpdateOfferEndpoint(ctx, request)

	res := _response.(response.UpdateOfferResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"offer": res.Offer})
}

func DeleteOffer(ctx *gin.Context) {
	offer_id := ctx.Param("offer_id")
	_request := request.DeleteOfferRequest{ID: offer_id}
	_response, _ := offerEndpoints.DeleteOfferEndpoint(ctx, _request)

	res := _response.(response.DeletOfferResponse)

	if res.Err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": res.Err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"offer_id": offer_id})
}
