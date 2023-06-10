package entity

import (
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"

	"github.com/gin-gonic/gin"
)

type KonsumenService interface {
	Inquiry(ctx *gin.Context, request *request.InquiryRequest, uid string) (*response.GeneralResponse, error)
	Payment(ctx *gin.Context, request *request.PaymentRequest, uid string) (*response.GeneralResponse, error)
}

type KonsumenRepository interface {
	GetTenorByNik(nik string, tenor string) (*response.GetDataTenorEntity, error)
	GetDataKonsumenByNik(nik string) (*response.GetDataKonsumenByNikEntity, error)
	InsertTransaction(request *request.InsertTransactionRequest) error
	UpdateTenor(tenor string, limitUpdate int, nik string) error
}
