package service_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kredit-plus/entity/mock"
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"
	"kredit-plus/helpers"
	"kredit-plus/konsumen/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestKonsumenService_Inquiry(t *testing.T) {
	//gagal test ini karena di akibatkan menggunakan enkripsi metode RSA yang dimana enkripsinya berubah ubah
	gin.SetMode(gin.TestMode)

	service := &service.KonsumenService{
		KonsumenRepository: &mock.MockKonsumenRepository{
			GetDataKonsumenByNikFn: func(nik string) (*response.GetDataKonsumenByNikEntity, error) {
				if nik == "4684671205414376" {
					return &response.GetDataKonsumenByNikEntity{
						Nik: "4684671205414376",
					}, nil
				} else {
					return nil, errors.New("Data konsumen not found")
				}
			},
			GetTenorByNikFn: func(nik string, tenor string) (*response.GetDataTenorEntity, error) {
				if nik == "4684671205414376" && tenor == "3" {
					return &response.GetDataTenorEntity{
						Nik:        "4684671205414376",
						LimitTenor: 10000,
						// Populate the other fields as needed
					}, nil
				} else {
					return nil, errors.New("Data tenor not found")
				}
			},
		},
	}

	router := gin.Default()
	router.POST("/inquiry", func(ctx *gin.Context) {
		var req request.InquiryRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uid := "sample-uid"

		res, err := service.Inquiry(ctx, &req, uid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
	t.Run("Successful Inquiry", func(t *testing.T) {
		bodyReq := request.InquiryRequest{
			Nik:            "4684671205414376",
			Tenor:          "3",
			Otr:            500,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/inquiry", nil)
		ctx.Set("Content-Type", "application/json")
		ctx.Set("uid", "sample-uid")
		ctx.Set("request_id", "sample-request-id")
		ctx.Set("uid", "sample-uid")
		ctx.Set("request_id", "sample-request-id")
		ctx.Set("uid", "sample-uid")
		ctx.Set("request_id", "sample-request-id")
		body, _ := json.Marshal(bodyReq)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		router.ServeHTTP(w, ctx.Request)

		assert.Equal(t, http.StatusOK, w.Code)
		fmt.Println("//gagal test ini karena di akibatkan menggunakan enkripsi metode RSA yang dimana enkripsinya berubah ubah")
		var resData response.ResponseDataInquiry
		resData.Nik = helpers.EncryptRSAKey(bodyReq.Nik)
		resData.Tenor = "3"
		resData.Otr = 500
		resData.Bunga = 5
		resData.NamaAsset = "Laptop"
		resData.AdminFee = 500000
		resData.ProdukCategory = "WG"
		// Assert the response body
		expectedRes := response.GeneralResponseTest{
			Code: "200",
			Msg:  "Sukses",
			Data: resData,
		}
		actualRes := response.GeneralResponseTest{}
		_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("Data Konsumen Not Found", func(t *testing.T) {
		// Mock request body
		bodyReq := request.InquiryRequest{
			Nik:            "0987654321",
			Tenor:          "5",
			Otr:            20000000,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}

		// Create a test context with request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/inquiry", nil)
		ctx.Set("Content-Type", "application/json")
		body, _ := json.Marshal(bodyReq)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert the response body
		expectedRes := response.GeneralResponse{
			Code: "400",
			Msg:  "Data konsumen not found",
			Data: nil,
		}
		actualRes := response.GeneralResponse{}
		_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
		assert.Equal(t, expectedRes, actualRes)
	})
}

func TestKonsumenService_Payment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	service := &service.KonsumenService{
		KonsumenRepository: &mock.MockKonsumenRepository{
			GetTenorByNikFn: func(nik string, tenor string) (*response.GetDataTenorEntity, error) {
				// Mock the response data
				if nik == "4684671205414376" && tenor == "4" {
					return &response.GetDataTenorEntity{
						Nik:        "4684671205414376",
						LimitTenor: 5000,
						// Populate the other fields as needed
					}, nil
				} else {
					return nil, errors.New("Data tenor not found")
				}
			},
			InsertTransactionFn: func(req *request.InsertTransactionRequest) error {
				// Mock the behavior of inserting transaction
				// You can perform additional assertions here if needed
				return nil
			},
			UpdateTenorFn: func(tenor string, limitTenor int, nik string) error {
				// Mock the behavior of updating tenor
				// You can perform additional assertions here if needed
				return nil
			},
		},
	}

	router := gin.Default()
	router.POST("/payment", func(ctx *gin.Context) {
		var params request.PaymentRequest
		if err := ctx.ShouldBindJSON(&params); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		uid := "sample-uid"

		// Call the service.Payment function and handle the response
		res, err := service.Payment(ctx, &params, uid)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, res)
	})
	t.Run("Successful Payment", func(t *testing.T) {
		// Mock request body
		bodyReq := request.PaymentRequest{
			Nik:            "Hobukcw1+L7LRq90tMwWOcdVq+ODd5mebpWlu9TkVHpmzdH24cczISPPNoH/xBCGHYT4qaV//4BBxAFUaKWpAsapFNUHB/w2pnQRlNdFbQIhVHNhOODYYyVnHFG5Ra2wNT05m7qywiqk0g4faMH6SaLkbm/l02Di/yJ3Z8Kl2ryNwgxzeMR//1YYCbBQyQmiXGmhzcQ+s3z/KoKLrSkOVHogurK61mu0kloFaJrOJrKlChk45UvE0wu9TpRC6D0qtqn25kqti6J4u+0pyQypmSD8AQXXtSvSO6cECDoOW7Jvq1qZJb+73G135WeYIbwNSrHpZgqV7zPiODQ5WMJ7Fg==",
			Tenor:          "4",
			Otr:            500,
			Bunga:          5,
			NamaAsset:      "Laptop",
			AdminFee:       500000,
			ProdukCategory: "WG",
		}

		// Create a test context with request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/payment", nil)
		ctx.Set("Content-Type", "application/json")

		// Set the CSRF token and X-Token header
		csrfToken := "sample-token"
		tokenAuth := helpers.OneWayEncrypt([]byte(csrfToken + "4684671205414376"))
		ctx.Request.Header.Set("X-CSRF-Token", csrfToken)
		ctx.Request.Header.Set("X-Token", tokenAuth)
		body, _ := json.Marshal(bodyReq)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert the response body
		expectedRes := response.GeneralResponse{
			Code: "200",
			Msg:  "Sukses",
			Data: nil,
		}
		actualRes := response.GeneralResponse{}
		_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
		assert.Equal(t, expectedRes, actualRes)
	})

	t.Run("Invalid CSRF Token", func(t *testing.T) {
		// Mock request body
		bodyReq := request.PaymentRequest{
			Nik:       "Hobukcw1+L7LRq90tMwWOcdVq+ODd5mebpWlu9TkVHpmzdH24cczISPPNoH/xBCGHYT4qaV//4BBxAFUaKWpAsapFNUHB/w2pnQRlNdFbQIhVHNhOODYYyVnHFG5Ra2wNT05m7qywiqk0g4faMH6SaLkbm/l02Di/yJ3Z8Kl2ryNwgxzeMR//1YYCbBQyQmiXGmhzcQ+s3z/KoKLrSkOVHogurK61mu0kloFaJrOJrKlChk45UvE0wu9TpRC6D0qtqn25kqti6J4u+0pyQypmSD8AQXXtSvSO6cECDoOW7Jvq1qZJb+73G135WeYIbwNSrHpZgqV7zPiODQ5WMJ7Fg==",
			Tenor:     "3",
			Otr:       500,
			Bunga:     5,
			NamaAsset: "Laptop",
			AdminFee:  500000,
		}

		// Create a test context with request body
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request, _ = http.NewRequest(http.MethodPost, "/payment", nil)
		ctx.Set("Content-Type", "application/json")

		// Set an invalid CSRF token and X-Token header
		csrfToken := "invalid-csrf-token"
		tokenAuth := "sample-token-auth"
		ctx.Set("X-CSRF-Token", csrfToken)
		ctx.Set("X-Token", tokenAuth)
		body, _ := json.Marshal(bodyReq)
		ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(body))

		// Execute the handler
		router.ServeHTTP(w, ctx.Request)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)

		// Assert the response body
		expectedRes := response.GeneralResponse{
			Code: "TNF",
			Msg:  "Invalid CSRF token",
			Data: nil,
		}
		actualRes := response.GeneralResponse{}
		_ = json.Unmarshal(w.Body.Bytes(), &actualRes)
		assert.Equal(t, expectedRes, actualRes)
	})
}
