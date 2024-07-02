package exception

import (
	"log"
	"net/http"

	"github.com/adityarifqyfauzan/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler(ctx *gin.Context, err any) {
	log.Println(err)
	if notFoundError(ctx, err) {
		return
	}
	if credentialError(ctx, err) {
		return
	}
	if unprocessableEntity(ctx, err) {
		return
	}
	if badRequestError(ctx, err) {
		return
	}
	internalServerError(ctx)
}

func unprocessableEntity(ctx *gin.Context, err any) bool {
	exception, ok := err.(UnprocessableEntityException)
	if !ok {
		return false
	}

	webResponse := utils.NewResponse(http.StatusUnprocessableEntity, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func notFoundError(ctx *gin.Context, err any) bool {
	exception, ok := err.(NotFoundException)
	if !ok {
		return false
	}

	webResponse := utils.NewResponse(http.StatusNotFound, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func credentialError(ctx *gin.Context, err any) bool {
	exception, ok := err.(CredentialException)
	if !ok {
		return false
	}

	webResponse := utils.NewResponse(http.StatusUnauthorized, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func badRequestError(ctx *gin.Context, err any) bool {
	exception, ok := err.(BadRequestException)
	if !ok {
		return false
	}

	webResponse := utils.NewResponse(http.StatusBadRequest, exception.Error, nil)
	utils.WriteToResponseBody(ctx, webResponse)

	return true
}

func internalServerError(ctx *gin.Context) bool {
	webResponse := utils.NewResponse(http.StatusInternalServerError, "Internal Server Error: Please try in few minutes", nil)
	utils.WriteToResponseBody(ctx, webResponse)
	return true
}
