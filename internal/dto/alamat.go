package dto

import "time"

type Alamat struct {
	ID           int64     `json:"id"`
	IdUser       int64     `json:"id_user"`
	JudulAlamat  string    `json:"judul_alamat"`
	NamaPenerima string    `json:"nama_penerima"`
	NoTelp       string    `json:"no_telp"`
	DetailAlamat string    `json:"detail_alamat"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateAlamatRequest struct {
	IdUser       int64  `json:"id_user" validate:"required"`
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type UpdateAlamatRequest struct {
	ID           int64   `json:"id" validate:"required"`
	JudulAlamat  *string `json:"judul_alamat"`  // pakai pointer untuk optional field
	NamaPenerima *string `json:"nama_penerima"` // pakai pointer untuk optional field
	NoTelp       *string `json:"no_telp"`       // pakai pointer untuk optional field
	DetailAlamat *string `json:"detail_alamat"` // pakai pointer untuk optional field
}
