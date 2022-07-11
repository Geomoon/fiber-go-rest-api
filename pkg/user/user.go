package user

import "time"

type User struct {
	Id        int
	Name      string
	LastName  string
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	// TODO Token string
}
