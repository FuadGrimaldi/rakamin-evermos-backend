package dto

type LoginRequest struct {
	NoTelp   string `json:"no_telp"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterRequest struct {
	Nama         string `json:"nama"`
	Password     string `json:"password"`
	NoTelp       string `json:"no_telp"`
	TanggalLahir string `json:"tanggal_lahir"`
	JenisKelamin string `json:"jenis_kelamin"`
	Tentang      string `json:"tentang"`
	Pekerjaan    string `json:"pekerjaan"`
	Email        string `json:"email"`
	IdProvinsi   string `json:"id_provinsi"`
	IdKota       string `json:"id_kota"`
	Role         string `json:"role"`
}

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type LoginResponse struct {
	Nama         string   `json:"nama"`
	NoTelp       string   `json:"no_telp"`
	TanggalLahir string   `json:"tanggal_Lahir"`
	Tentang      string   `json:"tentang"`
	Pekerjaan    string   `json:"pekerjaan"`
	Email        string   `json:"email"`
	IdProvinsi   Province `json:"id_provinsi"`
	IdKota       City     `json:"id_kota"`
	Token        string   `json:"token"`
}
