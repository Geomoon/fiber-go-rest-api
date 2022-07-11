package user

type Repository interface {
	Create(user *User) (*User, error)
	FindAll() *map[int]User
	FindById(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	Update(user *User) (*User, error)
	DeleteById(id int) error
}
