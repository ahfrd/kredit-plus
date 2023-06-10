package response

type GetDataTenorEntity struct {
	Nik        string `json:"nik"`
	Konsumen   string `json:"konsumen"`
	LimitTenor int    `json:"limitTenor"`
}

type GetDataKonsumenByNikEntity struct {
	Nik          string `json:"nik"`
	FullName     string `json:"fullName"`
	LegalName    string `json:"legalName"`
	TempatLahir  string `json:"tempatLahir"`
	TanggalLahir string `json:"tanggalLahir"`
	Gaji         int    `json:"gaji"`
	FotoKtp      string `json:"fotoKtp"`
	FotoSelfie   string `json:"fotoSelfie"`
}

type GeneralResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GeneralResponseTest struct {
	Code string              `json:"code"`
	Msg  string              `json:"msg"`
	Data ResponseDataInquiry `json:"data"`
}

type ResponseDataInquiry struct {
	Nik            string `json:"nik"`
	Tenor          string `json:"tenor"`
	Otr            int    `json:"otr"`
	Bunga          int    `json:"bunga"`
	NamaAsset      string `json:"namaAsset"`
	AdminFee       int    `json:"adminFee"`
	ProdukCategory string `json:"produkCategory"`
}
