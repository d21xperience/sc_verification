package service

import (
	"auth-service/model"
	"auth-service/repository"
	"auth-service/util"
	"errors"
)

type AuthService interface {
	Register(user *model.User) error
	Login(username, password string) (string, error)
	UpdateProfilePicture(userID string, filepath string) error
	GetUserByID(userID string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
}

type authServiceImpl struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (s *authServiceImpl) Register(user *model.User) error {
	// Cek apakah username sudah ada
	existingUser, _ := s.userRepo.FindByUsername(user.Username)
	if existingUser != nil {
		return errors.New("username already exists")
	}

	// Simpan user baru
	user.Password = util.EncryptPassword(user.Password) // Encrypt password
	return s.userRepo.Save(user)
}

func (s *authServiceImpl) Login(username, password string) (string, error) {
	// Ambil user berdasarkan username
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verifikasi password
	if !util.VerifyPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT
	token, err := util.GenerateJWT(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}

func (s *authServiceImpl) UpdateProfilePicture(userID string, filepath string) error {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user not found")
	}

	user.ProfilePicture = filepath
	return s.userRepo.Save(user)
}

// Get user by ID
func (s *authServiceImpl) GetUserByID(userID string) (*model.User, error) {
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// Get user by username
func (s *authServiceImpl) GetUserByUsername(username string) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
