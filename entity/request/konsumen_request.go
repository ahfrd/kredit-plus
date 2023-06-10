package request

type InquiryRequest struct {
	Nik            string `json:"nik"`
	Tenor          string `json:"tenor"`
	Otr            int    `json:"otr"`
	Bunga          int    `json:"bunga"`
	NamaAsset      string `json:"namaAsset"`
	AdminFee       int    `json:"adminFee"`
	ProdukCategory string `json:"produkCategory"`
}

type PaymentRequest struct {
	Nik            string `json:"nik"`
	Tenor          string `json:"tenor"`
	Otr            int    `json:"otr"`
	Bunga          int    `json:"bunga"`
	NamaAsset      string `json:"namaAsset"`
	AdminFee       int    `json:"adminFee"`
	ProdukCategory string `json:"produkCategory"`
}

type InsertTransactionRequest struct {
	NomorKontrak  string `json:"nomorKontrak"`
	Otr           int    `json:"otr"`
	AdminFee      int    `json:"adminFee"`
	JumlahCicilan int    `json:"jumlahCicilan"`
	JumlahBunga   string `json:"jumlahBunga"`
	NamaAsset     string `json:"namaAsset"`
	Nik           string `json:"nik"`
}
