package repository

import (
	"database/sql"
	"fmt"
	"kredit-plus/config"
	"kredit-plus/entity"
	"kredit-plus/entity/request"
	"kredit-plus/entity/response"
	"kredit-plus/helpers"
)

type konsumenRepository struct {
}

func NewKonsumenRepository() entity.KonsumenRepository {
	return &konsumenRepository{}
}

func (r *konsumenRepository) GetTenorByNik(nik string, tenor string) (*response.GetDataTenorEntity, error) {
	var resultDataTenor response.GetDataTenorEntity
	nik = helpers.SanitizeNumber(nik)
	tenor = helpers.SanitizeNumber(tenor)
	var dataNik sql.NullString
	var dataKonsumen sql.NullString
	var dataTenor sql.NullInt64

	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return nil, fmt.Errorf("failed Select SQL for limit_peminjaman : %v", err)
	}
	defer db.Close()
	var query string = fmt.Sprintf("select nik,konsumen,`%s` as limit_tenor from limit_peminjaman where nik = '%s'", tenor, nik)
	fmt.Println(query)
	db.QueryRow(query).Scan(
		&dataNik,
		&dataKonsumen,
		&dataTenor,
	)
	resultDataTenor.Nik = dataNik.String
	resultDataTenor.Konsumen = dataKonsumen.String
	resultDataTenor.LimitTenor = int(dataTenor.Int64)
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return &resultDataTenor, fmt.Errorf("failed Select SQL for limit_peminjaman : %v", err)
	}

	return &resultDataTenor, nil
}

func (r *konsumenRepository) GetDataKonsumenByNik(nik string) (*response.GetDataKonsumenByNikEntity, error) {
	var data response.GetDataKonsumenByNikEntity
	var Nik sql.NullString
	var fullName sql.NullString
	var legalName sql.NullString
	var tempatLahir sql.NullString
	var tanggalLahir sql.NullString
	var gaji sql.NullInt64
	var fotoKtp sql.NullString
	var fotoSelfie sql.NullString

	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return &data, err
	}
	nik = helpers.SanitizeNumber(nik)
	var query string = fmt.Sprintf(`select nik,full_name,legal_name,tempat_lahir,tanggal_lahir,gaji,foto_ktp,foto_selfie from konsumen where nik = "%s"`, nik)
	db.QueryRow(query).Scan(
		&Nik,
		&fullName,
		&legalName,
		&tempatLahir,
		&tanggalLahir,
		&gaji,
		&fotoKtp,
		&fotoSelfie,
	)
	fmt.Println(query)
	data.Nik = Nik.String
	data.FullName = fullName.String
	data.LegalName = legalName.String
	data.TempatLahir = tempatLahir.String
	data.TanggalLahir = tanggalLahir.String
	data.Gaji = int(gaji.Int64)
	data.FotoKtp = fotoKtp.String
	data.FotoSelfie = fotoSelfie.String
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed Select SQL for konsumen : %v", err)
	}

	return &data, nil
}
func (r *konsumenRepository) InsertTransaction(request *request.InsertTransactionRequest) error {
	var err error
	var res sql.Result
	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	queryInsert := "INSERT INTO transaction (nomor_kontrak,otr,admin_fee,jumlah_cicilan,jumlah_bunga,nama_asset,nik) values (?,?,?,?,?,?,?)"
	res, err = db.Exec(queryInsert, request.NomorKontrak, request.Otr, request.AdminFee, request.JumlahCicilan, request.JumlahBunga, request.NamaAsset, request.Nik)

	if err != nil {
		return fmt.Errorf("failed to insert error_general on transaction SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("terjadi kesalahan pada database")

	}
	fmt.Println(count)
	return nil
}

func (r *konsumenRepository) UpdateTenor(tenor string, limitUpdate int, nik string) error {
	tenor = helpers.SanitizeNumber(tenor)
	fmt.Println(limitUpdate)

	nik = helpers.SanitizeNumber(nik)
	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	var queryUpdate string = fmt.Sprintf("UPDATE limit_peminjaman set `%s` = ? where nik = ?", tenor)
	fmt.Println(queryUpdate)
	res, err := db.Exec(queryUpdate, limitUpdate, nik)
	defer db.Close()
	if err != nil {
		return fmt.Errorf("failed to limit_peminjaman SQL : %v", err)
	}

	if err != nil {
		return fmt.Errorf("failed to limit_peminjaman SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to populate status updated : %v", err)
	}
	fmt.Println(counter)
	if counter == 0 {
		return fmt.Errorf("failed to update tenor SQL")
	}
	return nil
}
