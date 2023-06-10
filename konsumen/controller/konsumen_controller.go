package controller

import (
	"encoding/json"
	"fmt"
	"kredit-plus/entity"
	"kredit-plus/entity/request"
	"kredit-plus/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type KonsumenController struct {
	KonsumenService entity.KonsumenService
}

func NewKonsumenController(konsumenService *entity.KonsumenService) KonsumenController {
	return KonsumenController{KonsumenService: *konsumenService}
}

func (c *KonsumenController) Inquiry(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.InquiryRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.KonsumenService.Inquiry(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (c *KonsumenController) Payment(ctx *gin.Context) {
	requestId := guuid.New()
	var bodyReq request.PaymentRequest
	if err := ctx.BindJSON(&bodyReq); err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	requestData, err := json.Marshal(bodyReq)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStart := helpers.LogRequest(ctx, string(requestData), requestId.String())
	fmt.Println(logStart)

	response, err := c.KonsumenService.Payment(ctx, &bodyReq, requestId.String())
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
