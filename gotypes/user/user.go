package user

import "app/email"

type User struct {
	Name  string
	Email string
}

func NewUser(name string, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func (u *User) Welcome() {
	email.Send(u.Email, "Welcome!")
}
