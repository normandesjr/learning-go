package user

import (
	"log"
)

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	log.Printf("User: Sending user email to %s<%s>\n", u.Name, u.Email)
	return nil
}

type Admin struct {
	User
	Level string
}

func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin email to %s<%s>\n", a.Name, a.Email)
	return nil
}
