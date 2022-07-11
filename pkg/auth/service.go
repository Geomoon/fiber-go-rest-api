package auth

import "accounts-api/pkg/user"

type Service interface {
	Login(email, password string) (*user.User, error)
	Signup(user *user.User) (*user.User, error)
}
