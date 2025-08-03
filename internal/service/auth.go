package service

import (
	"Evermos-Virtual-Intern/config"
	"Evermos-Virtual-Intern/internal/common"
	"Evermos-Virtual-Intern/internal/dto"
	"Evermos-Virtual-Intern/internal/entity"
	"Evermos-Virtual-Intern/internal/repository"
	"Evermos-Virtual-Intern/internal/util"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
	Register(ctx context.Context, request dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type authService struct {
	cfg        *config.Config
	repository repository.UserRepository
	tokoRepo     repository.TokoRepository 
}

func NewAuthService(cfg *config.Config, repository repository.UserRepository, tokoRepo repository.TokoRepository) AuthService {
	return &authService{cfg, repository, tokoRepo}
}

func (u *authService) Login(ctx context.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := u.repository.FindByContact(ctx, request.NoTelp)

	if err != nil {
		return nil, errors.New("No Telepon/password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))

	if err != nil {
		return nil, errors.New("No Telepon/password salah")
	}

	token, err := common.GenerateAccessToken(ctx, user)
	if err != nil {
		return nil, err
	}

	province, err := util.GetProvinceByID(user.IdProvinsi)
	if err != nil {
		return nil, err
	}

	city, err := util.GetCityByID(user.IdProvinsi, user.IdKota)
	if err != nil {
		return nil, err
	}


	response := &dto.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir.Format("02/01/2006"),
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi: dto.Province{
			ID:   province.ID,
			Name: province.Name,
		},
		IdKota: dto.City{
			ID:         city.ID,
			ProvinceID: city.ProvinceID,
			Name:       city.Name,
		},
		Token: token,
	}

	return response, nil
}

func (u *authService) Register(ctx context.Context, request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	if err != nil {
		return nil, err
	}

	role := "user"

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Lengkapi data user dari request
	user := &entity.User{
		Nama:         request.Nama,
		Email:        request.Email,
		Password:     string(hashedPassword),
		NoTelp:       request.NoTelp,
		Role:         role,
		TanggalLahir: tanggalLahir,
		JenisKelamin: request.JenisKelamin,
		Tentang:      request.Tentang,
		Pekerjaan:    request.Pekerjaan,
		IdProvinsi:   request.IdProvinsi,
		IdKota:       request.IdKota,
	}

	// Simpan ke database
	err = u.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	toko := &entity.Toko{
		IdUser:   user.ID,
	}
	err = u.tokoRepo.Create(ctx, toko)
	if err != nil {
		return nil, err
	}

	// Generate Token
	token, err := common.GenerateAccessToken(ctx, user)
	if err != nil {
		return nil, err
	}

	// Build Response
	response := &dto.RegisterResponse{
		ID:     user.ID,
		Nama:   user.Nama,
		NoTelp: user.NoTelp,
		Email:  user.Email,
		Token:  token,
	}

	return response, nil
}
