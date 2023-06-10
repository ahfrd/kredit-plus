package mock

import (
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"
)

type MockKonsumenRepository struct {
	GetDataKonsumenByNikFn func(nik string) (*response.GetDataKonsumenByNikEntity, error)
	GetTenorByNikFn        func(nik string, tenor string) (*response.GetDataTenorEntity, error)
	InsertTransactionFn    func(req *request.InsertTransactionRequest) error
	UpdateTenorFn          func(tenor string, limitTenor int, nik string) error
}

func (m *MockKonsumenRepository) GetDataKonsumenByNik(nik string) (*response.GetDataKonsumenByNikEntity, error) {
	return m.GetDataKonsumenByNikFn(nik)
}

func (m *MockKonsumenRepository) GetTenorByNik(nik string, tenor string) (*response.GetDataTenorEntity, error) {
	return m.GetTenorByNikFn(nik, tenor)
}

func (m *MockKonsumenRepository) InsertTransaction(req *request.InsertTransactionRequest) error {
	return m.InsertTransactionFn(req)
}

func (m *MockKonsumenRepository) UpdateTenor(tenor string, limitTenor int, nik string) error {
	return m.UpdateTenorFn(tenor, limitTenor, nik)
}
