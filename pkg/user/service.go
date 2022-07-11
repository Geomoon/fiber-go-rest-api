package user

type Service interface {
	Create(user *User) (*User, error)
	Update(user *User) (*User, error)
	FindAll() *[]User
	FindById(id int) (*User, error)
	DeleteById(id int) error
}
