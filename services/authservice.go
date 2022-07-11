package services

import (
	"accounts-api/pkg/user"
	"errors"
)

type AuthService struct {
	userRepo       user.Repository
	encryptService PasswordEncrypt
}

func NewAuthService(userRepo *user.Repository, encrypt *PasswordEncrypt) *AuthService {
	return &AuthService{
		userRepo:       *userRepo,
		encryptService: *encrypt,
	}
}

func (a *AuthService) Login(email, password string) (*user.User, error) {
	byEmail, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	passwordOk := a.encryptService.CheckPasswordHash(password, byEmail.Password)
	if passwordOk {
		return byEmail, nil
	}
	return nil, errors.New("incorrect password")

}

func (a *AuthService) Signup(user *user.User) (*user.User, error) {
	foundEmail, _ := a.userRepo.FindByEmail(user.Email)
	if foundEmail != nil {
		msg := "[Not Unique]: User with email: " + user.Email
		return nil, errors.New(msg)
	}

	foundUsername, _ := a.userRepo.FindByUsername(user.Username)
	if foundUsername != nil {
		msg := "[Not Unique]: User with username: " + user.Username
		return nil, errors.New(msg)
	}

	hashedPassword, _ := a.encryptService.HashPassword(user.Password)
	user.Password = hashedPassword

	created, err := a.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return created, nil
}
