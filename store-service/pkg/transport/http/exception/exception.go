package exception

import (
	"github.com/gin-gonic/gin"
)

type BadException struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

type InvalidParametersException struct {
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewBadException(err error) BadException {
	return BadException{
		Error:   err.Error(),
		Message: "Bad Request",
		Status:  "400",
	}
}

func NewInvalidParametersException(err error) *InvalidParametersException {
	return &InvalidParametersException{
		Error:   err.Error(),
		Message: "Invalid request parameters",
		Status:  "400",
	}
}

func HandleBindUriError(c *gin.Context, err error) *InvalidParametersException {
	if err != nil {
		return NewInvalidParametersException(err)
	}
	return nil
}