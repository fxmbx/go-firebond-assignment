package dto

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data         interface{} `json:"data"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"error,omitempty"`
}

func buildResponse(data interface{}, message string) Response {
	return Response{
		Data:    data,
		Message: message,
	}
}

func buildErrorResponse(data interface{}, message string, err error) Response {
	return Response{
		Data:         data,
		Message:      message,
		ErrorMessage: err.Error(),
	}
}

func InternalServerErrorResponse(ctx *gin.Context, message string, err error) {
	ctx.JSON(http.StatusInternalServerError, buildErrorResponse(nil, message, err))
}
func BadRequestResponse(ctx *gin.Context, message string, err error) {

	ctx.JSON(http.StatusBadRequest, buildErrorResponse(nil, message, err))
}
func UnauthrizedResponse(ctx *gin.Context, message string, err error) {
	if len(message) == 0 {
		message = "you are not authorized to perform this action"
	}
	ctx.JSON(http.StatusUnauthorized, buildErrorResponse(nil, message, err))
}

func OkResponse(ctx *gin.Context, message string, data interface{}) {
	if len(message) < 1 {
		message = "request completed"
	}
	ctx.JSON(http.StatusOK, buildResponse(data, message))
}

func CustomResponse(ctx *gin.Context, message string, err error, statusCode int, success bool, data interface{}) {
	if success {
		ctx.JSON(statusCode, buildResponse(data, message))
	} else {
		ctx.JSON(statusCode, buildErrorResponse(nil, message, err))
	}
}
