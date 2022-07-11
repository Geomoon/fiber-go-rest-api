package datasource

import (
	"accounts-api/pkg/user"
	"errors"
	"time"
)

type MKVDatabase struct {
	datasource map[int]user.User
}

func NewMKVDatabase() *MKVDatabase {
	return &MKVDatabase{
		datasource: map[int]user.User{},
	}
}

func (m *MKVDatabase) Create(user *user.User) (*user.User, error) {
	user.Id = len(m.datasource)
	user.CreatedAt = time.Now()
	m.datasource[user.Id] = *user
	return user, nil
}

func (m *MKVDatabase) FindAll() *map[int]user.User {
	return &m.datasource
}

func (m *MKVDatabase) FindById(id int) (*user.User, error) {
	u, exists := m.datasource[id]
	if exists {
		return &u, nil
	}
	msg := "[Not Found]: User with id: " + string(rune(id))
	return nil, errors.New(msg)
}

func (m *MKVDatabase) FindByEmail(email string) (*user.User, error) {
	for _, v := range m.datasource {
		if v.Email == email {
			return &v, nil
		}
	}
	msg := "[Not Found]: User with email: " + email
	return nil, errors.New(msg)
}

func (m *MKVDatabase) FindByUsername(username string) (*user.User, error) {
	for _, v := range m.datasource {
		if v.Username == username {
			return &v, nil
		}
	}
	msg := "[Not Found]: User with username: " + username
	return nil, errors.New(msg)
}

func (m *MKVDatabase) Update(user *user.User) (*user.User, error) {
	found, exists := m.datasource[user.Id]
	if !exists {
		msg := "[Not Found]: User with id: " + string(rune(user.Id))
		return nil, errors.New(msg)
	}
	found = *user
	return &found, nil
}

func (m *MKVDatabase) DeleteById(id int) error {
	_, exists := m.datasource[id]
	if !exists {
		msg := "[Not Found]: User with id: " + string(rune(id))
		return errors.New(msg)
	}
	delete(m.datasource, id)
	return nil
}
