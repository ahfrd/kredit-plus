package mock

import (
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"

	"github.com/gin-gonic/gin"
)

type MockKonsumenService struct {
	InquiryFn func(ctx *gin.Context, req *request.InquiryRequest, requestId string) (*response.GeneralResponse, error)
	PaymentFn func(ctx *gin.Context, req *request.PaymentRequest, requestId string) (*response.GeneralResponse, error)
}

func (m *MockKonsumenService) Inquiry(ctx *gin.Context, req *request.InquiryRequest, requestId string) (*response.GeneralResponse, error) {
	if m.InquiryFn != nil {
		return m.InquiryFn(ctx, req, requestId)
	}
	return nil, nil
}

func (m *MockKonsumenService) Payment(ctx *gin.Context, req *request.PaymentRequest, requestId string) (*response.GeneralResponse, error) {
	if m.PaymentFn != nil {
		return m.PaymentFn(ctx, req, requestId)
	}
	return nil, nil
}
