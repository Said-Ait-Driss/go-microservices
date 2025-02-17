package mapper

import (
	entities "store-service/pkg/entities"
	"store-service/pkg/transport/http/request"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func MapCreateStoreRequest(request request.CreateStoreRequest) entities.Store {
    return entities.Store{
		ID:         		primitive.NewObjectID(),
		Name:              	request.Name,
		Description:       	request.Description,
		Phone: 		      	request.Phone,
		Covers:            	request.Covers,
		Longitude:         	request.Longitude,
		Latitude:          	request.Latitude,
		Client_id:         	request.Client_id,
		Adresse:           	request.Adresse,
		City:    	     	request.City,
		Neighborhood:     	request.Neighborhood,
		Status:				request.Status,
		Open_from:           request.Open_from,
		Open_to:             request.Open_to,
		Category:          	request.Category,
		Created_at:     	time.Now(),
	}
}

func MapUpdateStoreRequest(request request.UpdateStoreRequest) entities.Store {
	objectID, err := primitive.ObjectIDFromHex(request.ID)
    if err != nil {
        return entities.Store{}
    }
    return entities.Store{
		ID:         		objectID,
		Name:              	request.Name,
		Description:       	request.Description,
		Phone: 		      	request.Phone,
		Covers:            	request.Covers,
		Longitude:         	request.Longitude,
		Latitude:          	request.Latitude,
		Client_id:         	request.Client_id,
		Adresse:           	request.Adresse,
		City:    	     	request.City,
		Neighborhood:     	request.Neighborhood,
		Category:          	request.Category,
		Updated_at:     	time.Now(),
	}
}

// start mapping category entity

func MapCreateCategoryRequest(request request.CreateCategoryRequest) entities.Categories {
    return entities.Categories{
		ID:         		primitive.NewObjectID(),
		Code:              	request.Code,
		Libelle:       		request.Libelle,
		Created_at:     	time.Now(),
	}
}

func MapUpdateCategoryRequest(request request.UpdateCategoryRequest) entities.Categories {
	objectID, err := primitive.ObjectIDFromHex(request.ID)
    if err != nil {
        return entities.Categories{}
    }
    return entities.Categories{
		ID:         		objectID,
		Code:              	request.Code,
		Libelle:       		request.Libelle,
		Updated_at:     	time.Now(),
	}
}

func MapUpdateStoreLocationRequest(request request.UpdateStoreLocationRequest) entities.Store {
	objectID, err := primitive.ObjectIDFromHex(request.ID)
    if err != nil {
        return entities.Store{}
    }
    return entities.Store{
		ID:         		objectID,
		Longitude:         	request.Longitude,
		Latitude:          	request.Latitude,
	}
}

func MapUpdateStoreStatusRequest(request request.UpdateStoreStatusRequest) entities.Store {
	objectID, err := primitive.ObjectIDFromHex(request.ID)
    if err != nil {
        return entities.Store{}
    }
    return entities.Store{
		ID:         		objectID,
		Status:         	request.Status,
	}
}