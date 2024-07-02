package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebResponse struct {
	Code     int       `json:"code" example:"200"`
	Message  string    `json:"message" example:"Here is your data!"`
	Data     any       `json:"data,omitempty"`
	MetaData *MetaData `json:"metadata,omitempty"`
}

type MetaData struct {
	Page  int `json:"page" example:"1"`
	Size  int `json:"size" example:"10"`
	Total int `json:"total" example:"13"`
}

func NewResponse(httpStatus int, message string, data any) WebResponse {
	return WebResponse{
		Code:    httpStatus,
		Message: message,
		Data:    data,
	}

}

func NewResponseWithPagination(message string, data any, metaData *MetaData) WebResponse {
	return WebResponse{
		Code:     http.StatusOK,
		Message:  message,
		Data:     data,
		MetaData: metaData,
	}

}

func WriteToResponseBody(ctx *gin.Context, webResponse WebResponse) {
	ctx.JSON(webResponse.Code, webResponse)
}
