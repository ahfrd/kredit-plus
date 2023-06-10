package service

import (
	"fmt"
	"kredit-plus/entity"
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"
	"kredit-plus/helpers"
	"strconv"
	"strings"
	"time"

	guuid "github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type KonsumenService struct {
	KonsumenRepository entity.KonsumenRepository
}

func NewKonsumenService(konsumenRepository *entity.KonsumenRepository) entity.KonsumenService {
	return &KonsumenService{
		KonsumenRepository: *konsumenRepository,
	}
}

func (s *KonsumenService) Inquiry(ctx *gin.Context, request *request.InquiryRequest, uid string) (*response.GeneralResponse, error) {
	getKonsumen, err := s.KonsumenRepository.GetDataKonsumenByNik(request.Nik)
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	if getKonsumen.Nik == "" {
		return &response.GeneralResponse{
			Code: "DNF",
			Msg:  "Data konsumen not found",
		}, nil
	}

	getAmountTenor, err := s.KonsumenRepository.GetTenorByNik(request.Nik, request.Tenor)
	if err != nil {
		return &response.GeneralResponse{
			Code: "400",
			Msg:  err.Error(),
		}, nil
	}
	if getAmountTenor.Nik == "" {
		return &response.GeneralResponse{
			Code: "DNF",
			Msg:  "Data tenor not found",
		}, nil
	}
	if request.Otr > getAmountTenor.LimitTenor {
		return &response.GeneralResponse{
			Code: "LMT",
			Msg:  "Limit anda tidak mencukupi",
		}, nil
	}
	csrf := guuid.New().String()
	fmt.Println(csrf)
	fmt.Println(getKonsumen.Nik)
	initToken := helpers.OneWayEncrypt([]byte(csrf + getKonsumen.Nik))
	ctx.Header("X-Token", initToken)
	ctx.Header("X-CSRF-Token", csrf)

	var resData response.ResponseDataInquiry
	resData.Nik = request.Nik
	resData.Tenor = request.Tenor
	resData.Otr = request.Otr
	resData.Bunga = request.Bunga
	resData.NamaAsset = request.NamaAsset
	resData.AdminFee = request.AdminFee
	resData.ProdukCategory = request.ProdukCategory

	return &response.GeneralResponse{
		Code: "200",
		Msg:  "Sukses",
		Data: resData,
	}, nil
}

func (s *KonsumenService) Payment(ctx *gin.Context, params *request.PaymentRequest, uid string) (*response.GeneralResponse, error) {
	decryptNik := helpers.DecryptRSAKey(params.Nik)

	csrfToken := ctx.GetHeader("X-CSRF-Token")
	tokenAuth := ctx.GetHeader("X-Token")

	nikCsrf := helpers.OneWayEncrypt([]byte(csrfToken + decryptNik))
	if tokenAuth != nikCsrf {
		return &response.GeneralResponse{
			Code: "TNF",
			Msg:  "Invalid CSRF token",
			Data: nil,
		}, nil
	}
	tenor, _ := strconv.Atoi(params.Tenor)
	installmentBeforeInterest := params.Otr / tenor
	interest := fmt.Sprintf("%.f", float64(params.Otr)*(float64(params.Bunga)/100))
	integerInterest, _ := strconv.Atoi(interest)
	installment := installmentBeforeInterest + integerInterest
	prodNameKey := helpers.SeparateWord(params.ProdukCategory)

	timeNow := time.Now().Format("2006-01-02")
	ranNum, _ := helpers.GenerateRandomNumber(4)
	strNum := strconv.Itoa(ranNum)
	romanInt := helpers.IntegerToRoman(strNum)

	unique := strings.ReplaceAll(uid, "-", "")
	kontrakNum := strings.ToUpper(prodNameKey) + "/" + timeNow + "/" + romanInt + "/" + unique[0:7]
	fmt.Println(kontrakNum)
	resultChan := make(chan *response.GeneralResponse)
	errChan := make(chan error)

	// Goroutine pertama untuk mendapatkan tenor dan memeriksa limit
	go func() {
		getAmountTenor, err := s.KonsumenRepository.GetTenorByNik(decryptNik, params.Tenor)
		if err != nil {
			errChan <- err
			return
		}
		if getAmountTenor.Nik == "" {
			resultChan <- &response.GeneralResponse{
				Code: "DNF",
				Msg:  "Data tenors not found",
				Data: nil,
			}
			return
		}
		if params.Otr > getAmountTenor.LimitTenor {
			resultChan <- &response.GeneralResponse{
				Code: "LMT",
				Msg:  "Limit anda tidak mencukupi",
				Data: nil,
			}
			return
		}

		entityInsertTransaction := &request.InsertTransactionRequest{
			NomorKontrak:  kontrakNum,
			Otr:           params.Otr,
			AdminFee:      params.AdminFee,
			JumlahCicilan: installment,
			JumlahBunga:   interest,
			NamaAsset:     params.NamaAsset,
			Nik:           decryptNik,
		}
		fmt.Println(entityInsertTransaction)
		updatedLimit := getAmountTenor.LimitTenor - params.Otr

		go func() {
			err := s.KonsumenRepository.InsertTransaction(entityInsertTransaction)
			if err != nil {
				errChan <- err
				return
			}

			err = s.KonsumenRepository.UpdateTenor(params.Tenor, updatedLimit, decryptNik)
			if err != nil {
				errChan <- err
				return
			}

			resultChan <- &response.GeneralResponse{
				Code: "200",
				Msg:  "Sukses",
				Data: nil,
			}
		}()

	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		fmt.Println(errChan)
		return &response.GeneralResponse{
			Code: "DE",
			Msg:  err.Error(),
			Data: nil,
		}, nil
	}
}
