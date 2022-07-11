package services

import "golang.org/x/crypto/bcrypt"

type PasswordEncrypt struct {
	cost int
}

func NewPasswordEncrypt(cost int) *PasswordEncrypt {
	return &PasswordEncrypt{cost: cost}
}

func (p *PasswordEncrypt) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	return string(bytes), err
}

func (p *PasswordEncrypt) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
