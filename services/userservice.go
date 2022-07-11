package services

import (
	"accounts-api/pkg/user"
	"errors"
)

type userservice struct {
	repository user.Repository
}

func NewUserService(repository user.Repository) *userservice {
	return &userservice{
		repository: repository,
	}
}

func (s *userservice) Create(user *user.User) (*user.User, error) {
	foundEmail, _ := s.repository.FindByEmail(user.Email)
	if foundEmail != nil {
		msg := "[Not Unique]: User with email: " + user.Email
		return nil, errors.New(msg)
	}

	foundUsername, _ := s.repository.FindByUsername(user.Username)
	if foundUsername != nil {
		msg := "[Not Unique]: User with username: " + user.Username
		return nil, errors.New(msg)
	}

	created, err := s.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *userservice) Update(user *user.User) (*user.User, error) {
	_, err := s.repository.FindById(user.Id)
	if err != nil {
		return nil, err
	}

	updated, errorUpdate := s.repository.Update(user)
	if errorUpdate != nil {
		return nil, errorUpdate
	}

	return updated, nil
}

func (s *userservice) FindAll() *[]user.User {
	var list []user.User

	users := s.repository.FindAll()

	for _, v := range *users {
		list = append(list, v)
	}
	return &list
}

func (s *userservice) FindById(id int) (*user.User, error) {
	return s.repository.FindById(id)
}

func (s *userservice) DeleteById(id int) error {
	return s.repository.DeleteById(id)
}
