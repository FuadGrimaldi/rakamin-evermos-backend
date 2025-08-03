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
	Register(ctx context.Context, request dto.RegisterRequest) (string, error)
}

type authService struct {
	cfg        *config.Config
	repository repository.UserRepository
}

func NewAuthService(cfg *config.Config, repository repository.UserRepository) AuthService {
	return &authService{cfg, repository}
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

func (u *authService) Register(ctx context.Context, request dto.RegisterRequest) (string, error) {
	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	role := "user"

	// Lengkapi data user dari request
	user := &entity.User{
		Nama:         request.Nama,
		Email:        request.Email,
		Password:     request.Password,
		NoTelp:       request.NoTelp,
		Role:         role,
		TanggalLahir: tanggalLahir,
		JenisKelamin: request.JenisKelamin,
		Tentang:      request.Tentang,
		Pekerjaan:    request.Pekerjaan,
		IdProvinsi:   request.IdProvinsi,
		IdKota:       request.IdKota,
	}

	

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	if err := u.repository.Create(ctx, user); err != nil {
		return "", err
	}

	token, err := common.GenerateAccessToken(ctx, user)
	if err != nil {
		return "", err
	}

	return token, nil
}