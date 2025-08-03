package dto

import "time"

type User struct {
	ID           int64     `json:"id"`
	Nama         string    `json:"nama"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IdProvinsi   int64     `json:"id_provinsi"`
	IdKota       int64     `json:"id_kota"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type NewUser struct {
	Nama         string    `json:"nama"`
	Password    string    `json:"password"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IdProvinsi   int64     `json:"id_provinsi"`
	IdKota       int64     `json:"id_kota"`
}

type UpdateUserRequest struct {
	Nama         *string    `json:"nama"`
	Password     *string    `json:"password"`
	NoTelp       *string    `json:"no_telp"`
	TanggalLahir *string `json:"tanggal_lahir"`
	JenisKelamin *string    `json:"jenis_kelamin"`
	Tentang      *string    `json:"tentang"`
	Pekerjaan    *string    `json:"pekerjaan"`
	Email        *string    `json:"email"`
	IdProvinsi   *string     `json:"id_provinsi"`
	IdCity   *string     `json:"id_kota"`
}

type ProfileResponse struct {
	Nama         string   `json:"nama"`
	NoTelp       string   `json:"no_telp"`
	TanggalLahir string   `json:"tanggal_Lahir"`
	Tentang      string   `json:"tentang"`
	Pekerjaan    string   `json:"pekerjaan"`
	Email        string   `json:"email"`
	IdProvinsi   Province `json:"id_provinsi"`
	IdKota       City     `json:"id_kota"`
}
