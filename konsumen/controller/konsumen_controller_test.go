package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"kredit-plus/entity/mock"
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"
	"kredit-plus/konsumen/controller"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestKonsumenController_Inquiry(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.KonsumenController{}

	router := gin.Default()
	router.POST("/inquiry", controller.Inquiry)

	t.Run("Successful Inquiry", func(t *testing.T) {
		bodyReq := request.InquiryRequest{
			Nik:            "412321321313",
			Tenor:          "5",
			Otr:            20000000,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/inquiry", bytes.NewReader(requestData))

		controller.KonsumenService = &mock.MockKonsumenService{
			InquiryFn: func(ctx *gin.Context, req *request.InquiryRequest, requestId string) (*response.GeneralResponse, error) {
				response := &response.GeneralResponse{
					Code: "200",
					Msg:  "Sukses",
				}
				return response, nil
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)

		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

	})

	t.Run("Failed Inquiry - BindJSON Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/inquiry", bytes.NewReader([]byte("s-json")))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})

	t.Run("Failed Inquiry - Service Error", func(t *testing.T) {
		// Mock request body
		bodyReq := request.InquiryRequest{
			Nik:            "41231231241242",
			Tenor:          "5",
			Otr:            20000000,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}
		requestData, _ := json.Marshal(bodyReq)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/inquiry", bytes.NewReader(requestData))

		controller.KonsumenService = &mock.MockKonsumenService{
			InquiryFn: func(ctx *gin.Context, req *request.InquiryRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service error")
			},
		}

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusBadRequest, w.Code)

	})
}

func TestKonsumenController_Payment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	controller := &controller.KonsumenController{}

	router := gin.Default()
	router.POST("/payment", controller.Payment)

	t.Run("Successful Payment", func(t *testing.T) {
		// Mock request body
		bodyReq := request.PaymentRequest{
			Nik:            "41231231241242",
			Tenor:          "5",
			Otr:            20000000,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}
		requestData, _ := json.Marshal(bodyReq)

		// Create a test context with request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/payment", bytes.NewReader(requestData))

		// Mock the KonsumenService method to return a successful response
		controller.KonsumenService = &mock.MockKonsumenService{
			PaymentFn: func(ctx *gin.Context, req *request.PaymentRequest, requestId string) (*response.GeneralResponse, error) {
				// Mock the response data
				response := &response.GeneralResponse{
					// Populate the response data as needed
				}
				return response, nil
			},
		}

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)

		// Parse the response body
		var response response.GeneralResponse
		_ = json.Unmarshal(w.Body.Bytes(), &response)

		// Assert the response as needed
		// assert.Equal(t, expectedResponse, response)
	})

	t.Run("Failed Payment - BindJSON Error", func(t *testing.T) {
		// Create a test context with invalid request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/payment", bytes.NewReader([]byte("invalid-json")))

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Assert the error message or response body as needed
	})

	t.Run("Failed Payment - Service Error", func(t *testing.T) {
		// Mock request body
		bodyReq := request.PaymentRequest{
			Nik:            "41231231241242",
			Tenor:          "5",
			Otr:            20000000,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}
		requestData, _ := json.Marshal(bodyReq)

		// Create a test context with request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/payment", bytes.NewReader(requestData))

		// Mock the KonsumenService method to return an error
		controller.KonsumenService = &mock.MockKonsumenService{
			PaymentFn: func(ctx *gin.Context, req *request.PaymentRequest, requestId string) (*response.GeneralResponse, error) {
				return nil, errors.New("service errors")
			},
		}

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusBadRequest, w.Code)

		// Assert the error message or response body as needed
	})
}
